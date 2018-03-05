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

	log.Fatal(http.ListenAndServe(addr, router))
}

// HandleSanityTest handles sanity check requests.
func (s *DumbflowServer) HandleSanityTest(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Success!\n"))
}
