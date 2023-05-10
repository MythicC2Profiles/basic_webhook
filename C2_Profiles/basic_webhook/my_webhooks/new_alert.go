package my_webhooks

import (
	"fmt"
	"github.com/MythicMeta/MythicContainer/logging"
	"github.com/MythicMeta/MythicContainer/webhookstructs"
)

func newAlertMessage(input webhookstructs.NewAlertWebhookMessage) {
	newMessage := webhookstructs.GetNewDefaultWebhookMessage()
	newMessage.Channel = webhookstructs.AllWebhookData.Get("my_webhooks").GetWebhookChannel(input, webhookstructs.WEBHOOK_TYPE_NEW_ALERT)
	var webhookURL = webhookstructs.AllWebhookData.Get("my_webhooks").GetWebhookURL(input, webhookstructs.WEBHOOK_TYPE_NEW_ALERT)
	if webhookURL == "" {
		logging.LogError(nil, "No webhook url specified for operation or locally", "data", newMessage)
		//go mythicrpc.SendMythicRPCOperationEventLogCreate(mythicrpc.MythicRPCOperationEventLogCreateMessage{
		//	Message:      "No webhook url specified, can't send webhook message",
		//	MessageLevel: mythicrpc.MESSAGE_LEVEL_WARNING,
		//})
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
		logging.LogDebug("webhook about to fire", "url", webhookURL, "message", newMessage)
		messageBytes, _ := json.MarshalIndent(newMessage, "", "  ")
		fmt.Printf("%s", string(messageBytes))

	*/

	webhookstructs.SubmitWebRequest("POST", webhookURL, newMessage)
}
