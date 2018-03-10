package app

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/coreos/etcd/clientv3"
	"github.com/satori/go.uuid"
	model "github.com/yolo3301/dumb-flow/pkg/df-model"
)

type DefaultTableDAO struct {
	ctx    context.Context
	client *clientv3.Client
}

func NewDefaultTableDAO() (*DefaultTableDAO, error) {
	ep := os.Getenv("ETCD_ENDPOINT")
	if ep == "" {
		return nil, fmt.Errorf("Can't find ECTD_ENDPOINT in env vars")
	}

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{ep},
		DialTimeout: 10 * time.Second,
	})
	if err != nil {
		return nil, err
	}

	return &DefaultTableDAO{ctx: context.TODO(), client: cli}, nil
}

func (dao DefaultTableDAO) Close() {
	dao.client.Close()
}

func (dao DefaultTableDAO) CreateOrUpdateWorkflowDef(workflowDef *model.WorkflowDef) error {
	content, err := json.Marshal(workflowDef)
	if err != nil {
		return err
	}

	res, err := dao.client.Put(dao.ctx, workflowDef.Key(), string(content))
	if err != nil {
		return err
	}

	log.Print(res.Header)
	return nil
}

func (dao DefaultTableDAO) DeleteWorkflowDef(workflowName string) error {
	res, err := dao.client.Delete(dao.ctx, workflowName)
	if err != nil {
		return err
	}

	log.Print(res.Header)
	return nil
}

// Caller can give optional exec ID otherwise will be auto generated.
func (dao DefaultTableDAO) CreateWorkflowExec(workflowName string) (string, error) {
	id := uuid.NewV4()
	exec := &model.WorkflowExec{
		ID:           id.String(),
		WorkflowName: workflowName,
		State:        model.WorkflowPending,
	}

	content, err := json.Marshal(exec)
	if err != nil {
		return "", err
	}

	res, err := dao.client.Put(dao.ctx, exec.Key(), string(content))
	if err != nil {
		return "", err
	}

	log.Print(res.Header)
	return id.String(), nil
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

func (dao DefaultTableDAO) SanityCheck() (string, error) {
	var buffer bytes.Buffer
	buffer.WriteString("put foo1 = bar1\n")
	_, err := dao.client.Put(dao.ctx, "foo1", "bar1")
	if err != nil {
		buffer.WriteString(err.Error() + "\n")
		return "", err
	}
	buffer.WriteString("put foo2 = bar2\n")
	_, err = dao.client.Put(dao.ctx, "foo2", "bar2")
	if err != nil {
		buffer.WriteString(err.Error() + "\n")
		return "", err
	}

	buffer.WriteString("get foo1\n")
	res, err := dao.client.Get(dao.ctx, "foo1")
	if err != nil {
		buffer.WriteString(err.Error() + "\n")
		return "", err
	}
	buffer.WriteString("got foo1: " + string(res.Kvs[0].Value) + "\n")

	buffer.WriteString("get foo prefix\n")
	res, err = dao.client.Get(dao.ctx, "foo", clientv3.WithPrefix())
	if err != nil {
		buffer.WriteString(err.Error() + "\n")
		return "", err
	}
	buffer.WriteString("got foo prefix: " + string(res.Kvs[0].Value) + " " + string(res.Kvs[1].Value) + "\n")

	return buffer.String(), nil
}
