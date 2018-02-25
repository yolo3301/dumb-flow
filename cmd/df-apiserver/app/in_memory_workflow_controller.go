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

// TestWorkflowController just dumb.
func (c InMemoryWorkflowController) TestWorkflowController() string {
	return "ok"
}
