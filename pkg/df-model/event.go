package model

import (
	"time"
)

// EventType defines the type of current event ,this type will be assigned when event created
type EventType string

type EventState string

const (
	WorkflowStartEvent    EventType = "WorkflowStart"
	WorkflowCompleteEvent EventType = "WorkflowComplete"
	WorkItemStartEvent    EventType = "WorkItemStart"
	WorkItemCompleteEvent EventType = "WorkItemComplete"
	WorkItemFailEvent     EventType = "WorkItemFail"
)

const (
	EventCreated    EventState = "Created"
	EventScheduling EventState = "Scheduling"
	EventHandled    EventState = "Handled"
)

// Event defines the model of a workflow event.
type Event struct {
	EventID          string
	WorkflowName     string
	WorkflowExecID   string
	WorkItemName     string
	WorkItemExecID   string
	EventType        EventType
	LastModifiedTime time.Time
	State            EventState
	Payload          string
}
