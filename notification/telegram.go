package notification

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/kry0sc0pic/mercurius/config"
)

// TelegramNotifier is responsible for sending notifications via Telegram
type TelegramNotifier struct{}

// NewTelegramNotifier creates a new instance of TelegramNotifier
func NewTelegramNotifier() *TelegramNotifier {
	return &TelegramNotifier{}
}

// Send sends a message to the configured Telegram chat
func (n *TelegramNotifier) Send(message string) error {
	return nil
	telegramConfig := config.Get().Telegram
	data := url.Values{}
	data.Set("chat_id", telegramConfig.ChatID)
	data.Set("text", message)

	req, err := http.NewRequest("POST", "https://api.telegram.org/bot"+telegramConfig.BotToken+"/sendMessage", strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

