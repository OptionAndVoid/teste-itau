package server

import (
	"fmt"
	"log/slog"
	"net/http"
)

type Server struct {
	Mux  *http.ServeMux
	host string
	port uint32
}

func NewServer() *Server {
	slog.Info("Server setup deu bom")
	return &Server{
		Mux:  nil,
		host: "localhost",
		port: 8080,
	}
}

func (s *Server) WithHost(host string) *Server {
	s.host = host
	return s
}

func (s *Server) WithPort(port uint32) *Server {
	s.port = port
	return s
}

func (s *Server) Run() error {
	listenAddr := fmt.Sprintf("%s:%d", s.host, s.port)
	slog.Info("starting server", "addr", listenAddr)
	return http.ListenAndServe(listenAddr, s.Mux)
}

func (s *Server) WithMux(mux *http.ServeMux) *Server {
	s.Mux = mux
	return s
}
