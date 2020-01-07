package handler

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"log"
	"time"
	"todolist-facebook-chatbot/adapters/slack"
	"todolist-facebook-chatbot/events"
	"todolist-facebook-chatbot/services"
)

type taskCreatedHandler struct {
	slackAdapter slack.Adapter
}

func init() {
	log.Println("Init create notifier")
	notifier := taskCreatedHandler{}
	_ = services.GetServiceContainer().Invoke(func(slackAdapter slack.Adapter) {
		notifier.slackAdapter = slackAdapter
	})
	events.TaskCreated.Register(notifier)
}

func (t taskCreatedHandler) notifyAdmin(email string, time time.Time) {
	log.Printf("Notify admin: User created with email: %v\n", email)
}

func (t taskCreatedHandler) sendToSlack(email string, time time.Time) {
	log.Printf("Send to slack: User created with email: %v\n", email)
	err := t.slackAdapter.PushAlert(fmt.Sprintf("User created with email: %v", email))
	if err != nil {
		logrus.Errorf("Push msg alert error %v", err)
	}
}

func (t taskCreatedHandler) Handle(payload events.TaskCreatedPayload) {
	t.sendToSlack(payload.Email, payload.Time)
	t.notifyAdmin(payload.Email, payload.Time)
}
