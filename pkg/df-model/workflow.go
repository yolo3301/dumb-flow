package model

// Workflow defines the model of a workflow.
type Workflow struct {
	Name    string
	Configs map[string]string // configuration data - option: store in etcd
}
