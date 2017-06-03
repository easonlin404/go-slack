package slack

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestSendMessage(t *testing.T) {
	// You need get your slack WebhookURL from https://api.slack.com/incoming-webhooks
	webhookURL := os.Getenv("webhookURL")
	message := Message{
		Text: "message",
	}

	r, e := New().Debug(true).
		WebhookURL(webhookURL).
		SendMessage(message)

	assert.NoError(t, e)
	assert.Equal(t, 200, r.StatusCode)

}

func TestSendMessageWithAttachments(t *testing.T) {
	// You need get your slack WebhookURL from https://api.slack.com/incoming-webhooks
	webhookURL := os.Getenv("webhookURL")
	message := Message{
		Attachments: []Attachment{
			{Title: "title", MrkdwnIn: []string{"text"},
				Fields: []Field{
					{Title: "T1", Value: 0, Short: true},
					{Title: "T2", Value: 0, Short: true},
				},
			},
		},
	}

	r, e := New().Debug(true).
		WebhookURL(webhookURL).
		SendMessage(message)

	assert.NoError(t, e)
	assert.Equal(t, 200, r.StatusCode)

}
