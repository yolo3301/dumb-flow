package model

import (
	"fmt"
	"time"
)

// EventType defines the type of current event ,this type will be assigned when event created
type EventType string

const (
	start    EventType = "Start"
	timeout  EventType = "Timeout"
	complete EventType = "Complete"
)

// Event defines the model of a workflow event.
type Event struct {
	EventID        string
	WorkflowName   string
	WorkflowExecID string
	WorkItemName   string
	EventType      EventType // work start / timeout / complete (immutable)
	Payload        string
	CreateTime     time.Time
	CompleteTime   time.Time
}

// GetPartitionKey returns the partition key for an event.
func (e *Event) GetPartitionKey() string {
	// mapping to Kafka Topic
	return fmt.Sprintf("%v-%v-%v", e.WorkflowName, e.WorkItemName, e.EventType)
}

// GetUniqueKey returns a unique key for an event.
func (e *Event) GetUniqueKey() string {
	return fmt.Sprintf("%v-%v", e.WorkflowExecID, e.EventID)
}
