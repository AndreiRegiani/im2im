# im2im

Proof of Concept: daemon to bridge between two instant messaging protocols.

The service listens for incoming messages on a given protocol and forwards them to another protocol. It's designed to be extensible to support the addition of new protocols and configuring their usage via a configuration file (YAML), any amount of bridges can be run at the same time within the same process.

## Supported protocols

* TCP
* Telegram Bot

## Overview

Example for sending and receiving messages between TCP socket ⇔ Telegram Bot API, using a microcontroller (thin client) in LAN with a Raspberry Pi (running im2im):

![Alt text](./assets/overview.png "Overview")

### im2im.yaml

```yaml
bridges:
  bridge1:
      from:
        tcp:
          host: ""  # daemon listens on all interfaces (reachable in LAN)
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
      tcp:
          host: 192.168.1.50  # microcontroller IP address in LAN
          port: 9002
```

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
        tcp:
          host: localhost
          port: 9001
      to:
        tcp:
          host: localhost
          port: 9002
```

### Connecting

Using Netcat as TCP client and server.

```bash
nc -l 9002  # server
```

```bash
make run  # daemon
```

```bash
nc localhost 9001  # client
```

## Implementation

* Go 1.20, each bridge spawns two goroutines (sender/receiver) and uses one channel to pass the message.

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md) for submitting pull requests.

## License

MIT License
Copyright (c) 2023 Andrei Regiani
