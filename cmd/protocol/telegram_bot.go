package protocol

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type TelegramBot struct {
	Token  string `yaml:"token"`
	ChatID int64  `yaml:"chat_id"`
}

type TelegramResult struct {
	Ok     bool           `json:"ok"`
	Result []TelegramItem `json:"result"`
}

type TelegramItem struct {
	Message TelegramItemMessage `json:"message"`
}

type TelegramItemMessage struct {
	Text string           `json:"text"`
	Chat TelegramItemChat `json:"chat"`
}

type TelegramItemChat struct {
	ChatID int64 `json:"id"`
}

func (tb *TelegramBot) InitFrom(channel chan string) {
	for {
		tb.polling(channel)
		time.Sleep(5 * time.Second)
	}
}

func (tb *TelegramBot) polling(channel chan string) {
	log.Println("Polling")

	client := http.Client{Timeout: 10 * time.Second}
	url := fmt.Sprintf("https://api.telegram.org/bot%s/getUpdates", tb.Token)

	resp, err := client.Get(url)
	if err != nil {
		log.Println("Error getting getUpdates")
	}

	if resp.StatusCode != http.StatusOK {
		log.Printf("Status code is not 200, response: %d", resp.StatusCode)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading res.Body")
		return
	}

	result := TelegramResult{}

	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Println("Error unmarshaling JSON")
		return
	}

	for _, item := range result.Result {
		message := item.Message.Text
		log.Printf("Telegram: to: \"%s\"", message)
		channel <- message
	}
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

		if resp.StatusCode != http.StatusOK {
			log.Printf("Status code is not 200, response: %d", resp.StatusCode)
			return
		}

	}
}
