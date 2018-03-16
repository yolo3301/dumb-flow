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

func (dao DefaultTableDAO) GetWorkflowDef(workflowName string) (*model.WorkflowDef, error) {
	res, err := dao.client.Get(dao.ctx, fmt.Sprintf("%v-%v", model.WorkflowDefPrefix, workflowName))
	if err != nil {
		return nil, err
	}

	if len(res.Kvs) == 0 {
		return nil, fmt.Errorf("Not found")
	}

	var workflowDef model.WorkflowDef
	err = json.Unmarshal([]byte(res.Kvs[0].Value), &workflowDef)
	if err != nil {
		return nil, err
	}

	return &workflowDef, nil
}

func (dao DefaultTableDAO) DeleteWorkflowDef(workflowName string) error {
	res, err := dao.client.Delete(dao.ctx, fmt.Sprintf("%v-%v", model.WorkflowDefPrefix, workflowName))
	if err != nil {
		return err
	}

	log.Print(res.Header)
	return nil
}

func (dao DefaultTableDAO) CreateOrUpdateWorkflowExec(workflowName string) (string, error) {
	id := uuid.NewV4()
	exec := &model.WorkflowExec{
		ID:           id.String(),
		WorkflowName: workflowName,
		State:        model.WorkflowPending,
	}

	err := dao.SaveWorkflowExec(exec)
	if err != nil {
		return "", nil
	}

	return id.String(), nil
}

func (dao DefaultTableDAO) GetWorkflowExec(workflowName, workflowExecID string) (*model.WorkflowExec, error) {
	res, err := dao.client.Get(dao.ctx, fmt.Sprintf("%v-%v-%v", model.WorkflowExecPrefix, workflowName, workflowExecID))
	if err != nil {
		return nil, err
	}

	if res.Count == 0 {
		return nil, fmt.Errorf("Not found")
	}

	var exec model.WorkflowExec
	err = json.Unmarshal([]byte(res.Kvs[0].Value), &exec)
	if err != nil {
		return nil, err
	}

	return &exec, nil
}

func (dao DefaultTableDAO) GetWorkflowExecs(workflowName string) ([]model.WorkflowExec, error) {
	res, err := dao.client.Get(dao.ctx, fmt.Sprintf("%v-%v", model.WorkflowExecPrefix, workflowName), clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}

	var execs []model.WorkflowExec
	for _, v := range res.Kvs {
		var e model.WorkflowExec
		err = json.Unmarshal([]byte(v.Value), &e)
		if err != nil {
			log.Printf("Failed to deserialize '%v'", v.Value)
		}

		execs = append(execs, e)
	}

	return execs, nil
}

func (dao DefaultTableDAO) DeleteWorkflowExec(workflowName, workflowExecID string) error {
	res, err := dao.client.Delete(dao.ctx, fmt.Sprintf("%v-%v-%v", model.WorkflowExecPrefix, workflowName, workflowExecID))
	if err != nil {
		return err
	}

	log.Print(res.Header)
	return nil
}

func (dao DefaultTableDAO) PauseWorkflowExec(workflowName, workflowExecID string) error {
	return dao.UpdateWorkflowExecState(workflowName, workflowExecID, model.WorkflowPaused)
}

func (dao DefaultTableDAO) AbortWorkflowExec(workflowName, workflowExecID string) error {
	return dao.UpdateWorkflowExecState(workflowName, workflowExecID, model.WorkflowAborted)
}

func (dao DefaultTableDAO) CompleteWorkflowExec(workflowName, workflowExecID string) error {
	return dao.UpdateWorkflowExecState(workflowName, workflowExecID, model.WorkflowCompleted)
}

func (dao DefaultTableDAO) GetEvents(workflowName, workflowExecID string, allowedStates map[model.EventState]bool) ([]model.Event, error) {
	res, err := dao.client.Get(dao.ctx, fmt.Sprintf("%v-%v-%v", model.EventPrefix, workflowName, workflowExecID), clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}

	var events []model.Event
	for _, v := range res.Kvs {
		var e model.Event
		err = json.Unmarshal([]byte(v.Value), &e)
		if err != nil {
			log.Printf("Failed to deserialize '%v'", v.Value)
		}

		if allowedStates[e.State] {
			events = append(events, e)
		}
	}

	return events, nil
}

func (dao DefaultTableDAO) CreateOrUpdateEvents(workflowName, workflowExecID string, events []model.Event) ([]model.Event, error) {
	var updatedEvents []model.Event
	for _, e := range events {
		e.WorkflowName = workflowName
		e.WorkflowExecID = workflowExecID
		e.LastModifiedTime = time.Now()

		content, err := json.Marshal(e)
		if err != nil {
			log.Printf("Failed to serialize some event")
		}

		res, err := dao.client.Put(dao.ctx, e.Key(), string(content))
		if err != nil {
			log.Printf("Failed to put '%v'", e.Key())
		}

		log.Print(res.Header)

		updatedEvents = append(updatedEvents, e)
	}

	if len(updatedEvents) < len(events) {
		return updatedEvents, fmt.Errorf("Failed to put some events")
	}

	return updatedEvents, nil
}

func (dao DefaultTableDAO) ResetEvents(workflowName, workflowExecID string) error {
	res, err := dao.client.Get(dao.ctx, fmt.Sprintf("%v-%v-%v", model.EventPrefix, workflowName, workflowExecID), clientv3.WithPrefix())
	if err != nil {
		return err
	}

	for _, v := range res.Kvs {
		var e model.Event
		err = json.Unmarshal([]byte(v.Value), &e)
		if err != nil {
			return fmt.Errorf("Failed to deserialize '%v'", v.Value)
		}

		// TODO make timeout a workflow config
		if e.State == model.EventScheduling && time.Since(e.LastModifiedTime) > 5*time.Minute {
			e.State = model.EventCreated
			e.LastModifiedTime = time.Now()
			content, err := json.Marshal(e)
			if err != nil {
				return fmt.Errorf("Failed to serialize '%v'", e.Key())
			}

			_, err = dao.client.Put(dao.ctx, e.Key(), string(content))
			if err != nil {
				log.Printf("Failed to put '%v'", e.Key())
			}
		}
	}

	return nil
}

func (dao DefaultTableDAO) UpdateWorkflowExecState(workflowName, workflowExecID string, state model.WorkflowState) error {
	exec, err := dao.GetWorkflowExec(workflowName, workflowExecID)
	if err != nil {
		return err
	}

	exec.State = state
	err = dao.SaveWorkflowExec(exec)
	if err != nil {
		return err
	}

	return nil
}

func (dao DefaultTableDAO) SaveWorkflowExec(workflowExec *model.WorkflowExec) error {
	content, err := json.Marshal(workflowExec)
	if err != nil {
		return err
	}

	res, err := dao.client.Put(dao.ctx, workflowExec.Key(), string(content))
	if err != nil {
		return err
	}

	log.Print(res.Header)
	return nil
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
