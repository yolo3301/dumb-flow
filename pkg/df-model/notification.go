package model

type Notification struct {
	WorkflowName   string `json:"workflowName"`
	WorkflowExecID string `json:"workflowExecID"`
}
