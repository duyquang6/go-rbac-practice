package events

import (
	"time"
)

var TaskDeleted taskDeleted

type TaskDeletedHandler interface {
	Handle(TaskDeletedPayload)
}

// TaskDeletedPayload is the data for when a user is Deleted
type TaskDeletedPayload struct {
	Email string
	Time  time.Time
}

type taskDeleted struct {
	handlers []TaskDeletedHandler
}

// Register adds an event handler for this event
func (t *taskDeleted) Register(handler TaskDeletedHandler) {
	t.handlers = append(t.handlers, handler)
}

// Trigger sends out an event with the payload
func (t taskDeleted) Trigger(payload TaskDeletedPayload) {
	for _, handler := range t.handlers {
		go handler.Handle(payload)
	}
}
