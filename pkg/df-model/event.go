package model

import (
	"fmt"
	"time"
)

// Event defines the model of a workflow event.
type Event struct {
	workflowName   string
	workflowExecID string
	workItemName   string
	workItemExecID string
	eventType      string
	state          string
	payload        string
	createTime     time.Time
	completeTime   time.Time
}

// GetPartitionKey returns the partition key for an event.
func (e *Event) GetPartitionKey() string {
	return fmt.Sprintf("%v-%v-%v", e.workflowName, e.workItemName, e.eventType)
}

// GetUniqueKey returns a unique key for an event.
func (e *Event) GetUniqueKey() string {
	return fmt.Sprintf("%v-%v", e.workflowExecID, e.workItemExecID)
}
