package main

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

type Server struct {
	Mux *http.ServeMux
}

func InitHttpServer() *Server {
	s := &Server{}
	s.registerRoutes()
	return s
}

func (s *Server) registerRoutes() {
	s.Mux = http.NewServeMux()
	s.Mux.HandleFunc("/incoming", incomingHandler())
	s.Mux.HandleFunc("/status", HandleStatus())
}

func HandleStatus() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("OK"))
		if err != nil {
			log.WithError(err).Warn("Status handler response")
		}
	}
}

func (s *Server) Close() error {
	log.Infof("shutting down server")
	return nil
}
