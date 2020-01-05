package events

import (
	"time"
)

var TaskCreated taskCreated

type TaskCreatedHandler interface {
	Handle(TaskCreatedPayload)
}

// TaskCreatedPayload is the data for when a user is created
type TaskCreatedPayload struct {
	Email string
	Time  time.Time
}

type taskCreated struct {
	handlers []TaskCreatedHandler
}

// Register adds an event handler for this event
func (t *taskCreated) Register(handler TaskCreatedHandler) {
	t.handlers = append(t.handlers, handler)
}

// Trigger sends out an event with the payload
func (t *taskCreated) Trigger(payload TaskCreatedPayload) {
	for _, handler := range t.handlers {
		go handler.Handle(payload)
	}
}
