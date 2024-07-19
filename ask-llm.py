#!/usr/bin/env python3

import asyncio
import os
import json
import urllib.request
import time

LLM_API_BASE_URL = os.environ.get("LLM_API_BASE_URL", "https://api.openai.com/v1")
LLM_API_KEY = os.environ.get("LLM_API_KEY") or os.environ.get("OPENAI_API_KEY")
LLM_CHAT_MODEL = os.environ.get("LLM_CHAT_MODEL")
LLM_STREAMING = os.environ.get("LLM_STREAMING", "yes") != "no"

LLM_DEBUG = os.environ.get("LLM_DEBUG")


async def chat(messages, handler=None):
    url = f"{LLM_API_BASE_URL}/chat/completions"
    auth_header = f"Bearer {LLM_API_KEY}" if LLM_API_KEY else None
    headers = {
        "Content-Type": "application/json",
        "User-Agent": "python-requests/2.31.0",
    }
    if auth_header:
        headers["Authorization"] = auth_header

    model = LLM_CHAT_MODEL or "gpt-4o-mini"
    stop = ["<|im_end|>", "<|end|>", "<|eot_id|>"]
    max_tokens = 200
    temperature = 0
    stream = LLM_STREAMING and callable(handler)

    body = {
        "messages": messages,
        "model": model,
        "stop": stop,
        "max_tokens": max_tokens,
        "temperature": temperature,
        "stream": stream,
    }
    json_body = json.dumps(body).encode("utf-8")
    request = urllib.request.Request(
        url, data=json_body, headers=headers, method="POST"
    )
    response = urllib.request.urlopen(request)

    if not stream:
        if response.status != 200:
            raise Exception(f"HTTP error: {response.status} {response.reason}")
        data = json.loads(response.read().decode("utf-8"))
        choices = data["choices"]
        first = choices[0]
        message = first["message"]
        content = message["content"]
        full_answer = content.strip()
        if handler:
            handler(full_answer)
        return full_answer
    else:

        def parse(line):
            partial = None
            prefix = line[:6]
            if prefix == "data: ":
                payload = line[6:]
                try:
                    choices = json.loads(payload)["choices"]
                    choice = choices[0]
                    delta = choice.get("delta", {})
                    partial = delta.get("content", "")
                except Exception as e:
                    pass
            return partial

        finished = False
        buffer = []
        answer = ""
        while not finished:
            raw_bytes = response.read(8)
            if not raw_bytes:
                break
            buffer.append(raw_bytes)
            lines = b"".join(buffer).decode("utf-8").splitlines(True)
            full_answer = ""
            for line in lines:
                if len(line) > 0:
                    if line[0] == ":":
                        continue
                    if line == "data: [DONE]":
                        finished = True
                        break
                    elif line:
                        partial = parse(line)
                        if partial is not None:
                            if len(full_answer) == 0:
                                full_answer = partial.strip()
                            else:
                                full_answer += partial
            if handler:
                handler(full_answer.replace(answer, ""))
            answer = full_answer

        return answer


SYSTEM_PROMPT = "Answer the question politely and concisely."


async def main():
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
        stream = lambda partial: print(partial, end="", flush=True)
        answer = await chat(messages, stream)
        messages.append({"role": "assistant", "content": answer})
        print()
        elapsed = time.time() - start
        if LLM_DEBUG:
            print(f"[{round(elapsed * 1000)} ms]")
        print()


asyncio.run(main())
