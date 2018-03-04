package model

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
	ID            string
	Name          string
	State         WorkflowState
	WorkItemExecs []WorkItemExec
}

type WorkItemExec struct {
	ID           string
	WorkItemName string
	State        string
}
