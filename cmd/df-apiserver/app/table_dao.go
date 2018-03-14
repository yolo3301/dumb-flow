package app

import (
	"github.com/yolo3301/dumb-flow/pkg/df-model"
)

// TableDAO defines the interface for table storage.
type TableDAO interface {
	CreateOrUpdateWorkflowDef(workflowDef *model.WorkflowDef) error
	DeleteWorkflowDef(workflowName string) error
	GetWorkflowDef(workflowName string) (*model.WorkflowDef, error)
	CreateOrUpdateWorkflowExec(workflowName string) (string, error)
	GetWorkflowExec(workflowName, workflowExecID string) (*model.WorkflowExec, error)
	GetWorkflowExecs(workflowName string) ([]model.WorkflowExec, error)
	DeleteWorkflowExec(workflowName, workflowExecID string) error
	PauseWorkflowExec(workflowName, workflowExecID string) error
	AbortWorkflowExec(workflowName, workflowExecID string) error
	CompleteWorkflowExec(workflowName, workflowExecID string) error
	GetEvents(workflowName, workflowExecID string, allowedStates map[model.EventState]bool) ([]model.Event, error)
	CreateOrUpdateEvents(workflowName, workflowExecID string, events []model.Event) ([]model.Event, error)
	ResetEvents(workflowName, workflowExecID string) error

	// Test only, should remove later.
	SanityCheck() (string, error)
}
