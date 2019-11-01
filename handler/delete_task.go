package handler

import (
	"log"
	"time"
	"todolist-facebook-chatbot/events"
)

type taskDeletedHandler struct {
	adminEmail string
	slackHook  string
}

func init() {
	log.Println("Init create notifier")
	notifier := taskDeletedHandler{
		adminEmail: "quangnguyen@blabla.com",
		slackHook:  "https://webhook.slack.com",
	}
	events.TaskDeleted.Register(notifier)
}

func (t taskDeletedHandler) notifyAdmin(email string, time time.Time) {
	log.Printf("Notify admin: User deleted with email: %v\n", email)
}

func (t taskDeletedHandler) sendToSlack(email string, time time.Time) {
	log.Printf("Send to slack: User deleted with email: %v\n", email)
}

func (t taskDeletedHandler) Handle(payload events.TaskDeletedPayload) {
	t.sendToSlack(payload.Email, payload.Time)
	t.notifyAdmin(payload.Email, payload.Time)
}
