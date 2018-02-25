package app

import (
	"github.com/yolo3301/dumb-flow/pkg/df-model"
)

// InMemoryWorkflowController in memory implementation of workflow controller.
type InMemoryWorkflowController struct {
	workflows map[string]model.Workflow
}

// NewInMemoryWorkflowController new
func NewInMemoryWorkflowController() InMemoryWorkflowController {
	return InMemoryWorkflowController{workflows: make(map[string]model.Workflow)}
}

func (c InMemoryWorkflowController) SaveWorkflow(wf model.Workflow) error {
	c.workflows[wf.Name] = wf
	return nil
}

func (c InMemoryWorkflowController) GetWorkflow(name string) (model.Workflow, error) {
	return c.workflows[name], nil
}
