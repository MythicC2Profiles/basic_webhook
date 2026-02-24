package my_webhooks

import (
	"errors"
	"strings"

	"github.com/MythicMeta/MythicContainer/webhookstructs"
)

// sendMessage detects whether the basic_webhook URL is for Slack or Discord and sends the message
func sendMessage(webhookURL string, newMessage webhookstructs.SlackWebhookMessage) error {
	if webhookURL == "" {
		return errors.New("no basic_webhook URL provided")
	}

	// Detect Discord basic_webhook URL
	if strings.Contains(webhookURL, "discord.com") {
		return sendDiscordMessage(webhookURL, newMessage)
	}
    
	// Detect Google Chat basic_webhook URL
    if strings.Contains(webhookURL, "chat.googleapis.com") {
		return sendGoogleChatMessage(webhookURL, newMessage)
	}

	// Detect Slack basic_webhook URL
	if strings.Contains(webhookURL, "slack.com") {
		_, _, err := webhookstructs.SubmitWebRequest("POST", webhookURL, newMessage)
		return err
	}

	// Unsupported basic_webhook type
	return errors.New("unsupported basic_webhook type: currently only Slack and Discord are supported")
}
