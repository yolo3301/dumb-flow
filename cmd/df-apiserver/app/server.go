package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	model "github.com/yolo3301/dumb-flow/pkg/df-model"
)

// DumbflowServer is the API server.
type DumbflowServer struct {
	eventController    EventController
	workflowController WorkflowController
}

// NewDumbflowServer creates a API server.
func NewDumbflowServer(eController EventController, wfController WorkflowController) *DumbflowServer {
	return &DumbflowServer{
		eventController:    eController,
		workflowController: wfController}
}

// Run starts the server.
func (s *DumbflowServer) Run() {
	addr := ":13301"
	port := os.Getenv("DF_PORT")
	if port != "" {
		addr = fmt.Sprintf(":%v", port)
	}

	router := mux.NewRouter()
	router.HandleFunc("/sanitycheck", s.HandleSanityTest).Methods("GET")
	router.HandleFunc("/event", s.HandleSaveEvent).Methods("PUT")
	router.HandleFunc("/events/{eventkey}", s.HandleGetEvents).Methods("GET")

	log.Fatal(http.ListenAndServe(addr, router))
}

// HandleSanityTest handles sanity check requests.
func (s *DumbflowServer) HandleSanityTest(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Success!\n"))
}

func (s *DumbflowServer) HandleSaveEvent(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}

	var e model.Event
	err := json.NewDecoder(r.Body).Decode(&e)
	if err != nil {
		http.Error(w, err.Error(), 400)
	}

	s.eventController.SaveEvent(e)
}

func (s *DumbflowServer) HandleGetEvents(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	eventKey := vars["eventkey"]
	events, err := s.eventController.GetEvents(eventKey)
	if err != nil {
		http.Error(w, err.Error(), 400)
	}
	json.NewEncoder(w).Encode(events)
}
