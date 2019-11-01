package handler

import (
	"log"
	"time"
	"todolist-facebook-chatbot/events"
)

type taskCreatedHandler struct {
	adminEmail string
	slackHook  string
}

func init() {
	log.Println("Init create notifier")
	notifier := taskCreatedHandler{
		adminEmail: "quangnguyen@blabla.com",
		slackHook:  "https://webhook.slack.com",
	}
	events.TaskCreated.Register(notifier)
}

func (t taskCreatedHandler) notifyAdmin(email string, time time.Time) {
	log.Printf("Notify admin: User created with email: %v\n", email)
}

func (t taskCreatedHandler) sendToSlack(email string, time time.Time) {
	log.Printf("Send to slack: User created with email: %v\n", email)
}

func (t taskCreatedHandler) Handle(payload events.TaskCreatedPayload) {
	t.sendToSlack(payload.Email, payload.Time)
	t.notifyAdmin(payload.Email, payload.Time)
}
