package slack

import (
	"github.com/nlopes/slack"
)

type Adapter interface {
	PushAlert(text string) error
}

type slackAPI struct {
	WebHookURL string
}

func NewSlackAPI(webhookURL string) Adapter {
	return &slackAPI{WebHookURL: webhookURL}
}

func (s *slackAPI) PushAlert(text string) error {
	slackMsg := &slack.WebhookMessage{
		Text: text,
	}
	return slack.PostWebhook(s.WebHookURL, slackMsg)
}
