package model

import (
	"fmt"
)

// Workflow defines the model of a workflow.
type WorkflowDef struct {
	Name    string            `json:"name,omitempty"`
	Configs map[string]string `json:"configs,omitempty"` // custom workflow configs
}

type WorkItemDef struct {
	Name         string            `json:"name,omitempty"`
	WorkflowName string            `json:"workflowName,omitempty"`
	Configs      map[string]string `json:"configs,omitempty"`
}

func (def *WorkflowDef) Key() string {
	return def.Name
}

func (def *WorkItemDef) Key() string {
	return fmt.Sprintf("%v-%v", def.WorkflowName, def.Name)
}
