package model

type HandleEventsRet struct {
	WorkflowExecID  string
	HandledEvents   []Event
	CreatedEvents   []Event
	UnhandledEvents []Event
}
