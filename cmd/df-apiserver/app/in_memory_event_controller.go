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

// TestEventController dumb.
func (c InMemoryEventController) TestEventController() string {
	return "ok"
}
