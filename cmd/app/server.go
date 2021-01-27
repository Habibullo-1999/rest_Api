package app

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Habibullo-1999/rest_Api/pkg/note"
	"github.com/gorilla/mux"

)

const (
	GET = "GET"
	POST = "POST"
	PUT = "PUT"
	DELETE = "DELETE"
	UPDATE ="UPDATE"
)

type Server struct {
	mux     *mux.Router
	noteSvc *note.Service
}

func NewServer(mux *mux.Router, noteSvc *note.Service) *Server {
	return &Server{mux: mux, noteSvc: noteSvc}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

func (s *Server) Init() {
	s.mux.HandleFunc("/notes", s.handleGetNotes).Methods(GET)
}

func (s *Server) handleGetNotes(w http.ResponseWriter, r *http.Request) {
	items, err := s.noteSvc.GetNote(r.Context())
	if err != nil {
		log.Print(err)
		return
	}
	resJson(w, items)

}

func resJson(w http.ResponseWriter, iData interface{}) {

	data, err := json.Marshal(iData)

	if err != nil {
		errWriter(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(data)

	if err != nil {

		log.Print(err)
	}
}

// function for writing an error in responseWriter
func errWriter(w http.ResponseWriter, httpSts int, err error) {
	log.Print(err)
	http.Error(w, http.StatusText(httpSts), httpSts)
}
