package protocol

type TelegramBot struct {
	Token  string `yaml:"token"`
	ChatID int64  `yaml:"chat_id"`
}
