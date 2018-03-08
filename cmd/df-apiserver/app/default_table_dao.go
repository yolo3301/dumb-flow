package app

import (
	"fmt"
	"os"
	"time"

	"github.com/coreos/etcd/clientv3"
	model "github.com/yolo3301/dumb-flow/pkg/df-model"
)

type DefaultTableDAO struct {
	ectdEndpoint string
}

func NewDefaultTableDAO() DefaultTableDAO {
	ep := os.Getenv("ETCD_ENDPOINT")
	// temp disable for test
	// if ep == "" {
	// 	panic("Can't find ECTD_ENDPOINT in env vars")
	// }

	return DefaultTableDAO{ectdEndpoint: ep}
}

func (dao DefaultTableDAO) newEtcdClient() (*clientv3.Client, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{dao.ectdEndpoint},
		DialTimeout: 10 * time.Second,
	})

	if err != nil {
		return nil, err
	}

	return cli, nil
}

func (dao DefaultTableDAO) CreateOrUpdateWorkflowDef(workflowDef *model.WorkflowDef) error {
	return fmt.Errorf("Not implemented")
}

func (dao DefaultTableDAO) DeleteWorkflowDef(workflowName string) error {
	return fmt.Errorf("Not implemented")
}

// Caller can give optional exec ID otherwise will be auto generated.
func (dao DefaultTableDAO) CreateWorkflowExec(workflowName string, workflowExecID string) (string, error) {
	return "", fmt.Errorf("Not implemented")
}

func (dao DefaultTableDAO) PauseWorkflowExec(workflowName string, workflowExecID string) error {
	return fmt.Errorf("Not implemented")
}

func (dao DefaultTableDAO) AbortWorkflowExec(workflowName string, workflowExecID string) error {
	return fmt.Errorf("Not implemented")
}

func (dao DefaultTableDAO) CompleteWorkflow(workflowName string, workflowExecID string) error {
	return fmt.Errorf("Not implemented")
}

func (dao DefaultTableDAO) GetEvents(workflowName string, workflowExecID string, allowedStates []model.EventState) ([]model.Event, error) {
	return nil, fmt.Errorf("Not implemented")
}

// Returns events successfully created/updated.
func (dao DefaultTableDAO) CreateOrUpdateEvents(workflowName string, workflowExecID string, events []model.Event) ([]model.Event, error) {
	return nil, fmt.Errorf("Not implemented")
}
