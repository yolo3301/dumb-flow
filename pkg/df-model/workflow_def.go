package model

import (
	"fmt"
)

const WorkflowDefPrefix = "wfdef"
const WorkItemDefPrefix = "widef"

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
	return fmt.Sprintf("%v-%v", WorkflowDefPrefix, def.Name)
}

func (def *WorkItemDef) Key() string {
	return fmt.Sprintf("%v-%v-%v", WorkItemDefPrefix, def.WorkflowName, def.Name)
}
