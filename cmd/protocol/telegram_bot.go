package protocol

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type TelegramBot struct {
	Token  string `yaml:"token"`
	ChatID int64  `yaml:"chat_id"`
}

func (tb *TelegramBot) InitTo(channel chan string) {
	client := http.Client{Timeout: 10 * time.Second}

	for {
		message := <-channel
		url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", tb.Token)

		data := map[string]interface{}{
			"chat_id": tb.ChatID,
			"text":    message,
		}

		jsonData, err := json.Marshal(data)
		if err != nil {
			log.Println("Error serialiazing JSON")
			return
		}

		resp, err := client.Post(url, "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			log.Println("Error posting to Telegram Bot API")
			return
		}

		if resp.StatusCode != 200 {
			log.Printf("Status code is not 200, response: %d", resp.StatusCode)
			return
		}
	}
}
