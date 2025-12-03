package my_webhooks

import (
	"errors"
	"strings"

	"github.com/MythicMeta/MythicContainer/webhookstructs"
)

// sendMessage detects whether the webhook URL is for Slack or Discord and sends the message
func sendMessage(webhookURL string, newMessage webhookstructs.SlackWebhookMessage) error {
	if webhookURL == "" {
		return errors.New("no webhook URL provided")
	}

	// Detect Discord webhook URL
	if strings.Contains(webhookURL, "discord.com/api/webhooks") {
		return sendDiscordMessage(webhookURL, newMessage)
	}

	// Detect Slack webhook URL
	if strings.Contains(webhookURL, "hooks.slack.com") {
		_, _, err := webhookstructs.SubmitWebRequest("POST", webhookURL, newMessage)
		return err
	}

	// Unsupported webhook type
	return errors.New("unsupported webhook type: currently only Slack and Discord are supported")
}
