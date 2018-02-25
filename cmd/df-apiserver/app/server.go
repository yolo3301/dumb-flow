package app

import (
	"log"
	"net/http"
)

// DumbflowServer is the API server.
type DumbflowServer struct {
	eventController    *EventController
	workflowController *WorkflowController
}

func NewDumbflowServer(eController *EventController, wfController *WorkflowController) *DumbflowServer {
	return &DumbflowServer{
		eventController:    eController,
		workflowController: wfController}
}

func (s *DumbflowServer) HandleSanityTest(w http.ResponseWriter, r *http.Request) {
	log.Print("Sanity checked")
}

func Run() {
	// mapping http req to different handler
	log.Print("Not Implement")
}
