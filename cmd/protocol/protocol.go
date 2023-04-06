package protocol

type Config struct {
	Bridges map[string]Bridge `yaml:"bridges"`
}

type Bridge struct {
	From struct {
		Netcat      *Netcat      `yaml:"netcat,omitempty"`
		TelegramBot *TelegramBot `yaml:"telegram_bot,omitempty"`
	} `yaml:"from"`
	To struct {
		Netcat      *Netcat      `yaml:"netcat,omitempty"`
		TelegramBot *TelegramBot `yaml:"telegram_bot,omitempty"`
	} `yaml:"to"`
}
