package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

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

	wfDefSubrouter.HandleFunc("/", s.HandleDeleteWorkflowDef).Methods("PUT")
	wfDefSubrouter.HandleFunc("/", s.HandleGetWorkflowDef).Methods("GET")
	wfDefSubrouter.HandleFunc("/", s.HandleDeleteWorkflowDef).Methods("DELETE")

	wfExecSubrouter.HandleFunc("/", s.HandleCreateOrUpdateWorkflowExec).Methods("PUT")
	wfExecSubrouter.HandleFunc("/{workflowExecId}", s.HandleGetWorkflowExec).Methods("GET")
	wfExecSubrouter.HandleFunc("/", s.HandleGetWorkflowExecs).Methods("GET")
	wfExecSubrouter.HandleFunc("/{workflowExecId}", s.HandleDeleteWorkflowExec).Methods("DELETE")
	wfExecSubrouter.HandleFunc("/{workflowExecId}", s.HandleWorkflowExecAction).Methods("POST")

	eventsSubrouter.HandleFunc("/", s.HandleCreateOrUpdateEvents).Methods("PUT")
	eventsSubrouter.HandleFunc("/", s.HandleGetEvents).Methods("GET")
	eventsSubrouter.HandleFunc("/", s.HandleEventsAction).Methods("POST")

	log.Fatal(http.ListenAndServe(addr, router))
}

func (s *DumbflowServer) HandleCreateOrUpdateWorkflowDef(w http.ResponseWriter, r *http.Request) {

}

func (s *DumbflowServer) HandleGetWorkflowDef(w http.ResponseWriter, r *http.Request) {

}

func (s *DumbflowServer) HandleDeleteWorkflowDef(w http.ResponseWriter, r *http.Request) {

}

func (s *DumbflowServer) HandleCreateOrUpdateWorkflowExec(w http.ResponseWriter, r *http.Request) {

}

func (s *DumbflowServer) HandleGetWorkflowExec(w http.ResponseWriter, r *http.Request) {

}

func (s *DumbflowServer) HandleDeleteWorkflowExec(w http.ResponseWriter, r *http.Request) {

}

func (s *DumbflowServer) HandleGetWorkflowExecs(w http.ResponseWriter, r *http.Request) {

}

func (s *DumbflowServer) HandleWorkflowExecAction(w http.ResponseWriter, r *http.Request) {

}

func (s *DumbflowServer) HandleCreateOrUpdateEvents(w http.ResponseWriter, r *http.Request) {

}

func (s *DumbflowServer) HandleGetEvents(w http.ResponseWriter, r *http.Request) {

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
