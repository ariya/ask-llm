# Ask LLM

[![asciicast](https://asciinema.org/a/646222.svg)](https://asciinema.org/a/646222)

This is a straightforward, zero-dependency CLI tool to interact with any LLM service.

It is available in several flavors:

* Python version. Compatible with [CPython](https://python.org) or [PyPy](https://pypy.org),  v3.10 or higher.
* JavaScript version. Compatible with [Node.js](https://nodejs.org) (>= v18) or [Bun](https://bun.sh) (>= v1.0).
* Clojure version. Compatible with [Babashka](https://babashka.org/) (>= 1.3).
* Go version. Compatible with [Go](https://golang.org), v1.19 or higher.

Ask LLM is compatible with either a cloud-based (managed) LLM service (e.g. [OpenAI GPT model](https://platform.openai.com/docs), [Grog](https://groq.com), [OpenRouter](https://openrouter.ai), etc) or with a locally hosted LLM server (e.g. [llama.cpp](https://github.com/ggerganov/llama.cpp), [LocalAI](https://localai.io), [Ollama](https://ollama.com), etc). Please continue reading for detailed instructions.

Interact with the LLM with:
```bash
./ask-llm.py         # for Python user
./ask-llm.js         # for Node.js user
./ask-llm.clj        # for Clojure user
go run ask-llm.go    # for Go user
```

or pipe the question directly to get an immediate answer:
```bash
echo "Why is the sky blue?" | ./ask-llm.py
```

or request the LLM to perform a certain task:
```bash
echo "Translate into German: thank you" | ./ask-llm.py
```

## Using Local LLM Servers

Supported local LLM servers include [llama.cpp](https://github.com/ggerganov/llama.cpp), [Nitro](https://nitro.jan.ai), [Ollama](https://ollama.com), and [LocalAI](https://localai.io).

To utilize [llama.cpp](https://github.com/ggerganov/llama.cpp) locally with its inference engine, ensure to load a quantized model such as [Phi-3 Mini](https://huggingface.co/microsoft/Phi-3-mini-4k-instruct-gguf), [LLama-3 8B](https://huggingface.co/QuantFactory/Meta-Llama-3-8B-Instruct-GGUF), or [OpenHermes 2.5](https://huggingface.co/TheBloke/OpenHermes-2.5-Mistral-7B-GGUF). Adjust the environment variable `LLM_API_BASE_URL` accordingly:
```bash
/path/to/llama.cpp/server -m Phi-3-mini-4k-instruct-q4.gguf
export LLM_API_BASE_URL=http://127.0.0.1:8080/v1
```

To utilize [Nitro](https://nitro.jan.ai) locally, refer to its [Quickstart guide](https://nitro.jan.ai/quickstart#step-4-load-model) for loading a model like [Phi-3 Mini](https://huggingface.co/microsoft/Phi-3-mini-4k-instruct-gguf), [LLama-3 8B](https://huggingface.co/QuantFactory/Meta-Llama-3-8B-Instruct-GGUF), or [OpenHermes 2.5](https://huggingface.co/TheBloke/OpenHermes-2.5-Mistral-7B-GGUF) and set the environment variable `LLM_API_BASE_URL`:
```bash
export LLM_API_BASE_URL=http://localhost:3928/v1
```

To use [Ollama](https://ollama.com) locally, load a model and configure the environment variable `LLM_API_BASE_URL`:
```bash
ollama pull phi3
export LLM_API_BASE_URL=http://127.0.0.1:11434/v1
export LLM_CHAT_MODEL='phi3'
```

For [LocalAI](https://localai.io), initiate its container and adjust the environment variable `LLM_API_BASE_URL`:
```bash
docker run -ti -p 8080:8080 localai/localai tinyllama-chat
export LLM_API_BASE_URL=http://localhost:3928/v1
```

## Using Managed LLM Services

To use [OpenAI GPT model](https://platform.openai.com/docs), configure the environment variable `OPENAI_API_KEY` with your API key:
```bash
export OPENAI_API_KEY="sk-yourownapikey"
```

To utilize other LLM services, populate the relevant environment variables as demonstrated in the following examples:

* [Anyscale](https://www.anyscale.com/)
```bash
export LLM_API_BASE_URL=https://api.endpoints.anyscale.com/v1
export LLM_API_KEY="yourownapikey"
export LLM_CHAT_MODEL="meta-llama/Llama-3-8b-chat-hf"
```

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
export LLM_CHAT_MODEL="accounts/fireworks/models/llama-v3-8b-instruct"
```

* [Grog](https://groq.com/)
```bash
export LLM_API_BASE_URL=https://api.groq.com/openai/v1
export LLM_API_KEY="yourownapikey"
export LLM_CHAT_MODEL="gemma-7b-it"
```

* [Lepton](https://lepton.ai)
```bash
export LLM_API_BASE_URL=https://mixtral-8x7b.lepton.run/api/v1/
export LLM_API_KEY="yourownapikey"
```

* [OpenRouter](https://openrouter.ai/)
```bash
export LLM_API_BASE_URL=https://openrouter.ai/api/v1
export LLM_API_KEY="yourownapikey"
export LLM_CHAT_MODEL="mistralai/mistral-7b-instruct:free"
```

* [Together](https://www.together.ai/)
```bash
export LLM_API_BASE_URL=https://api.together.xyz/v1
export LLM_API_KEY="yourownapikey"
export LLM_CHAT_MODEL="meta-llama/Llama-3-8b-chat-hf"
```