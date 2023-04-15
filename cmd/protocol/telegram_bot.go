package protocol

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
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
	Message  TelegramItemMessage `json:"message"`
	UpdateID int64               `json:"update_id"`
}

type TelegramItemMessage struct {
	Text string           `json:"text"`
	Chat TelegramItemChat `json:"chat"`
}

type TelegramItemChat struct {
	ChatID int64 `json:"id"`
}

func (tb *TelegramBot) InitFrom(channel chan string) {
	var offset int64

	for {
		tb.fromPolling(channel, &offset)
		time.Sleep(5 * time.Second)
	}
}

func (tb *TelegramBot) fromPolling(channel chan string, offset *int64) {
	log.Println("Polling")

	client := http.Client{Timeout: 10 * time.Second}
	url := fmt.Sprintf("https://api.telegram.org/bot%s/getUpdates", tb.Token)

	if *offset != 0 {
		url = fmt.Sprintf("%s?offset=%d", url, *offset)
	}

	resp, err := client.Get(url)
	if err != nil {
		log.Println("TelegramBot: InitFrom: error fetching getUpdates")
	}

	if resp.StatusCode != http.StatusOK {
		log.Printf("TelegramBot: InitFrom: status code: %d", resp.StatusCode)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("TelegramBot: InitFrom: error reading response body")
		return
	}

	result := TelegramResult{}

	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Println("TelegramBot: InitFrom: error unmarshaling JSON")
		return
	}

	var newOffset int64

	for _, item := range result.Result {
		message := item.Message.Text
		log.Printf("TelegramBot: from: \"%s\"", message)
		newOffset = item.UpdateID + 1

		// Prevent re-sending the historic of messages that happens when first
		// polling without an offset
		if *offset != 0 {
			channel <- message
		}
	}

	*offset = newOffset
}

func (tb *TelegramBot) InitTo(channel chan string) {
	client := http.Client{Timeout: 10 * time.Second}

	for {
		message := <-channel
		log.Printf("TelegramBot: to: \"%s\"", message)

		// Telegram Bot API doesn't allow empty messages
		if strings.TrimSpace(message) == "" {
			continue
		}

		data := map[string]interface{}{
			"chat_id": tb.ChatID,
			"text":    message,
		}

		jsonData, err := json.Marshal(data)
		if err != nil {
			log.Println("TelegramBot: InitTo: error serialiazing JSON")
			return
		}

		url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", tb.Token)

		resp, err := client.Post(url, "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			log.Println("TelegramBot: InitTo: error posting to Telegram Bot API")
			return
		}

		if resp.StatusCode != http.StatusOK {
			log.Printf("TelegramBot: InitTo: status code: %d", resp.StatusCode)
			return
		}
	}
}
