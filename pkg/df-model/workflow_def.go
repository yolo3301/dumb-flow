package model

// Workflow defines the model of a workflow.
type WorkflowDef struct {
	Name      string
	WorkItems []WorkItemDef
	Configs   map[string]string // custom workflow configs
}

type WorkItemDef struct {
	Name    string
	Configs map[string]string
}
