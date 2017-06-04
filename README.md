# go-slack

A simple, flexible Golang wrapper around the [Slack webhook API](https://api.slack.com). Makes it easy to send notifications to Slack from your application(inspired by nodejs [slack-notify](https://github.com/andrewchilds/slack-notify)).

[![Travis branch](https://img.shields.io/travis/easonlin404/go-slack/master.svg)](https://travis-ci.org/easonlin404/go-slack)
[![Codecov branch](https://img.shields.io/codecov/c/github/easonlin404/go-slack/master.svg)](https://codecov.io/gh/easonlin404/go-slack)

## Features
* Send plain text and message with attachments
* Send attachments with markdown syntax 
* Debug
* Timeout
* Logger
* [Context](https://golang.org/pkg/context/)(todo)

## Installation
```sh
$ go get -u github.com/easonlin404/go-slack
```

## Useage
Basic send text message
```go
// You need get your slack WebhookURL from https://api.slack.com/incoming-webhooks
	webhookURL := "https://hooks.slack.com/services/T00000000/B00000000/XXXXXXXXXXXXXXXXXXXXXXXX"
  
	message := Message{
		Text: "message",
	}

	r, e := slack.New().
		WebhookURL(webhookURL).
		SendMessage(message)
```
Send message with attachments
```go
// You need get your slack WebhookURL from https://api.slack.com/incoming-webhooks
	webhookURL := "https://hooks.slack.com/services/T00000000/B00000000/XXXXXXXXXXXXXXXXXXXXXXXX"
  
	message := Message{
		Attachments: []Attachment{
			{Title: "title",
				Fields: []Field{
					{Title: "T1", Value: 0, Short: true},
					{Title: "T2", Value: 0, Short: true},
				},
			},
		},
	}

	r, e := slack.New().
		WebhookURL(webhookURL).
		SendMessage(message)

```
Attachments enable markdown

```go
// You need get your slack WebhookURL from https://api.slack.com/incoming-webhooks
	webhookURL := "https://hooks.slack.com/services/T00000000/B00000000/XXXXXXXXXXXXXXXXXXXXXXXX"
  
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

	r, e := slack.New().
		WebhookURL(webhookURL).
		SendMessage(message)

```

Debug
```go
// You need get your slack WebhookURL from https://api.slack.com/incoming-webhooks
	webhookURL := "https://hooks.slack.com/services/T00000000/B00000000/XXXXXXXXXXXXXXXXXXXXXXXX"
  
	message := Message{
		Text: "message",
	}

	r, e := slack.New().Debug(true).
		WebhookURL(webhookURL).
		SendMessage(message)
