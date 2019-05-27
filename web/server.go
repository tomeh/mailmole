package web

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Mailmole Http server.
type Server struct {
	*http.Server
}

// Create a new instance of the Mailmole Http server.
func NewServer(host string, port int) *Server {
	addr := fmt.Sprintf("%s:%d", host, port)

	return &Server{
		&http.Server{
			Addr:         addr,
			Handler:      getHandler(),
			WriteTimeout: getWriteTimeout(),
			ReadTimeout:  getReadTimeout(),
		},
	}
}

func (s *Server) GetBaseUrl() string {
	return s.Addr
}

func (s *Server) Start() {
	log.Fatal(s.ListenAndServe())
}

func getHandler() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/", appHandler)
	r.PathPrefix("/static/").Handler(staticHandler())

	return r
}

func getWriteTimeout() time.Duration {
	return 15 * time.Second
}

func getReadTimeout() time.Duration {
	return 15 * time.Second
}
