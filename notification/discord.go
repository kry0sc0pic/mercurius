package notification

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/kry0sc0pic/mercurius/config"
)

// DiscordNotifier is responsible for sending notifications via Discord
type DiscordNotifier struct{}

// NewDiscordNotifier creates a new instance of DiscordNotifier
func NewDiscordNotifier() *DiscordNotifier {
	return &DiscordNotifier{}
}

// Send sends a message to the configured Discord channel
func (n *DiscordNotifier) Send(message string) error {
	discordConfig := config.Get().Discord
	data := map[string]string{
		"content": message,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", discordConfig.WebhookURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

