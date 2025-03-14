package my_webhooks

import (
	"encoding/json"
	"fmt"
	"github.com/MythicMeta/MythicContainer/logging"
	"github.com/MythicMeta/MythicContainer/mythicrpc"
	"github.com/MythicMeta/MythicContainer/webhookstructs"
)

func newCallbackWebhook(input webhookstructs.NewCallbackWebookMessage) {
	newMessage := webhookstructs.GetNewDefaultWebhookMessage()
	newMessage.Channel = webhookstructs.AllWebhookData.Get("my_webhooks").GetWebhookChannel(input, webhookstructs.WEBHOOK_TYPE_NEW_CALLBACK)
	var webhookURL = webhookstructs.AllWebhookData.Get("my_webhooks").GetWebhookURL(input, webhookstructs.WEBHOOK_TYPE_NEW_CALLBACK)
	if webhookURL == "" {
		logging.LogError(nil, "No webhook url specified for operation or locally", "data", newMessage)
		go mythicrpc.SendMythicRPCOperationEventLogCreate(mythicrpc.MythicRPCOperationEventLogCreateMessage{
			Message:      "No webhook url specified, can't send new callback webhook message",
			MessageLevel: mythicrpc.MESSAGE_LEVEL_INFO,
		})
		return
	}
	newMessage.Attachments[0].Title = "New Callback!"
	newMessage.Attachments[0].Color = "#b366ff"
	if newMessage.Attachments[0].Blocks != nil {
		(*newMessage.Attachments[0].Blocks)[0].Text.Text = fmt.Sprintf("You have a new callback!") // <!here> <-- add this if you want to ping the whole channel
	}
	// construct the fields list
	fieldsBlockStarter := []webhookstructs.SlackWebhookMessageAttachmentBlockText{
		{
			Type: "mrkdwn",
			Text: fmt.Sprintf("*Operation*\n%s", input.OperationName),
		},
	}
	fieldsBlockStarter = append(fieldsBlockStarter,
		webhookstructs.SlackWebhookMessageAttachmentBlockText{
			Type: "mrkdwn",
			Text: fmt.Sprintf("*Callback ID*\n%d", input.Data.DisplayID),
		})
	integrityLevelString := "MEDIUM"
	switch input.Data.IntegrityLevel {
	case 1:
		integrityLevelString = "LOW"
	case 2:
		integrityLevelString = "MEDIUM"
	case 3:
		integrityLevelString = "HIGH"
	case 4:
		integrityLevelString = "SYSTEM"
	}
	fieldsBlockStarter = append(fieldsBlockStarter,
		webhookstructs.SlackWebhookMessageAttachmentBlockText{
			Type: "mrkdwn",
			Text: fmt.Sprintf("*Integrity Level*\n%s", integrityLevelString),
		})
	ipArray := []string{}
	err := json.Unmarshal([]byte(input.Data.IPs), &ipArray)
	if err != nil {
		logging.LogError(err, "failed to unmarshal ip array")
		fieldsBlockStarter = append(fieldsBlockStarter,
			webhookstructs.SlackWebhookMessageAttachmentBlockText{
				Type: "mrkdwn",
				Text: fmt.Sprintf("*IP*\n%s", input.Data.IPs),
			})
	} else if len(ipArray) > 0 {
		fieldsBlockStarter = append(fieldsBlockStarter,
			webhookstructs.SlackWebhookMessageAttachmentBlockText{
				Type: "mrkdwn",
				Text: fmt.Sprintf("*IP*\n%s", ipArray[0]),
			})
	} else {
		fieldsBlockStarter = append(fieldsBlockStarter,
			webhookstructs.SlackWebhookMessageAttachmentBlockText{
				Type: "mrkdwn",
				Text: fmt.Sprintf("*IP*\n%s", input.Data.IPs),
			})
	}

	fieldsBlockStarter = append(fieldsBlockStarter,
		webhookstructs.SlackWebhookMessageAttachmentBlockText{
			Type: "mrkdwn",
			Text: fmt.Sprintf("*Type*\n%s", input.Data.AgentType),
		})
	fieldBlock := webhookstructs.SlackWebhookMessageAttachmentBlock{
		Type:   "section",
		Fields: &fieldsBlockStarter,
	}
	messageBlockText := webhookstructs.SlackWebhookMessageAttachmentBlockText{
		Type: "mrkdwn",
		Text: fmt.Sprintf("%s", input.Data.Description),
	}
	messageBlock := webhookstructs.SlackWebhookMessageAttachmentBlock{
		Type: "section",
		Text: &messageBlockText,
	}
	dividerBlock := webhookstructs.SlackWebhookMessageAttachmentBlock{
		Type: "divider",
	}
	// add the block to the blocks list
	tempBlockList := append(*(newMessage.Attachments[0].Blocks), fieldBlock, dividerBlock, messageBlock)
	newMessage.Attachments[0].Blocks = &tempBlockList
	// now actually send the message
	/*
		logging.LogDebug("webhook about to fire", "url", webhookURL, "message", newMessage)
		messageBytes, _ := json.MarshalIndent(newMessage, "", "  ")
		fmt.Printf("%s", string(messageBytes))

	*/

	webhookstructs.SubmitWebRequest("POST", webhookURL, newMessage)
}
