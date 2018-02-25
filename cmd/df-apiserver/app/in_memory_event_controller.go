package app

import (
	"github.com/yolo3301/dumb-flow/pkg/df-model"
)

// InMemoryEventController an in memory implementation of event controller.
type InMemoryEventController struct {
	events map[string][]model.Event
}

// NewInMemoryEventController new
func NewInMemoryEventController() InMemoryEventController {
	return InMemoryEventController{events: make(map[string][]model.Event)}
}

// SaveEvent not thread safe in this implementation
func (c InMemoryEventController) SaveEvent(event model.Event) error {
	key := event.GetPartitionKey()
	c.events[key] = append(c.events[key], event)
	return nil
}

// GetEvents from the map.
func (c InMemoryEventController) GetEvents(key string) ([]model.Event, error) {
	return c.events[key], nil
}
