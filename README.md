# Ask LLM

[![asciicast](https://asciinema.org/a/646222.svg)](https://asciinema.org/a/646222)

This is a straightforward, zero-dependency CLI tool to interact with any LLM service.

It is available in several flavors:

* Python version. Compatible with [CPython](https://python.org) or [PyPy](https://pypy.org),  v3.10 or higher.
* JavaScript version. Compatible with [Node.js](https://nodejs.org) (>= v18) or [Bun](https://bun.sh) (>= v1.0).
* Clojure version. Compatible with [Babashka](https://babashka.org/) (>= 1.3).

Once a suitable inference engine is set up (local or remote, read the next section), interact with the LLM:
```bash
./ask-llm.py  # for Python user
./ask-llm.js  # for Node.js user
./ask-llm.clj # for Clojure user
```

or pipe the question directly to get an immediate answer:
```bash
echo "Why is the sky blue?" | ./ask-llm.py
```

or request the LLM to perform a certain task:
```bash
echo "Translate into German: thank you" | ./ask-llm.py
```

To use it locally with [llama.cpp](https://github.com/ggerganov/llama.cpp) inference engine, make sure to load a quantized model (example: [TinyLLama](https://huggingface.co/TheBloke/TinyLlama-1.1B-Chat-v1.0-GGUF), [Gemma 2B](https://huggingface.co/google/gemma-2b-it-GGUF), [OpenHermes 2.5](https://huggingface.co/TheBloke/OpenHermes-2.5-Mistral-7B-GGUF), etc) with the suitable chat template. Set the environment variable `LLM_API_BASE_URL` accordingly:
```bash
~/llama.cpp/server -m gemma-2b-it-q4_k_m.gguf --chat-template gemma
export LLM_API_BASE_URL=http://127.0.0.1:8080/v1
```

To use it locally with [Nitro](https://nitro.jan.ai/), follow its [Quickstart guide](https://nitro.jan.ai/quickstart#step-4-load-model) to load a model (e.g. [TinyLLama](https://huggingface.co/TheBloke/TinyLlama-1.1B-Chat-v1.0-GGUF), [OpenHermes 2.5](https://huggingface.co/TheBloke/OpenHermes-2.5-Mistral-7B-GGUF), etc) and set the environment variable `LLM_API_BASE_URL`:
```bash
export LLM_API_BASE_URL=http://localhost:3928/v1
```

To use it locally with [Ollama](https://ollama.com/), load a model and set the environment variable `LLM_API_BASE_URL`:
```bash
ollama pull gemma:2b
export LLM_API_BASE_URL=http://127.0.0.1:11434/v1
export LLM_CHAT_MODEL='gemma:2b'
```

To use it locally with [LocalAI](https://localai.io), launch its container and the set environment variable `LLM_API_BASE_URL`:
```bash
docker run -ti -p 8080:8080 localai/localai tinyllama-chat
export LLM_API_BASE_URL=http://localhost:3928/v1
```

To use [OpenAI GPT model](https://platform.openai.com/docs), set the environment variable `OPENAI_API_KEY` to your API key:
```bash
export OPENAI_API_KEY="sk-yourownapikey"
```

To use it with other LLM services, populate relevant environment variables as shown in these examples:

* [Deep Infra](https://deepinfra.com)
```bash
export LLM_API_BASE_URL=https://api.deepinfra.com/v1/openai
export LLM_API_KEY="yourownapikey"
export LLM_CHAT_MODEL="mistralai/Mistral-7B-Instruct-v0.1"
```

* [Fireworks](https://fireworks.ai/)
```bash
export LLM_API_BASE_URL=https://api.fireworks.ai/inference/v1
export LLM_API_KEY="yourownapikey"
export LLM_CHAT_MODEL="accounts/fireworks/models/mistral-7b-instruct-4k"
```

* [Lepton](https://lepton.ai)
```bash
export LLM_API_BASE_URL=https://mixtral-8x7b.lepton.run/api/v1/
export LLM_API_KEY="yourownapikey"
```

* [OpenRouter](https://openrouter.ai/)
```bash
export LLM_API_BASE_URL=https://openrouter.ai/api/v1
export LLM_API_KEY="sk-yourownapikey"
export LLM_CHAT_MODEL="mistralai/mistral-7b-instruct"
```

* [Together](https://www.together.ai/)
```bash
export LLM_API_BASE_URL=https://api.together.xyz/v1
export LLM_API_KEY="sk-yourownapikey"
export LLM_CHAT_MODEL="mistralai/Mistral-7B-Instruct-v0.2"
```