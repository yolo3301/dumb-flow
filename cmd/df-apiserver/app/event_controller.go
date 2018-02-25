package app

import (
	"github.com/yolo3301/dumb-flow/pkg/df-model"
)

// EventController is the interface for event controller.
type EventController interface {
	SaveEvent(event model.Event) error
	GetEvents(key string) ([]model.Event, error)
}
