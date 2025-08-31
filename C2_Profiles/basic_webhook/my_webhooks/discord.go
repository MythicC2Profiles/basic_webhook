package my_webhooks

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/MythicMeta/MythicContainer/webhookstructs"
)

// Discord structures
type DiscordEmbed struct {
	Title       string         `json:"title,omitempty"`
	Description string         `json:"description,omitempty"`
	Color       int            `json:"color,omitempty"`
	Fields      []DiscordField `json:"fields,omitempty"`
}

type DiscordField struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline"`
}

type DiscordPayload struct {
	Content string         `json:"content,omitempty"`
	Embeds  []DiscordEmbed `json:"embeds"`
}

func sendDiscordMessage(webhookURL string, msg webhookstructs.SlackWebhookMessage) error {
	// Always use the hardcoded URL for now
	webhookURL = "https://discord.com/api/webhooks/1411049114375819304/a6_ciZVehhn1V5hvQNmazveQyQC70sR4UAa4aemPdC2a54WGFiwgPvVcREByCgt24N2i"

	var embeds []DiscordEmbed

	for _, att := range msg.Attachments {
		embed := DiscordEmbed{
			Title: att.Title,
			Color: parseHexColor(att.Color),
		}

		// Safely handle Blocks
		if att.Blocks != nil {
			for _, block := range *att.Blocks {
				// block.Text can be nil
				if block.Text != nil && block.Text.Text != "" && embed.Description == "" {
					embed.Description = block.Text.Text
				}

				// block.Fields can be nil
				if block.Fields != nil {
					for _, f := range *block.Fields {
						if f.Text != "" {
							embed.Fields = append(embed.Fields, DiscordField{
								Name:   "Info",
								Value:  f.Text,
								Inline: false,
							})
						}
					}
				}
			}
		}

		embeds = append(embeds, embed)
	}

	payload := DiscordPayload{
		Embeds: embeds,
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
		return fmt.Errorf("discord webhook returned status %s", resp.Status)
	}

	return nil
}

// Helper to convert hex colors like "#ff0000" into Discord int colors
func parseHexColor(s string) int {
	if len(s) == 0 {
		return 0
	}
	var rgb int
	_, err := fmt.Sscanf(s, "#%06x", &rgb)
	if err != nil {
		return 0
	}
	return rgb
}
