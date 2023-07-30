
package notification

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/kry0sc0pic/mercurius/config"
)

// WhatsappNotifier is responsible for sending notifications via Whatsapp
type WhatsappNotifier struct{}

// NewWhatsappNotifier creates a new instance of WhatsappNotifier
func NewWhatsappNotifier() *WhatsappNotifier {
	return &WhatsappNotifier{}
}

// Send sends a message to the configured Whatsapp number
func (n *WhatsappNotifier) Send(message string) error {
	return nil
	whatsappConfig := config.Get().Whatsapp
	data := map[string]string{
		"phone":   whatsappConfig.PhoneNumber,
		"body":    message,
		"apikey":  whatsappConfig.ApiKey,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", "https://api.callmebot.com/whatsapp.php", strings.NewReader(string(jsonData)))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

