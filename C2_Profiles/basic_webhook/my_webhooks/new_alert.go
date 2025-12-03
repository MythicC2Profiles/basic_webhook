package my_webhooks

import (
	"fmt"
	"time"

	"github.com/MythicMeta/MythicContainer/logging"
	"github.com/MythicMeta/MythicContainer/mythicrpc"
	"github.com/MythicMeta/MythicContainer/webhookstructs"
)

var throttleTime = 60 * time.Second
var newAlertLastTime = time.Now()

func newAlertMessage(input webhookstructs.NewAlertWebhookMessage) {
	newMessage := webhookstructs.GetNewDefaultWebhookMessage()
	newMessage.Channel = webhookstructs.AllWebhookData.Get("my_webhooks").GetWebhookChannel(input, webhookstructs.WEBHOOK_TYPE_NEW_ALERT)
	var webhookURL = webhookstructs.AllWebhookData.Get("my_webhooks").GetWebhookURL(input, webhookstructs.WEBHOOK_TYPE_NEW_ALERT)
	if time.Now().Sub(newAlertLastTime).Abs() <= throttleTime {
		logging.LogInfo("Not sending basic_webhook because <10s has passed since last message")
		return
	} else {
		newAlertLastTime = time.Now()
	}

	if webhookURL == "" {
		logging.LogError(nil, "No basic_webhook url specified for operation or locally", "data", newMessage)
		go mythicrpc.SendMythicRPCOperationEventLogCreate(mythicrpc.MythicRPCOperationEventLogCreateMessage{
			Message:      "No basic_webhook url specified, can't send alert basic_webhook message",
			MessageLevel: mythicrpc.MESSAGE_LEVEL_INFO,
		})
		return
	}

	newMessage.Attachments[0].Title = "New Event Alert!"
	newMessage.Attachments[0].Color = "#ff0000"
	if newMessage.Attachments[0].Blocks != nil {
		(*newMessage.Attachments[0].Blocks)[0].Text.Text = fmt.Sprintf("Source: %s", input.Data.Source)
	}

	// construct the fields list
	fieldsBlockStarter := []webhookstructs.SlackWebhookMessageAttachmentBlockText{
		{
			Type: "mrkdwn",
			Text: fmt.Sprintf("%s", input.Data.Message),
		},
	}
	fieldBlock := webhookstructs.SlackWebhookMessageAttachmentBlock{
		Type:   "section",
		Fields: &fieldsBlockStarter,
	}
	// add the block to the blocks list
	tempBlockList := append(*(newMessage.Attachments[0].Blocks), fieldBlock)
	newMessage.Attachments[0].Blocks = &tempBlockList
	// now actually send the message
	/*
		logging.LogDebug("basic_webhook about to fire", "url", webhookURL, "message", newMessage)
		messageBytes, _ := json.MarshalIndent(newMessage, "", "  ")
		fmt.Printf("%s", string(messageBytes))

	*/

	err := sendMessage(webhookURL, newMessage)
	if err != nil {
		logging.LogError(err, "failed to send webhook")
	}
}
