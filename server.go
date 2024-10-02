package hiddentunes

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

type Server struct {
	HTTPServer *http.Server
}

func NewServer(port string, handler http.Handler) *Server {
	return &Server{
		HTTPServer: &http.Server{
			Addr:           ":" + port,
			Handler:        handler,
			MaxHeaderBytes: 1 << 20,
			ReadTimeout:    time.Second * 10,
			WriteTimeout:   time.Second * 10,
		},
	}
}

func (s *Server) Run() error {
	logrus.Printf("Server started on port %s", s.HTTPServer.Addr)
	return s.HTTPServer.ListenAndServe()
}
