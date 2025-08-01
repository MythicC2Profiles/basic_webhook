package my_webhooks

import (
	"fmt"
	"github.com/MythicMeta/MythicContainer/webhookstructs"
)

const version = "0.0.2"

func Initialize() {
	myWebhooks := webhookstructs.WebhookDefinition{
		Name:                "MyBasicWebhooks",
		Description:         fmt.Sprintf("Basic webhook functionality for feedback, callbacks, alerts, and startup notifications.\nVersion: %s", version),
		NewFeedbackFunction: newfeedbackWebhook,
		NewCallbackFunction: newCallbackWebhook,
		NewStartupFunction:  newStartupMessage,
		NewAlertFunction:    newAlertMessage,
		NewCustomFunction:   newCustomMessage,
	}
	webhookstructs.AllWebhookData.Get("my_webhooks").AddWebhookDefinition(myWebhooks)
}
