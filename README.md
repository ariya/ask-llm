# Ask LLM

[![Test on OpenAI](https://github.com/ariya/ask-llm/actions/workflows/test-openai.yml/badge.svg)](https://github.com/ariya/ask-llm/actions/workflows/test-openai.yml) [![Test on DeepInfra](https://github.com/ariya/ask-llm/actions/workflows/test-deepinfra.yml/badge.svg)](https://github.com/ariya/ask-llm/actions/workflows/test-deepinfra.yml) [![Test on Fireworks](https://github.com/ariya/ask-llm/actions/workflows/test-fireworks.yml/badge.svg)](https://github.com/ariya/ask-llm/actions/workflows/test-fireworks.yml) [![Test on Groq](https://github.com/ariya/ask-llm/actions/workflows/test-groq.yml/badge.svg)](https://github.com/ariya/ask-llm/actions/workflows/test-groq.yml) [![Test on Lepton](https://github.com/ariya/ask-llm/actions/workflows/test-lepton.yml/badge.svg)](https://github.com/ariya/ask-llm/actions/workflows/test-lepton.yml) [![Test on Novita](https://github.com/ariya/ask-llm/actions/workflows/test-novita.yml/badge.svg)](https://github.com/ariya/ask-llm/actions/workflows/test-novita.yml) [![Test on Octo](https://github.com/ariya/ask-llm/actions/workflows/test-octo.yml/badge.svg)](https://github.com/ariya/ask-llm/actions/workflows/test-octo.yml) [![Test on OpenRouter](https://github.com/ariya/ask-llm/actions/workflows/test-openrouter.yml/badge.svg)](https://github.com/ariya/ask-llm/actions/workflows/test-openrouter.yml) [![Test on Together](https://github.com/ariya/ask-llm/actions/workflows/test-together.yml/badge.svg)](https://github.com/ariya/ask-llm/actions/workflows/test-together.yml)

[![asciicast](https://asciinema.org/a/646222.svg)](https://asciinema.org/a/646222)

This is a straightforward, zero-dependency CLI tool to interact with any LLM service.

It is available in several flavors:

* Python version. Compatible with [CPython](https://python.org) or [PyPy](https://pypy.org),  v3.10 or higher.
* JavaScript version. Compatible with [Node.js](https://nodejs.org) (>= v18) or [Bun](https://bun.sh) (>= v1.0).
* Clojure version. Compatible with [Babashka](https://babashka.org/) (>= 1.3).
* Go version. Compatible with [Go](https://golang.org), v1.19 or higher.

Ask LLM is compatible with either a cloud-based (managed) LLM service (e.g. [OpenAI GPT model](https://platform.openai.com/docs), [Groq](https://groq.com), [OpenRouter](https://openrouter.ai), etc) or with a locally hosted LLM server (e.g. [llama.cpp](https://github.com/ggerganov/llama.cpp), [LocalAI](https://localai.io), [Ollama](https://ollama.com), etc). Please continue reading for detailed instructions.

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

Supported local LLM servers include [llama.cpp](https://github.com/ggerganov/llama.cpp), [Jan](https://jan.ai), [Ollama](https://ollama.com), and [LocalAI](https://localai.io).

To utilize [llama.cpp](https://github.com/ggerganov/llama.cpp) locally with its inference engine, ensure to load a quantized model such as [Phi-3 Mini](https://huggingface.co/microsoft/Phi-3-mini-4k-instruct-gguf), [LLama-3 8B](https://huggingface.co/QuantFactory/Meta-Llama-3-8B-Instruct-GGUF), or [OpenHermes 2.5](https://huggingface.co/TheBloke/OpenHermes-2.5-Mistral-7B-GGUF). Adjust the environment variable `LLM_API_BASE_URL` accordingly:
```bash
/path/to/llama.cpp/server -m Phi-3-mini-4k-instruct-q4.gguf
export LLM_API_BASE_URL=http://127.0.0.1:8080/v1
```

To use [Jan](https://jan.ai) with its local API server, refer to [its documentation](https://jan.ai/docs/local-api) and load a model like [Phi-3 Mini](https://huggingface.co/microsoft/Phi-3-mini-4k-instruct-gguf), [LLama-3 8B](https://huggingface.co/QuantFactory/Meta-Llama-3-8B-Instruct-GGUF), or [OpenHermes 2.5](https://huggingface.co/TheBloke/OpenHermes-2.5-Mistral-7B-GGUF) and set the environment variable `LLM_API_BASE_URL`:
```bash
export LLM_API_BASE_URL=http://127.0.0.1:1337/v1
export LLM_CHAT_MODEL='llama3-8b-instruct'
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

Supported LLM services include [Deep Infra](https://deepinfra.com), [Fireworks](https://fireworks.ai), [Groq](https://groq.com), [Hyperbolic](https://www.hyperbolic.xyz), [Lepton](https://lepton.ai), [Novita](https://novita.ai), [Octo](https://octo.ai), [OpenAI](https://platform.openai.com), [OpenRouter](https://openrouter.ai), and [Together](https://www.together.ai).

For configuration specifics, refer to the relevant section. The examples use Llama-3.1 8B (or GPT-4o Mini for OpenAI), but any LLM with at least 7B parameters should work just as well, such as Mistral 7B, Qwen-2 7B, or Gemma-2 9B.


* [Deep Infra](https://deepinfra.com)
```bash
export LLM_API_BASE_URL=https://api.deepinfra.com/v1/openai
export LLM_API_KEY="yourownapikey"
export LLM_CHAT_MODEL="meta-llama/Meta-Llama-3.1-8B-Instruct"
```

* [Fireworks](https://fireworks.ai/)
```bash
export LLM_API_BASE_URL=https://api.fireworks.ai/inference/v1
export LLM_API_KEY="yourownapikey"
export LLM_CHAT_MODEL="accounts/fireworks/models/llama-v3p1-8b-instruct"
```

* [Groq](https://groq.com/)
```bash
export LLM_API_BASE_URL=https://api.groq.com/openai/v1
export LLM_API_KEY="yourownapikey"
export LLM_CHAT_MODEL="llama-3.1-8b-instant"
```

* [Hyperbolic](https://www.hyperbolic.xyz)
```bash
export LLM_API_BASE_URL=https://api.hyperbolic.xyz/v1
export LLM_API_KEY="yourownapikey"
export LLM_CHAT_MODEL="meta-llama/Meta-Llama-3.1-8B-Instruct"
```

* [Lepton](https://lepton.ai)
```bash
export LLM_API_BASE_URL=https://llama3-1-8b.lepton.run/api/v1
export LLM_API_KEY="yourownapikey"
export LLM_CHAT_MODEL="llama3-1-8b"
```

* [Novita](https://novita.ai)
```bash
export LLM_API_BASE_URL=https://api.novita.ai/v3/openai
export LLM_API_KEY="yourownapikey"
export LLM_CHAT_MODEL="meta-llama/llama-3.1-8b-instruct"
```

* [Octo](https://octo.ai)
```bash
export LLM_API_BASE_URL=https://text.octoai.run/v1/
export LLM_API_KEY="yourownapikey"
export LLM_CHAT_MODEL="meta-llama-3.1-8b-instruct"
```

* [OpenAI](https://platform.openai.com)
```bash
export LLM_API_BASE_URL=https://api.openai.com/v1
export LLM_API_KEY="yourownapikey"
export LLM_CHAT_MODEL="gpt-4o-mini"
```

* [OpenRouter](https://openrouter.ai/)
```bash
export LLM_API_BASE_URL=https://openrouter.ai/api/v1
export LLM_API_KEY="yourownapikey"
export LLM_CHAT_MODEL="meta-llama/llama-3.1-8b-instruct"
```

* [Together](https://www.together.ai/)
```bash
export LLM_API_BASE_URL=https://api.together.xyz/v1
export LLM_API_KEY="yourownapikey"
export LLM_CHAT_MODEL="meta-llama/Meta-Llama-3.1-8B-Instruct-Turbo"
```
