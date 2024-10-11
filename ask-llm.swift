#!/usr/bin/env swift

import Foundation
import os

let LLM_API_BASE_URL =
  ProcessInfo.processInfo.environment["LLM_API_BASE_URL"] ?? "https://api.openai.com/v1"
let LLM_API_KEY =
  ProcessInfo.processInfo.environment["LLM_API_KEY"]
  ?? ProcessInfo.processInfo.environment["OPENAI_API_KEY"]
let LLM_CHAT_MODEL = ProcessInfo.processInfo.environment["LLM_CHAT_MODEL"]
let LLM_DEBUG = ProcessInfo.processInfo.environment["LLM_DEBUG"]

struct ChatRequest: Encodable {
  let messages: [[String: String]]
  let model: String
  let stop: [String]
  let maxTokens: Int
  let temperature: Double
  let stream: Bool

  enum CodingKeys: String, CodingKey {
    case messages, model, stop, temperature, stream
    case maxTokens = "max_tokens"
  }
}

struct ChatResponse: Decodable {
  let choices: [Choice]
}

struct Choice: Decodable {
  let message: Message
}

struct Message: Decodable {
  let content: String
}

let SYSTEM_PROMPT = "Answer the question politely and concisely."

func chat(messages: [[String: String]], handler: ((String) -> Void)? = nil) async throws -> String {
  let url = URL(string: "\(LLM_API_BASE_URL)/chat/completions")!

  let body = ChatRequest(
    messages: messages,
    model: LLM_CHAT_MODEL ?? "gpt-4o-mini",
    stop: ["<|im_end|>", "<|end|>", "<|eot_id|>"],
    maxTokens: 200,
    temperature: 0,
    stream: false
  )

  var request = URLRequest(url: url)
  request.httpMethod = "POST"
  request.addValue("application/json", forHTTPHeaderField: "Content-Type")
  request.addValue("python-requests/2.31.0", forHTTPHeaderField: "User-Agent")
  if let apiKey = LLM_API_KEY {
    request.addValue("Bearer \(apiKey)", forHTTPHeaderField: "Authorization")
  }
  request.httpBody = try JSONEncoder().encode(body)

  let (data, response) = try await URLSession.shared.data(for: request)

  guard let httpResponse = response as? HTTPURLResponse, httpResponse.statusCode == 200 else {
    let statusCode = (response as? HTTPURLResponse)?.statusCode ?? -1
    throw NSError(
      domain: "LLMError", code: 1,
      userInfo: [NSLocalizedDescriptionKey: "HTTP error: \(statusCode)"])
  }

  let decodedData = try JSONDecoder().decode(ChatResponse.self, from: data)

  if let answer = decodedData.choices.first?.message.content {
    if let handler = handler {
      handler(answer)
    }
    return answer
  }
  return ""
}

func main() async {
  print("Using LLM at \(LLM_API_BASE_URL).")
  print("Press Ctrl+D to exit.")
  print()

  var messages: [[String: String]] = [
    ["role": "system", "content": SYSTEM_PROMPT]
  ]

  while true {
    print(">>", terminator: " ")

    if let input = readLine() {
      messages.append(["role": "user", "content": input])

      let startTime = Date()
      do {
        let streamHandler: (String) -> Void = { partial in
          print(partial, terminator: "")
        }
        let answer = try await chat(messages: messages, handler: streamHandler)
        messages.append(["role": "assistant", "content": answer])
        print()
        let endTime = Date()
        let elapsed = endTime.timeIntervalSince(startTime)
        if LLM_DEBUG != nil {
          print("[\(Int(elapsed * 1000)) ms]")
        }
        print()
      } catch {
        print("Error: \(error)")
      }
    } else {
      break
    }
  }
}

await main()
