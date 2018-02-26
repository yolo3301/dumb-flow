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

// EventState defines current status of this event
type EventState string

const (
	running     EventState = "Running"
	pending     EventState = "Pending"
	unavailable EventState = "Unavailable"
)

// Event defines the model of a workflow event.
type Event struct {
<<<<<<< Updated upstream
	EventID        string
	WorkflowName   string
	WorkflowExecID string
	WorkItemName   string
	EventType      string // work start / timeout / complete (immutable)
	Payload        string
	CreateTime     time.Time
	CompleteTime   time.Time
||||||| merged common ancestors
	workflowName   string
	workflowExecID string
	workItemName   string
	workItemExecID string
	eventType      string    // work start / timeout / complete (assigned when event created) 
	state          string    // running / pending / ... (changable)
	payload        string    // metadata - 
	createTime     time.Time 
	completeTime   time.Time
=======
<<<<<<< HEAD
	workflowName   string
	workflowExecID string
	workItemName   string
	workItemExecID string
	eventType      EventType  // work start / timeout / complete (assigned when event created)
	state          EventState // running / pending / ... (changable)
	payload        string     // metadata -
	createTime     time.Time
	completeTime   time.Time
||||||| merged common ancestors
	workflowName   string
	workflowExecID string
	workItemName   string
	workItemExecID string
	eventType      string    // work start / timeout / complete (assigned when event created) 
	state          string    // running / pending / ... (changable)
	payload        string    // metadata - 
	createTime     time.Time 
	completeTime   time.Time
=======
	EventID        string
	WorkflowName   string
	WorkflowExecID string
	WorkItemName   string
	EventType      string // work start / timeout / complete (immutable)
	Payload        string
	CreateTime     time.Time
	CompleteTime   time.Time
>>>>>>> c7abb6b91857149f6fd7a6d8731e2a15aaa8905f
>>>>>>> Stashed changes
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
