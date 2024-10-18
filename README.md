# Ask LLM

[![asciicast](https://asciinema.org/a/646222.svg)](https://asciinema.org/a/646222)

This is a straightforward, zero-dependency CLI tool to interact with any LLM service.

It is available in several flavors:

* Python version. Compatible with [CPython](https://python.org) or [PyPy](https://pypy.org),  v3.10 or higher.
* JavaScript version. Compatible with [Node.js](https://nodejs.org) (>= v18) or [Bun](https://bun.sh) (>= v1.0).
* Clojure version. Compatible with [Babashka](https://babashka.org/) (>= 1.3).
* Swift version. Compatible with [Swift](https://www.swift.org), v5.10 or higher.
* Go version. Compatible with [Go](https://golang.org), v1.19 or higher.

Ask LLM is compatible with either a cloud-based (managed) LLM service (e.g. [OpenAI GPT model](https://platform.openai.com/docs), [Groq](https://groq.com), [OpenRouter](https://openrouter.ai), etc) or with a locally hosted LLM server (e.g. [llama.cpp](https://github.com/ggerganov/llama.cpp), [LM Studio](https://lmstudio.ai), [Ollama](https://ollama.com), etc). Please continue reading for detailed instructions.

Interact with the LLM with:
```bash
./ask-llm.py         # for Python user
./ask-llm.js         # for Node.js user
./ask-llm.clj        # for Clojure user
./ask-llm.swift      # for Swift user
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

Supported local LLM servers include [llama.cpp](https://github.com/ggerganov/llama.cpp), [Jan](https://jan.ai), [Ollama](https://ollama.com), [LocalAI](https://localai.io), [LM Studio](https://lmstudio.ai), and [Msty](https://msty.app).

To utilize [llama.cpp](https://github.com/ggerganov/llama.cpp) locally with its inference engine, load a quantized model like [Llama-3.2 3B](https://huggingface.co/bartowski/Llama-3.2-3B-Instruct-GGUF) or [Phi-3.5 Mini](https://huggingface.co/bartowski/Phi-3.5-mini-instruct-GGUF). Then set the `LLM_API_BASE_URL` environment variable:
```bash
/path/to/llama-server -m Llama-3.2-3B-Instruct-Q4_K_M.gguf
export LLM_API_BASE_URL=http://127.0.0.1:8080/v1
```

To use [Jan](https://jan.ai) with its local API server, refer to [its documentation](https://jan.ai/docs/local-api). Load a model like [Llama-3.2 3B](https://huggingface.co/bartowski/Llama-3.2-3B-Instruct-GGUF) or [Phi-3.5 Mini](https://huggingface.co/bartowski/Phi-3.5-mini-instruct-GGUF), and set the following environment variables:
```bash
export LLM_API_BASE_URL=http://127.0.0.1:1337/v1
export LLM_CHAT_MODEL='llama3-8b-instruct'
```

To use [Ollama](https://ollama.com) locally, load a model and configure the environment variable `LLM_API_BASE_URL`:
```bash
ollama pull llama3.2
export LLM_API_BASE_URL=http://127.0.0.1:11434/v1
export LLM_CHAT_MODEL='llama3.2'
```

For [LocalAI](https://localai.io), initiate its container and adjust the environment variable `LLM_API_BASE_URL`:
```bash
docker run -ti -p 8080:8080 localai/localai llama-3.2-3b-instruct:q4_k_m
export LLM_API_BASE_URL=http://localhost:3928/v1
```

For [LM Studio](https://lmstudio.ai), pick a model (e.g., Llama-3.2 3B). Next, go to the Developer tab, select the model to load, and click the Start Server button. Then, set the `LLM_API_BASE_URL` environment variable, noting that the server by default runs on port `1234`:
```bash
export LLM_API_BASE_URL=http://127.0.0.1:1234/v1
```

For [Msty](https://msty.app), choose a model (e.g., Llama-3.2 3B) and ensure the local AI is running. Go to the Settings menu, under Local AI, and note the Service Endpoint (which defaults to port `10002`). Then set the `LLM_API_BASE_URL` environment variable accordingly:
```bash
export LLM_API_BASE_URL=http://127.0.0.1:10002/v1
```

## Using Managed LLM Services

[![Test on AI21](https://github.com/ariya/ask-llm/actions/workflows/test-ai21.yml/badge.svg)](https://github.com/ariya/ask-llm/actions/workflows/test-ai21.yml)
[![Test on DeepInfra](https://github.com/ariya/ask-llm/actions/workflows/test-deepinfra.yml/badge.svg)](https://github.com/ariya/ask-llm/actions/workflows/test-deepinfra.yml)
[![Test on DeepSeek](https://github.com/ariya/ask-llm/actions/workflows/test-deepseek.yml/badge.svg)](https://github.com/ariya/ask-llm/actions/workflows/test-deepseek.yml)
[![Test on Fireworks](https://github.com/ariya/ask-llm/actions/workflows/test-fireworks.yml/badge.svg)](https://github.com/ariya/ask-llm/actions/workflows/test-fireworks.yml)
[![Test on Groq](https://github.com/ariya/ask-llm/actions/workflows/test-groq.yml/badge.svg)](https://github.com/ariya/ask-llm/actions/workflows/test-groq.yml)
[![Test on Hyperbolic](https://github.com/ariya/ask-llm/actions/workflows/test-hyperbolic.yml/badge.svg)](https://github.com/ariya//ask-llm/actions/workflows/test-hyperbolic.yml)
[![Test on Lepton](https://github.com/ariya/ask-llm/actions/workflows/test-lepton.yml/badge.svg)](https://github.com/ariya/ask-llm/actions/workflows/test-lepton.yml)
[![Test on Mistral](https://github.com/ariya/ask-llm/actions/workflows/test-mistral.yml/badge.svg)](https://github.com/ariya/ask-llm/actions/workflows/test-mistral.yml)
[![Test on Nebius](https://github.com/ariya/ask-llm/actions/workflows/test-nebius.yml/badge.svg)](https://github.com/ariya/ask-llm/actions/workflows/test-nebius.yml)
[![Test on Novita](https://github.com/ariya/ask-llm/actions/workflows/test-novita.yml/badge.svg)](https://github.com/ariya/ask-llm/actions/workflows/test-novita.yml)
[![Test on Octo](https://github.com/ariya/ask-llm/actions/workflows/test-octo.yml/badge.svg)](https://github.com/ariya/ask-llm/actions/workflows/test-octo.yml)
[![Test on OpenAI](https://github.com/ariya/ask-llm/actions/workflows/test-openai.yml/badge.svg)](https://github.com/ariya/ask-llm/actions/workflows/test-openai.yml)
[![Test on OpenRouter](https://github.com/ariya/ask-llm/actions/workflows/test-openrouter.yml/badge.svg)](https://github.com/ariya/ask-llm/actions/workflows/test-openrouter.yml)
[![Test on Together](https://github.com/ariya/ask-llm/actions/workflows/test-together.yml/badge.svg)](https://github.com/ariya/ask-llm/actions/workflows/test-together.yml)

Supported LLM services include [AI21](https://studio.ai21.com), [Deep Infra](https://deepinfra.com), [DeepSeek](https://platform.deepseek.com/), [Fireworks](https://fireworks.ai), [Groq](https://groq.com), [Hyperbolic](https://www.hyperbolic.xyz), [Lepton](https://lepton.ai), [Mistral](https://console.mistral.ai), [Nebius](https://studio.nebius.ai), [Novita](https://novita.ai), [Octo](https://octo.ai), [OpenAI](https://platform.openai.com), [OpenRouter](https://openrouter.ai), and [Together](https://www.together.ai).

For configuration specifics, refer to the relevant section. The quality of answers can vary based on the model's performance.

* [AI21](https://studio.ai21.com)
```bash
export LLM_API_BASE_URL=https://api.ai21.com/studio/v1
export LLM_API_KEY="yourownapikey"
export LLM_CHAT_MODEL=jamba-1.5-mini
```

* [Deep Infra](https://deepinfra.com)
```bash
export LLM_API_BASE_URL=https://api.deepinfra.com/v1/openai
export LLM_API_KEY="yourownapikey"
export LLM_CHAT_MODEL="meta-llama/Meta-Llama-3.1-8B-Instruct"
```

* [DeepSeek](https://platform.deepseek.com)
```bash
export LLM_API_BASE_URL=https://api.deepseek.com/v1
export LLM_API_KEY="yourownapikey"
export LLM_CHAT_MODEL="deepseek-chat"
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

* [Mistral](https://console.mistral.ai)
```bash
export LLM_API_BASE_URL=https://api.mistral.ai/v1
export LLM_API_KEY="yourownapikey"
export LLM_CHAT_MODEL="open-mistral-7b"
```

* [Nebius](https://studio.nebius.ai)
```bash
export LLM_API_BASE_URL=https://api.studio.nebius.ai/v1
export LLM_API_KEY="yourownapikey"
export LLM_CHAT_MODEL="meta-llama/Meta-Llama-3.1-8B-Instruct"
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
