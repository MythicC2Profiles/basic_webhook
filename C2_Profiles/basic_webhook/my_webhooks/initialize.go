package my_webhooks

import (
	"github.com/MythicMeta/MythicContainer/webhookstructs"
)

func Initialize() {
	myWebhooks := webhookstructs.WebhookDefinition{
		Name:                "MyBasicWebhooks",
		Description:         "Basic webhook functionality for feedback, callbacks, alerts, and startup notifications",
		NewFeedbackFunction: newfeedbackWebhook,
		NewCallbackFunction: newCallbackWebhook,
		NewStartupFunction:  newStartupMessage,
		NewAlertFunction:    newAlertMessage,
		NewCustomFunction:   newCustomMessage,
	}
	webhookstructs.AllWebhookData.Get("my_webhooks").AddWebhookDefinition(myWebhooks)
}
