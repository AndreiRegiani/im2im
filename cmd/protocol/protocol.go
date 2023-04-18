package protocol

type Config struct {
	Bridges map[string]Bridge `yaml:"bridges"`
}

type Bridge struct {
	From struct {
		TCP         *TCP         `yaml:"tcp,omitempty"`
		TelegramBot *TelegramBot `yaml:"telegram_bot,omitempty"`
	} `yaml:"from"`
	To struct {
		TCP         *TCP         `yaml:"tcp,omitempty"`
		TelegramBot *TelegramBot `yaml:"telegram_bot,omitempty"`
	} `yaml:"to"`
}
