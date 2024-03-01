# Ask LLM

This is a straightforward CLI tool to interact with any LLM service.

It is available in several flavors:

* Python version. Compatible with [CPython](https://python.org) or [PyPy](https://pypy.org),  v3.10 or higher.
* JavaScript version. Compatible with [Node.js](https://nodejs.org) (>= v18) or [Bun](https://bun.sh) (>= v1.0).
* Clojure version. Compatible with [Babashka](https://babashka.org/) (>= 1.3).

To use it locally with [llama.cpp](https://github.com/ggerganov/llama.cpp) inference engine (or its wrapper such as [Nitro](https://nitro.jan.ai/), [LocalAI](https://localai.io/), etc), make sure to load a suitable model that utilizes the [ChatML format](https://github.com/openai/openai-python/blob/release-v0.28.0/chatml.md) (example: [TinyLLama](https://huggingface.co/TheBloke/TinyLlama-1.1B-Chat-v1.0-GGUF), [OpenHermes 2.5](https://huggingface.co/TheBloke/OpenHermes-2.5-Mistral-7B-GGUF), etc). Set the environment variable `LLM_API_BASE_URL` accordingly:
```bash
~/llama.cpp/server -m tinyllama-1.1b-chat-v1.0.Q4_K_M.gguf
export LLM_API_BASE_URL=http://127.0.0.1:8080/v1
./ask-llm.py  # for Python user
./ask-llm.js  # for Node.js user
./ask-llm.clj # for Clojure user
```

To use it locally with [Ollama](https://ollama.com/), start by loading a model and setting the environment variable `LLM_API_BASE_URL`:
```bash
ollama pull gemma:2b
export LLM_API_BASE_URL=http://127.0.0.1:11434/v1
export LLM_CHAT_MODEL='gemma:2b'
./ask-llm.py  # for Python user
./ask-llm.js  # for Node.js user
./ask-llm.clj # for Clojure user
```

For integration with the [OpenAI GPT model](https://platform.openai.com/docs), set the environment variable `OPENAI_API_KEY` to your OpenAI API key:
```bash
export OPENAI_API_KEY="sk-yourownapikey"
./ask-llm.py  # for Python user
./ask-llm.js  # for Node.js user
./ask-llm.clj # for Clojure user
```

To use it with other LLM services, populate relevant environment variables as demonstrated in the examples below:

* [OpenRouter](https://openrouter.ai/)
```bash
export LLM_API_BASE_URL=https://openrouter.ai/api/v1
export LLM_API_KEY="sk-yourownapikey"
export LLM_CHAT_MODEL="mistralai/mistral-7b-instruct"
./ask-llm.py  # for Python user
./ask-llm.js  # for Node.js user
./ask-llm.clj # for Clojure user
```

* [Together](https://www.together.ai/)
```bash
export LLM_API_BASE_URL=https://api.together.xyz/v1
export LLM_API_KEY="sk-yourownapikey"
export LLM_CHAT_MODEL="mistralai/Mistral-7B-Instruct-v0.2"
./ask-llm.py  # for Python user
./ask-llm.js  # for Node.js user
./ask-llm.clj # for Clojure user
```

* [Fireworks](https://fireworks.ai/):
```bash
export LLM_API_BASE_URL=https://api.fireworks.ai/inference/v1
export LLM_API_KEY="yourownapikey"
export LLM_CHAT_MODEL="accounts/fireworks/models/mistral-7b-instruct-4k"
./ask-llm.py  # for Python user
./ask-llm.js  # for Node.js user
./ask-llm.clj # for Clojure user
```

* [Deep Infra](https://deepinfra.com):
```bash
export LLM_API_BASE_URL=https://api.deepinfra.com/v1/openai
export LLM_API_KEY="yourownapikey"
export LLM_CHAT_MODEL="mistralai/Mistral-7B-Instruct-v0.1"
./ask-llm.py  # for Python user
./ask-llm.js  # for Node.js user
./ask-llm.clj # for Clojure user
```