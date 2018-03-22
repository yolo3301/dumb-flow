package app

import (
	"github.com/yolo3301/dumb-flow/pkg/df-model"
)

// QueueDAO defines interface for queue storage.
type QueueDAO interface {
	EnqueueNotification(notifications []model.Notification) ([]model.Notification, error)
	DequeueNotification(count int) ([]model.Notification, error)
	EnqueueEvents(topic string, events []model.Event) ([]model.Event, error)
	DequeueEvents(topic string, count int) ([]model.Event, error)

	Close()
	// Test only, remove later.
	SanityCheck() (string, error)
}
