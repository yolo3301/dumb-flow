package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/yolo3301/dumb-flow/pkg/df-model"

	"github.com/gorilla/mux"
)

// DumbflowServer is the API server.
type DumbflowServer struct {
	tableDAO TableDAO
	queueDAO QueueDAO
}

// NewDumbflowServer creates a API server.
func NewDumbflowServer(tableDAO TableDAO, queueDAO QueueDAO) *DumbflowServer {
	return &DumbflowServer{tableDAO: tableDAO, queueDAO: queueDAO}
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
	router.HandleFunc("/tablesanitycheck", s.HandleTableSanityCheck).Methods("GET")

	wfDefSubrouter := router.PathPrefix("/workflowDef/{workflowName}").Subrouter()
	wfExecSubrouter := wfDefSubrouter.PathPrefix("/workflowExec").Subrouter()
	eventsSubrouter := wfExecSubrouter.PathPrefix("/{workflowExecId}/events").Subrouter()

	wfDefSubrouter.HandleFunc("", s.HandleCreateOrUpdateWorkflowDef).Methods("PUT")
	wfDefSubrouter.HandleFunc("", s.HandleGetWorkflowDef).Methods("GET")
	wfDefSubrouter.HandleFunc("", s.HandleDeleteWorkflowDef).Methods("DELETE")

	wfExecSubrouter.HandleFunc("", s.HandleCreateWorkflowExec).Methods("PUT")
	wfExecSubrouter.HandleFunc("/{workflowExecId}", s.HandleGetWorkflowExec).Methods("GET")
	wfExecSubrouter.HandleFunc("", s.HandleGetWorkflowExecs).Methods("GET")
	wfExecSubrouter.HandleFunc("/{workflowExecId}", s.HandleDeleteWorkflowExec).Methods("DELETE")
	wfExecSubrouter.HandleFunc("/{workflowExecId}", s.HandleWorkflowExecAction).Methods("POST")

	eventsSubrouter.HandleFunc("", s.HandleCreateEvents).Methods("PUT")
	eventsSubrouter.HandleFunc("", s.HandleGetEvents).Methods("GET")
	eventsSubrouter.HandleFunc("", s.HandleEventsAction).Methods("POST")

	log.Fatal(http.ListenAndServe(addr, router))
}

func (s *DumbflowServer) HandleCreateOrUpdateWorkflowDef(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		http.Error(w, "Missing request body", 400)
		return
	}

	vars := mux.Vars(r)
	workflowName := vars["workflowName"]

	var workflowDef model.WorkflowDef
	err := json.NewDecoder(r.Body).Decode(&workflowDef)
	if err != nil {
		http.Error(w, err.Error(), 400)
	}

	if workflowDef.Name != workflowName {
		http.Error(w, "Workflow name mismatch between url and the body", 400)
	}

	err = s.tableDAO.CreateOrUpdateWorkflowDef(&workflowDef)
	if err != nil {
		http.Error(w, err.Error(), 400)
	}
}

func (s *DumbflowServer) HandleGetWorkflowDef(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	workflowName := vars["workflowName"]

	workflowDef, err := s.tableDAO.GetWorkflowDef(workflowName)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	if workflowDef == nil {
		http.Error(w, "Workflow def not found", 404)
	}

	json.NewEncoder(w).Encode(workflowDef)
}

func (s *DumbflowServer) HandleDeleteWorkflowDef(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	workflowName := vars["workflowName"]

	err := s.tableDAO.DeleteWorkflowDef(workflowName)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
}

func (s *DumbflowServer) HandleCreateWorkflowExec(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	workflowName := vars["workflowName"]

	execID, err := s.tableDAO.CreateWorkflowExec(workflowName)
	if err != nil {
		http.Error(w, err.Error(), 400)
	}

	fmt.Fprint(w, execID)
}

func (s *DumbflowServer) HandleGetWorkflowExec(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	workflowName := vars["workflowName"]
	execID := vars["workflowExecId"]

	workflowExec, err := s.tableDAO.GetWorkflowExec(workflowName, execID)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	if workflowExec == nil {
		http.Error(w, "Workflow exec not found", 404)
	}

	json.NewEncoder(w).Encode(workflowExec)
}

func (s *DumbflowServer) HandleDeleteWorkflowExec(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	workflowName := vars["workflowName"]
	execID := vars["workflowExecId"]

	err := s.tableDAO.DeleteWorkflowExec(workflowName, execID)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
}

func (s *DumbflowServer) HandleGetWorkflowExecs(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	workflowName := vars["workflowName"]

	workflowExecs, err := s.tableDAO.GetWorkflowExecs(workflowName)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	json.NewEncoder(w).Encode(workflowExecs)
}

func (s *DumbflowServer) HandleWorkflowExecAction(w http.ResponseWriter, r *http.Request) {

}

func (s *DumbflowServer) HandleCreateEvents(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		http.Error(w, "Missing request body", 400)
		return
	}

	// todo add check for workflow name and exec id

	var events []model.Event
	err := json.NewDecoder(r.Body).Decode(&events)
	if err != nil {
		http.Error(w, err.Error(), 400)
	}

	createdEvents, err := s.tableDAO.CreateOrUpdateEvents(events)
	if err != nil {
		http.Error(w, err.Error(), 400)
	}

	if len(createdEvents) < len(events) {
		http.Error(w, "Some events failed to create", 500)
	}
}

func (s *DumbflowServer) HandleGetEvents(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	workflowName := vars["workflowName"]
	execID := vars["workflowExecId"]

	// todo add state filter

	events, err := s.tableDAO.GetEvents(workflowName, execID, nil)
	if err != nil {
		http.Error(w, err.Error(), 400)
	}

	json.NewEncoder(w).Encode(events)
}

func (s *DumbflowServer) HandleEventsAction(w http.ResponseWriter, r *http.Request) {

}

// HandleSanityTest handles sanity check requests.
func (s *DumbflowServer) HandleSanityTest(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Success!\n"))
}

func (s *DumbflowServer) HandleTableSanityCheck(w http.ResponseWriter, r *http.Request) {
	res, err := s.tableDAO.SanityCheck()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Write([]byte(res))
}
