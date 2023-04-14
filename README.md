# im2im

Proof of Concept: daemon to bridge between two instant messaging protocols.

The service listens for incoming messages on a given protocol and forwards them to another protocol. It's designed to be extensible to support the addition of new protocols and configuring their usage via a configuration file (YAML).

## Overview

Example for sending and receiving messages between netcat ⇔ Telegram Bot API.

![Alt text](./assets/overview.png "Overview")

### im2im.yaml

```yaml
bridges:
  bridge1:
      from:
        netcat:
          port: 9001
      to:
        telegram_bot:
          token: ABCDEFGHIJKLMNOPQRSTUVWXYZ
          chat_id: 123456789
  bridge2:
    from:
      telegram_bot:
          token: ABCDEFGHIJKLMNOPQRSTUVWXYZ
          chat_id: 123456789
    to:
      netcat:
          port: 9002
```

## Implementation

* Go 1.20, each bridge spawns two goroutines (sender/receiver) and uses one channel to relay the message.

## Installation

### From source

```bash
make build
```

### Docker

```bash
make docker-build
```

## Usage

### Configuration

```bash
cp im2im.yaml.example im2im.yaml
```

Minimal example:

```yaml
bridges:
  bridge0:
      from:
        netcat:
          port: 9001
      to:
        netcat:
          port: 9002
```

### Connecting

```bash
nc -l 9002
```

```bash
go run cmd/im2im.go
```

```bash
nc localhost 9001
```

## Supported protocols

* netcat (TCP socket)
* Telegram Bot
* ...

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md) for submitting pull requests.

## License

MIT License
Copyright (c) 2023 Andrei Regiani
