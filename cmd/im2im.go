package main

import (
	"flag"
	"log"
	"os"

	"github.com/AndreiRegiani/im2im/cmd/protocol"
	"gopkg.in/yaml.v3"
)

func loadConfig() protocol.Config {
	configFlag := flag.String("config", "im2im.yaml", "Path to the config YAML file")
	flag.Parse()

	data, err := os.ReadFile(*configFlag)
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	var config protocol.Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("Error unmarshaling YAML: %v", err)
	}

	return config
}

func runBridges() {
	config := loadConfig()

	for label, bridge := range config.Bridges {
		log.Printf("* Bridge: %s", label)

		channel := make(chan string)

		if bridge.From.Netcat != nil {
			go bridge.From.Netcat.InitFrom(channel)
		}

		if bridge.To.Netcat != nil {
			go bridge.To.Netcat.InitTo(channel)
		}

		if bridge.From.TelegramBot != nil {
			go bridge.From.TelegramBot.InitFrom(channel)
		}

		if bridge.To.TelegramBot != nil {
			go bridge.To.TelegramBot.InitTo(channel)
		}
	}

	select {}
}

func main() {
	runBridges()
}
