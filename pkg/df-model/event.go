package model

import (
	"fmt"
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
	ID               string     `json:"id,omitempty"`
	WorkflowName     string     `json:"workflowName,omitempty"`
	WorkflowExecID   string     `json:"workflowExecId,omitempty"`
	WorkItemName     string     `json:"workItemName,omitempty"`
	WorkItemExecID   string     `json:"workItemExecId,omitempty"`
	EventType        EventType  `json:"eventTyep,omitempty"`
	LastModifiedTime time.Time  `json:"lastModifiedTime,omitempty"`
	State            EventState `json:"State,omitempty"`
	Payload          string     `json:"payload,omitempty"`
}

func (e *Event) Key() string {
	return fmt.Sprintf("%v-%v-%v-%v-%v", e.WorkflowName, e.WorkflowExecID, e.WorkItemName, e.WorkItemExecID, e.ID)
}
