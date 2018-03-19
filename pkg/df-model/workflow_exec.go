package model

import (
	"fmt"
)

const WorkflowExecPrefix = "wfexec"
const WorkItemExecPrefix = "wiexec"

type WorkflowState string

const (
	WorkflowPending    WorkflowState = "Pending"
	WorkflowInProgress WorkflowState = "InProgress"
	WorkflowPaused     WorkflowState = "Paused"
	WorkflowAborted    WorkflowState = "Aborted"
	WorkflowError      WorkflowState = "Error"
	WorkflowTimedout   WorkflowState = "Timedout"
	WorkflowCompleted  WorkflowState = "Completed"
)

type WorkflowExec struct {
	ID           string        `json:"id,omitempty"`
	WorkflowName string        `json:"workflowName,omitempty"`
	State        WorkflowState `json:"state,omitempty"`
}

type WorkItemExec struct {
	ID             string `json:"id,omitempty"`
	WorkflowName   string `json:"workflowName,omitempty"`
	WorkflowExecID string `json:"workflowExecId,omitempty"`
	WorkItemName   string `json:"workItemName,omitempty"`
	State          string `json:"state,omitempty"`
}

func (exec *WorkflowExec) Key() string {
	return fmt.Sprintf("%v-%v-%v", WorkflowExecPrefix, exec.WorkflowName, exec.ID)
}

func (exec *WorkItemExec) Key() string {
	return fmt.Sprintf("%v-%v-%v-%v-%v", WorkItemExecPrefix, exec.WorkflowName, exec.WorkflowExecID, exec.WorkItemName, exec.ID)
}
