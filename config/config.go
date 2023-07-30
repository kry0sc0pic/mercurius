
package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Telegram struct {
		BotToken string
		ChatID   string
	}
	Whatsapp struct {
		ApiKey      string
		PhoneNumber string
	}
	Discord struct {
		WebhookURL string
	}
}

var config Config

// LoadConfig loads configuration from .env file
func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.Telegram.BotToken = os.Getenv("TELEGRAM_BOT_TOKEN")
	config.Telegram.ChatID = os.Getenv("TELEGRAM_CHAT_ID")

	config.Whatsapp.ApiKey = os.Getenv("WHATSAPP_API_KEY")
	config.Whatsapp.PhoneNumber = os.Getenv("WHATSAPP_PHONE_NUMBER")

	config.Discord.WebhookURL = os.Getenv("DISCORD_WEBHOOK_URL")
}

// Get returns the loaded configuration
func Get() *Config {
	return &config
}

