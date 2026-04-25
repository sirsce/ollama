# Ollama

> A fork of [ollama/ollama](https://github.com/ollama/ollama) — Get up and running with large language models locally.

[![Latest Release](https://github.com/ollama/ollama/actions/workflows/latest.yaml/badge.svg)](https://github.com/ollama/ollama/actions/workflows/latest.yaml)

## Overview

Ollama allows you to run large language models (LLMs) locally on your machine. It provides a simple API for running and managing models, and a library of pre-built models ready to use.

## Features

- Run LLMs locally with a simple command
- REST API for integration with your applications
- Support for a wide range of models (Llama 3, Mistral, Gemma, and more)
- Cross-platform: macOS, Linux, and Windows
- GPU acceleration support (NVIDIA CUDA, AMD ROCm, Apple Metal)
- OpenAI-compatible API endpoint

## Quick Start

### Install

**macOS / Linux:**
```bash
curl -fsSL https://ollama.com/install.sh | sh
```

**Windows:**

Download the installer from the [releases page](../../releases).

**Docker:**
```bash
docker pull ollama/ollama
docker run -d -v ollama:/root/.ollama -p 11434:11434 --name ollama ollama/ollama
```

### Run a Model

```bash
ollama run llama3
```

## API Usage

Once Ollama is running, you can interact with it via the REST API:

```bash
# Generate a response
curl http://localhost:11434/api/generate -d '{
  "model": "llama3",
  "prompt": "Why is the sky blue?"
}'

# Chat
curl http://localhost:11434/api/chat -d '{
  "model": "llama3",
  "messages": [
    { "role": "user", "content": "Hello!" }
  ]
}'
```

## Building from Source

### Prerequisites

- Go 1.22+
- CMake 3.24+
- GCC / Clang

### Build

```bash
git clone https://github.com/ollama/ollama.git
cd ollama
go build .
```

### Build with GPU Support

**NVIDIA CUDA:**
```bash
cmake -B build -DGGML_CUDA=ON
cmake --build build --config Release
```

**AMD ROCm:**
```bash
cmake -B build -DGGML_HIPBLAS=ON
cmake --build build --config Release
```

## Model Library

Ollama supports a growing library of models. See the full list at [ollama.com/library](https://ollama.com/library).

| Model | Parameters | Size |
|-------|-----------|------|
| Llama 3 | 8B | 4.7GB |
| Llama 3 | 70B | 40GB |
| Mistral | 7B | 4.1GB |
| Gemma 2 | 9B | 5.5GB |
| Phi-3 | 3.8B | 2.3GB |
| Qwen2.5 | 7B | 4.4GB |
| Qwen2.5 Coder | 7B | 4.4GB |
| DeepSeek-Coder V2 | 16B | 8.9GB |

> **Personal note:** I've been primarily using `qwen2.5-coder:7b` and `deepseek-coder-v2` for local coding assistance. Qwen2.5 Coder has noticeably better performance on code-related tasks compared to the base Qwen2.5 model. For general chat I still reach for `llama3`.
>
> **Tip:** I run Ollama behind a small nginx reverse proxy on my home server so I can reach it from other devices on the LAN without exposing the port directly. See [this gist](https://gist.github.com) for the config I use.
>
> **Note:** I also set `OLLAMA_KEEP_ALIVE=24h` in my systemd service so models stay loaded in memory overnight — makes repeated use much snappier without having to wait for the model to reload each time. I also set `OLLAMA_NUM_PARALLEL=2` so I can handle requests from a couple of devices simultaneously without them queuing up.
