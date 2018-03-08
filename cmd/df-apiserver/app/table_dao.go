package app

import (
	"github.com/yolo3301/dumb-flow/pkg/df-model"
)

// TableDAO defines the interface for table storage.
type TableDAO interface {
	CreateOrUpdateWorkflowDef(workflowDef *model.WorkflowDef) error
	DeleteWorkflowDef(workflowName string) error
	// Caller can give optional exec ID otherwise will be auto generated.
	CreateWorkflowExec(workflowName string, workflowExecID string) (string, error)
	PauseWorkflowExec(workflowName string, workflowExecID string) error
	AbortWorkflowExec(workflowName string, workflowExecID string) error
	CompleteWorkflow(workflowName string, workflowExecID string) error
	GetEvents(workflowName string, workflowExecID string, allowedStates []model.EventState) ([]model.Event, error)
	// Returns events successfully created/updated.
	CreateOrUpdateEvents(workflowName string, workflowExecID string, events []model.Event) ([]model.Event, error)

	// Test only, should remove later.
	SanityCheck() (string, error)
}
