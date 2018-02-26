package model

import (
	"fmt"
	"time"
)

// EventType defines the type of current event ,this type will be assigned when event created
type EventType string

const (
	WorkflowStartEvent    EventType = "WorkflowStart"
	WorkflowCompleteEvent EventType = "WorkflowComplete"
	WorkItemStartEvent    EventType = "WorkItemStart"
	WorkItemTimeoutEvent  EventType = "WorkItemTimeout"
	WorkItemCompleteEvent EventType = "WorkItemComplete"
)

// Event defines the model of a workflow event.
type Event struct {
	EventID        string
	WorkflowName   string
	WorkflowExecID string
	WorkItemName   string
	WorkItemExecID string
	EventType      EventType
	Payload        string
	CreateTime     time.Time
	CompleteTime   time.Time
}

// GetPartitionKey returns the partition key for an event.
func (e *Event) GetPartitionKey() string {
	if e.WorkItemName == "" {
		return fmt.Sprintf("%v-%v", e.WorkflowName, e.EventType)
	}

	return fmt.Sprintf("%v-%v-%v", e.WorkflowName, e.WorkItemName, e.EventType)
}

// GetUniqueKey returns a unique key for an event.
func (e *Event) GetUniqueKey() string {
	if e.WorkItemExecID == "" {
		return fmt.Sprintf("%v-%v", e.WorkflowExecID, e.EventID)
	}

	return fmt.Sprintf("%v-%v-%v", e.WorkflowExecID, e.WorkItemExecID, e.EventID)
}
