package app

import (
	"github.com/yolo3301/dumb-flow/pkg/df-model"
)

// TableDAO defines the interface for table storage.
type TableDAO interface {
	CreateOrUpdateWorkflowDef(workflowDef *model.WorkflowDef) error
	DeleteWorkflowDef(workflowName string) error
	GetWorkflowDef(workflowName string) (*model.WorkflowDef, error)
	GetWorkflowDefs() ([]model.WorkflowDef, error)
	CreateWorkflowExec(workflowName string) (string, error)
	GetWorkflowExec(workflowName, workflowExecID string) (*model.WorkflowExec, error)
	GetWorkflowExecs(workflowName string) ([]model.WorkflowExec, error)
	DeleteWorkflowExec(workflowName, workflowExecID string) error
	DeleteWorkflowExecCascade(workflowName, workflowExecID string) error
	UpdateWorkflowExecState(workflowName, workflowExecID string, state model.WorkflowState) error
	CreateWorkItemExec(workflowName, workflowExecID, workItemName string) (string, error)
	GetWorkItemExec(workflowName, workflowExecID, workItemName, workItemExecID string) (*model.WorkItemExec, error)
	GetWorkItemExecs(workflowName, workflowExecID, workItemName string) ([]model.WorkItemExec, error)
	DeleteWorkItemExec(workflowName, workflowExecID, workItemName, workItemExecID string) error
	UpdateWorkItemExecState(workflowName, workflowExecID, workItemName, workItemExecID, state string) error
	GetEvents(workflowName, workflowExecID string, allowedStates map[model.EventState]bool) ([]model.Event, error)
	CreateOrUpdateEvents(events []model.Event) ([]model.Event, error)
	DeleteEvents(workflowName, workflowExecID string) error
	ResetEvents(workflowName, workflowExecID string) error

	// Test only, should remove later.
	SanityCheck() (string, error)
}
