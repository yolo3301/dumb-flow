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
	GetWorkflowExec(workflowName string, workflowExecID string) (*model.WorkflowExec, error)
	GetWorkflowExecs(workflowName string) ([]model.WorkflowExec, error)
	DeleteWorkflowExec(workflowName string, workflowExecID string) error
	PauseWorkflowExec(workflowName string, workflowExecID string) error
	AbortWorkflowExec(workflowName string, workflowExecID string) error
	CompleteWorkflowExec(workflowName string, workflowExecID string) error
	GetEvents(workflowName string, workflowExecID string, allowedStates map[model.EventState]bool) ([]model.Event, error)
	CreateOrUpdateEvents(workflowName string, workflowExecID string, events []model.Event) ([]model.Event, error)
	ResetEvent(workflowName string, workflowExecID string, workflowItemName string, workflowItemExecID string, eventID string) error

	// Test only, should remove later.
	SanityCheck() (string, error)
}
