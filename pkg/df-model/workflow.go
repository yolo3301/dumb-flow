package model

// Workflow defines the model of a workflow.
type Workflow struct {
	name    string
	configs map[string]string
}
