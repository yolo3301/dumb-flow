package app

import (
	"fmt"

	model "github.com/yolo3301/dumb-flow/pkg/df-model"
)

type DefaultQueueDAO struct {
}

func (dao DefaultQueueDAO) EnqueueNotification(notifications []model.Notification) ([]model.Notification, error) {
	return nil, fmt.Errorf("Not implemented")
}

func (dao DefaultQueueDAO) DequeueNotification(count int) ([]model.Notification, error) {
	return nil, fmt.Errorf("Not implemented")
}

func (dao DefaultQueueDAO) EnqueueEvents(topic string, events []model.Event) ([]model.Event, error) {
	return nil, fmt.Errorf("Not implemented")
}

func (dao DefaultQueueDAO) DequeueEvents(topic string, events []model.Event) ([]model.Event, error) {
	return nil, fmt.Errorf("Not implemented")
}
