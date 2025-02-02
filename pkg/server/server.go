package server

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
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

func (s *Server) WithMux(mux *http.ServeMux) *Server {
	s.Mux = mux
	return s
}

func (s *Server) RunWithTls(keyFilePath string, certFilePath string) error {
	key, err := os.Stat(keyFilePath)
	if err != nil {
		return fmt.Errorf("invalid ssl key file")
	}
	if key.IsDir() {
		return fmt.Errorf("ssl key file cannot be a dir")
	}

	cert, err := os.Stat(certFilePath)
	if err != nil {
		return fmt.Errorf("invalid ssl cert file")
	}
	if cert.IsDir() {
		return fmt.Errorf("ssl cert file cannot be a dir")
	}

	listenAddr := fmt.Sprintf("%s:%d", s.host, s.port)
	slog.Info("starting https server", "addr", listenAddr)
	return http.ListenAndServeTLS(listenAddr, certFilePath, keyFilePath, s.Mux)
}

func (s *Server) Run() error {
	listenAddr := fmt.Sprintf("%s:%d", s.host, s.port)
	slog.Info("starting http server", "addr", listenAddr)
	return http.ListenAndServe(listenAddr, s.Mux)
}
