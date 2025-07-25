# whats-in-it

A CLI utility that analyzes the content of a file and answers the question “What’s in it?” in a single, concise sentence. Supports multiple AI backends and multiple output languages. The interface-driven design and clean-code approach ensure that integrating further LLMs is both straightforward and maintainable.

![demonstration gif](assets/demo.gif)

## Features

- **Multi-backend support**: Easily switch between GigaChat and YandexGPT models.
- **Config-driven**: All settings managed via a YAML config file.
- **Robust error handling**: Gracefully handles missing files, empty content, and API errors.
- **Structured logging**: JSON-formatted logs for easy machine parsing.

---

## Table of Contents

1. [Prerequisites](#prerequisites)
2. [Installation](#installation)
3. [Configuration](#configuration)
4. [Usage](#usage)
5. [Examples](#examples)

---

## Prerequisites

- Go 1.20 or higher
- Valid API credentials for at least one backend:
  - **GigaChat**: Auth Key from [Sberbank Studio](https://developers.sber.ru/studio/registration)
  - **YandexGPT**: OAuth token from [Yandex Cloud](https://yandex.cloud/en/docs/iam/concepts/authorization/oauth-token)

---

## Installation

This project provides both a Makefile target and an install script for easy setup.

### Using the installer script

1. Ensure you have root permissions (or use `sudo`).
2. Run:
   ```bash
   sudo bash install.sh
   ```
   This will:
   - Create `/etc/whats-in-it` and `/etc/whats-in-it/certs`.
   - Download the required TLS certificates for GigaChat into `/etc/whats-in-it/certs/{first.pem,second.pem}`.
   - Build the `wii` binary into `/usr/local/bin/wii`.
   - Copy the default config file to `/etc/whats-in-it/config.yaml` and set permissions.

### Using Makefile

You can also install or uninstall via the Makefile:

- To install:
  ```bash
  make install
  ```
- To uninstall:
  ```bash
  make uninstall
  ```

Both targets invoke the respective `install.sh` and `uninstall.sh` scripts under the hood.

## Configuration

The utility reads settings from `/etc/whats-in-it/config.yaml`. A sample config:

```yaml
language: "english"
current_model: "giga_chat"

giga_chat:
  model: "GigaChat-2"
  scope: "GIGACHAT_API_PERS"
  token_endpoint: "https://ngw.devices.sberbank.ru:9443/api/v2/oauth"
  chat_endpoint: "https://gigachat.devices.sberbank.ru/api/v1/chat/completions"
  auth_key: "your_auth_key"

yandex_gpt:
  model_uri: "gpt://<your_model_uri>/yandexgpt"
  token: "your_auth_token"
  token_endpoint: "https://iam.api.cloud.yandex.net/iam/v1/tokens"
  chat_endpoint: "https://llm.api.cloud.yandex.net/foundationModels/v1/completion"
```

---

## Usage

```bash
wii [FILE]
```

- If the file does not exist, you will see an error message:

  ```bash
  wii: example.txt: No such file or directory
  ```

- If the file is empty, you get a localized warning:

  ```bash
  File is empty.
  ```

Example output:

```bash
$ wii story.txt
This file contains a short sci-fi story about time travel.
```

---

## Examples

```bash
# Analyze in German
vim /etc/whats-in-it/config.yaml  # set language: "deutsch"
wii document.txt
# Output: "Die Datei enthält eine vollständige Anleitung zur Installation von Docker."
```

---
