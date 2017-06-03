package slack

import (
	"github.com/easonlin404/esrest"
	"net/http"
)

type Slack struct {
	webhookURL string
	debug      bool
}

func New() *Slack {
	return &Slack{}
}

type Message struct {
	//TODO: other
	Text        string       `json:"text"`
	Attachments []Attachment `json:"attachments"`
}

type Attachment struct {
	Title    string   `json:"title"`
	Text     string   `json:"text"`
	MrkdwnIn []string `json:"mrkdwn_in"`
	Fields   []Field  `json:"fields"`
}

type Field struct {
	Title string `json:"title"`
	Value int    `json:"value"`
	Short bool   `json:"short"`
}

func (s *Slack) WebhookURL(webhookURL string) *Slack {
	s.webhookURL = webhookURL
	return s
}

func (s *Slack) Debug(debug bool) *Slack {
	s.debug = debug
	return s
}

func (s *Slack) SendMessage(message Message) (*http.Response, error) {
	return esrest.New().
		Header("Content-Type", "text/plain").
		Debug(s.debug).
		Post(s.webhookURL).
		Body(message).
		Do()
}
