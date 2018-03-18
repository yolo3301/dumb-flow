package dfclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	uuid "github.com/satori/go.uuid"
	"github.com/yolo3301/dumb-flow/pkg/df-model"
)

type DumbflowClient struct {
	Endpoint string
}

func NewDumbflowClient() (*DumbflowClient, error) {
	ep := os.Getenv("DF_ENDPOINT")
	if ep == "" {
		return nil, fmt.Errorf("Can't find DF_ENDPOINT in env vars")
	}

	client := &DumbflowClient{
		Endpoint: "http://" + ep,
	}

	return client, nil
}

func (client *DumbflowClient) CreateWorkflowDef(workflowName string, configs map[string]string) error {
	workflowDef := &model.WorkflowDef{
		Name:    workflowName,
		Configs: configs,
	}

	body, err := json.Marshal(workflowDef)
	if err != nil {
		return err
	}

	requestURL := fmt.Sprintf("%v/workflowDef/%v", client.Endpoint, workflowName)
	req, err := http.NewRequest(http.MethodPut, requestURL, bytes.NewReader(body))
	if err != nil {
		return err
	}

	httpClient := &http.Client{}
	_, err = httpClient.Do(req)
	if err != nil {
		return err
	}

	return nil
}

func (client *DumbflowClient) GetWorkflowDef(workflowName string) (*model.WorkflowDef, error) {
	requestURL := fmt.Sprintf("%v/workflowDef/%v", client.Endpoint, workflowName)
	bodyBytes, err := client.sendGetRequest(requestURL)
	if err != nil {
		return nil, err
	}

	var workflowDef model.WorkflowDef
	err = json.Unmarshal(bodyBytes, &workflowDef)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	return &workflowDef, nil
}

func (client *DumbflowClient) DeleteWorkflowDef(workflowName string) error {
	requestURL := fmt.Sprintf("%v/workflowDef/%v", client.Endpoint, workflowName)
	return client.sendDeleteRequest(requestURL)
}

func (client *DumbflowClient) CreateWorkflowExec(workflowName string) (string, error) {
	requestURL := fmt.Sprintf("%v/workflowDef/%v/workflowExec", client.Endpoint, workflowName)
	req, err := http.NewRequest(http.MethodPut, requestURL, nil)
	if err != nil {
		return "", err
	}

	httpClient := &http.Client{}
	res, err := httpClient.Do(req)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()
	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", nil
	}

	return string(bodyBytes), nil
}

func (client *DumbflowClient) GetWorkflowExec(workflowName, workflowExecID string) (*model.WorkflowExec, error) {
	requestURL := fmt.Sprintf("%v/workflowDef/%v/workflowExec/%v", client.Endpoint, workflowName, workflowExecID)
	bodyBytes, err := client.sendGetRequest(requestURL)
	if err != nil {
		return nil, err
	}

	var workflowExec model.WorkflowExec
	err = json.Unmarshal(bodyBytes, &workflowExec)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	return &workflowExec, nil
}

func (client *DumbflowClient) DeleteWorkflowExec(workflowName, workflowExecID string) error {
	requestURL := fmt.Sprintf("%v/workflowDef/%v/workflowExec/%v", client.Endpoint, workflowName, workflowExecID)
	return client.sendDeleteRequest(requestURL)
}

func (client *DumbflowClient) CreateEvent(workflowName, workflowExecID, workItemName, workItemExecID, payload string, eventType model.EventType) (string, error) {
	id := uuid.NewV4()
	event := &model.Event{
		ID:             id.String(),
		WorkflowName:   workflowName,
		WorkflowExecID: workflowExecID,
		WorkItemName:   workItemName,
		WorkItemExecID: workItemExecID,
		Payload:        payload,
		EventType:      eventType,
	}
	events := []model.Event{*event}

	body, err := json.Marshal(events)
	if err != nil {
		return "", err
	}

	requestURL := fmt.Sprintf("%v/workflowDef/%v/workflowExec/%v/events", client.Endpoint, workflowName, workflowExecID)
	req, err := http.NewRequest(http.MethodPut, requestURL, bytes.NewReader(body))
	if err != nil {
		return "", err
	}

	httpClient := &http.Client{}
	_, err = httpClient.Do(req)
	if err != nil {
		return "", err
	}

	return id.String(), nil
}

func (client *DumbflowClient) GetEvents(workflowName, workflowExecID string) ([]model.Event, error) {
	requestURL := fmt.Sprintf("%v/workflowDef/%v/workflowExec/%v/events", client.Endpoint, workflowName, workflowExecID)
	bodyBytes, err := client.sendGetRequest(requestURL)
	if err != nil {
		return nil, err
	}

	var events []model.Event
	err = json.Unmarshal(bodyBytes, &events)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	return events, nil
}

func (client *DumbflowClient) sendGetRequest(requestURL string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		return nil, err
	}

	httpClient := &http.Client{}
	res, err := httpClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}

func (client *DumbflowClient) sendDeleteRequest(requestURL string) error {
	req, err := http.NewRequest(http.MethodDelete, requestURL, nil)
	if err != nil {
		return err
	}

	httpClient := &http.Client{}
	_, err = httpClient.Do(req)

	if err != nil {
		return err
	}

	return nil
}
