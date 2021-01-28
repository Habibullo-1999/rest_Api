package app

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Habibullo-1999/rest_Api/pkg/note"
	"github.com/Habibullo-1999/rest_Api/pkg/types"
	"github.com/gorilla/mux"

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
	s.mux.HandleFunc("/notes", s.handleGetNotes).Methods(types.GET)
	s.mux.HandleFunc("/notes", s.handleSaveNote).Methods(types.POST)
	s.mux.HandleFunc("/notes", s.handleUpdateNote).Methods(types.PUT)
	s.mux.HandleFunc("/notes/getById/{id}", s.handleGetNotesById).Methods(types.GET)
}

//All notes
func (s *Server) handleGetNotes(w http.ResponseWriter, r *http.Request) {
	items, err := s.noteSvc.GetNotes(r.Context())
	if err != nil {
		log.Print(err)
		return
	}
	resJson(w, items)

}
//Save notes
func (s *Server) handleSaveNote(w http.ResponseWriter, r *http.Request) {
	// item := types.Note{}

	note := &types.Note{}

	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		errWriter(w, http.StatusInternalServerError, err)
		return
	}

	noteY, err := s.noteSvc.SaveNote(r.Context(), note)
	if err != nil {
		errWriter(w, http.StatusInternalServerError, err)
		return
	}

	resJson(w, noteY)
}

// Update notes by Id
func (s *Server) handleUpdateNote(w http.ResponseWriter, r *http.Request) {

	note := &types.Note{}

	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		errWriter(w, http.StatusInternalServerError, err)
		return
	}

	noteY, err := s.noteSvc.UpdateNote(r.Context(), note)
	if err != nil {
		errWriter(w, http.StatusInternalServerError, err)
		return
	}

	resJson(w, noteY)
}
// Get by id notes
func (s *Server) handleGetNotesById(w http.ResponseWriter, r *http.Request) {

	id, ok := mux.Vars(r)["id"]
	if !ok {
		log.Print("Status bad Request")
		return
	}
	noteId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		errWriter(w, http.StatusBadRequest, err)
		return
	}
	
	item, err := s.noteSvc.GetById(r.Context(), noteId)
	if err != nil {
		errWriter(w, http.StatusBadRequest, err)
		return
	}
	resJson(w, item)
}


