package app

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"time"

	"github.com/coreos/etcd/clientv3"
	model "github.com/yolo3301/dumb-flow/pkg/df-model"
)

type DefaultTableDAO struct {
	etcdEndpoint string
	ctx          context.Context
}

func NewDefaultTableDAO() DefaultTableDAO {
	ep := os.Getenv("ETCD_ENDPOINT")
	if ep == "" {
		panic("Can't find ECTD_ENDPOINT in env vars")
	}

	return DefaultTableDAO{etcdEndpoint: ep, ctx: context.TODO()}
}

func (dao DefaultTableDAO) newEtcdClient() (*clientv3.Client, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{dao.etcdEndpoint},
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

func (dao DefaultTableDAO) SanityCheck() (string, error) {
	etcd, err := dao.newEtcdClient()
	if err != nil {
		return "", err
	}
	defer etcd.Close()

	var buffer bytes.Buffer
	buffer.WriteString("put foo1 = bar1\n")
	_, err = etcd.Put(dao.ctx, "foo1", "bar1")
	if err != nil {
		buffer.WriteString(err.Error() + "\n")
		return "", err
	}
	buffer.WriteString("put foo2 = bar2\n")
	_, err = etcd.Put(dao.ctx, "foo2", "bar2")
	if err != nil {
		buffer.WriteString(err.Error() + "\n")
		return "", err
	}

	buffer.WriteString("get foo1\n")
	res, err := etcd.Get(dao.ctx, "foo1")
	if err != nil {
		buffer.WriteString(err.Error() + "\n")
		return "", err
	}
	buffer.WriteString("got foo1: " + string(res.Kvs[0].Value) + "\n")

	buffer.WriteString("get foo prefix\n")
	res, err = etcd.Get(dao.ctx, "foo", clientv3.WithPrefix())
	if err != nil {
		buffer.WriteString(err.Error() + "\n")
		return "", err
	}
	buffer.WriteString("got foo prefix: " + string(res.Kvs[0].Value) + " " + string(res.Kvs[1].Value) + "\n")

	return buffer.String(), nil
}
