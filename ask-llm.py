#!/usr/bin/env python3

import os
import json
import urllib.request
import time

LLM_API_BASE_URL = os.environ.get("LLM_API_BASE_URL", "https://api.openai.com/v1")
LLM_API_KEY = os.environ.get("LLM_API_KEY") or os.environ.get("OPENAI_API_KEY")
LLM_CHAT_MODEL = os.environ.get("LLM_CHAT_MODEL")

LLM_DEBUG = os.environ.get("LLM_DEBUG")


def chat(messages):
    url = f"{LLM_API_BASE_URL}/chat/completions"
    auth_header = f"Bearer {LLM_API_KEY}" if LLM_API_KEY else None
    headers = {
        "Content-Type": "application/json",
        "User-Agent": "python-requests/2.31.0",
    }
    if auth_header:
        headers["Authorization"] = auth_header
    body = {
        "messages": messages,
        "model": LLM_CHAT_MODEL or "gpt-3.5-turbo",
        "max_tokens": 200,
        "temperature": 0,
    }
    json_body = json.dumps(body).encode("utf-8")
    request = urllib.request.Request(url, data=json_body, headers=headers)
    response = urllib.request.urlopen(request)
    if response.status != 200:
        raise Exception(f"HTTP error: {response.status} {response.reason}")
    data = json.loads(response.read().decode("utf-8"))
    choices = data["choices"]
    first = choices[0]
    message = first["message"]
    content = message["content"]
    answer = content.strip()
    return answer


SYSTEM_PROMPT = "Answer the question politely and concisely."


def main():
    print(f"Using LLM at {LLM_API_BASE_URL}.")
    print("Press Ctrl+D to exit.")
    print()

    messages = [{"role": "system", "content": SYSTEM_PROMPT}]

    while True:
        try:
            question = input(">> ")
        except EOFError:
            break

        messages.append({"role": "user", "content": question})
        start = time.time()
        answer = chat(messages)
        messages.append({"role": "assistant", "content": answer})
        print(answer)
        elapsed = time.time() - start
        if LLM_DEBUG:
            print(f"[{round(elapsed, 3)} s]")
        print()


if __name__ == "__main__":
    main()
