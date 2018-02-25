package main

import (
	"github.com/yolo3301/dumb-flow/cmd/df-apiserver/app"
)

func main() {
	eventController := app.NewInMemoryEventController()
	workflowController := app.NewInMemoryWorkflowController()
	server := app.NewDumbflowServer(eventController, workflowController)
	server.Run()
}
