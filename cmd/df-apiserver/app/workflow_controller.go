package app

import (
	"github.com/yolo3301/dumb-flow/pkg/df-model"
)

// WorkflowController is the interface for workflow controller.
type WorkflowController interface {
	SaveWorkflow(wf model.Workflow) error
	GetWorkflow(name string) (model.Workflow, error)
}
