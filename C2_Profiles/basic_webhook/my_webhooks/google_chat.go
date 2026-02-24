import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"github.com/MythicMeta/MythicContainer/webhookstructs"
)
type GoogleChatMessage struct {
	Text string `json:"text,omitempty"`
}
func sendGoogleChatMessage(webhookURL string, msg webhookstructs.SlackWebhookMessage) error {
	var builder strings.Builder
	for _, att := range msg.Attachments {
		if att.Title != "" {
			builder.WriteString(fmt.Sprintf("*%s*\n", att.Title))
		}
		if att.Blocks != nil {
			for _, block := range *att.Blocks {
				if block.Text != nil && block.Text.Text != "" {
					builder.WriteString(fmt.Sprintf("%s\n", block.Text.Text))
				}
				if block.Fields != nil {
					for _, f := range *block.Fields {
						if f.Text != "" {
							builder.WriteString(fmt.Sprintf("%s\n", f.Text))
						}
					}
				}
			}
		}
		builder.WriteString("\n")
	}
	payload := GoogleChatMessage{
		Text: builder.String(),
	}
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("google chat webhook returned status %s", resp.Status)
	}
	return nil
}
