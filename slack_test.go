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
		Attachments: []Attachment{
			{Title: "Order Status", MrkdwnIn: []string{"text"},
				Fields: []Field{
					{Title: "Request", Value: 0, Short: true},
					{Title: "Failed", Value: 0, Short: true},
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
