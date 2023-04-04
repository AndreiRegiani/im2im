# im2im

Proof of Concept: deamon to bridge between two instant messaging protocols.

The service listens for incoming messages on a given protocol and forwards them to another protocol. It's designed to be extensible to support the addition of new protocols and configuring their usage via a configuration file.

![Alt text](./assets/overview.png "Overview")

## Implementation

* Go 1.20

## Installation

### From source

```bash
```

### Docker

```bash
```

## Usage

### Configuration

```bash
cp im2im.yaml.example im2im.yaml
```

### im2im.yaml

```yaml
```

## Supported protocols

* Telnet (`telnet`)
* Telegram Bot (`telegram-bot`)

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md) for submitting pull requests.

## License

MIT License
Copyright (c) 2023 Andrei Regiani
