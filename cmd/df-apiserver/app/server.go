package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// DumbflowServer is the API server.
type DumbflowServer struct {
	eventController    *EventController
	workflowController *WorkflowController
}

// NewDumbflowServer creates a API server.
func NewDumbflowServer(eController EventController, wfController WorkflowController) *DumbflowServer {
	return &DumbflowServer{
		eventController:    &eController,
		workflowController: &wfController}
}

// Run starts the server.
func (s *DumbflowServer) Run() {
	router := mux.NewRouter()
	router.HandleFunc("/sanitycheck", s.HandleSanityTest).Methods("GET")

	// TODO port move the env
	log.Fatal(http.ListenAndServe(":13301", router))
}

// HandleSanityTest handles sanity check requests.
func (s *DumbflowServer) HandleSanityTest(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Success!\n"))
}
