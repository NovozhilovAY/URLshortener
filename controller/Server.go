package controller

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Server struct {
	router   *mux.Router
	handlers *Handlers
}

func NewServer(handlers *Handlers) *Server {
	return &Server{mux.NewRouter(), handlers}
}

func (s *Server) Run() {
	s.router.HandleFunc("/", s.handlers.AddUrl).Methods("POST")
	s.router.HandleFunc("/{code}", s.handlers.GetUrl).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", s.router))
}
