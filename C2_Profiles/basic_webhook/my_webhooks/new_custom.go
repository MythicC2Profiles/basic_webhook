package my_webhooks

import (
	"fmt"
	"github.com/MythicMeta/MythicContainer/logging"
	"github.com/MythicMeta/MythicContainer/mythicrpc"
	"github.com/MythicMeta/MythicContainer/webhookstructs"
)

func newCustomMessage(input webhookstructs.NewCustomWebhookMessage) {
	newMessage := webhookstructs.GetNewDefaultWebhookMessage()
	newMessage.Channel = webhookstructs.AllWebhookData.Get("my_webhooks").GetWebhookChannel(input, webhookstructs.WEBHOOK_TYPE_NEW_CUSTOM)
	var webhookURL = webhookstructs.AllWebhookData.Get("my_webhooks").GetWebhookURL(input, webhookstructs.WEBHOOK_TYPE_NEW_CUSTOM)
	if webhookURL == "" {
		logging.LogError(nil, "No webhook url specified for operation or locally", "data", newMessage)
		go mythicrpc.SendMythicRPCOperationEventLogCreate(mythicrpc.MythicRPCOperationEventLogCreateMessage{
			Message:      "No webhook url specified, can't send custom webhook message",
			MessageLevel: mythicrpc.MESSAGE_LEVEL_INFO,
		})
		return
	}

	newMessage.Attachments[0].Title = fmt.Sprintf("%s Message!", input.OperatorUsername)
	newMessage.Attachments[0].Color = "#ff0000"
	// construct the fields list
	blockPieces := []webhookstructs.SlackWebhookMessageAttachmentBlockText{}
	for key, val := range input.Data {
		blockPieces = append(blockPieces, webhookstructs.SlackWebhookMessageAttachmentBlockText{
			Type: "mrkdwn",
			Text: fmt.Sprintf("*%s*\n%s", key, val),
		})
	}
	fieldBlock := webhookstructs.SlackWebhookMessageAttachmentBlock{
		Type:   "section",
		Fields: &blockPieces,
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
