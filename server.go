package hiddentunes

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/acme/autocert"
)

type Server struct {
	HTTPServer *http.Server
	Manager    *autocert.Manager
}

func NewServer(port string, handler http.Handler, manager *autocert.Manager) *Server {
	return &Server{
		HTTPServer: &http.Server{
			Addr:           ":" + port,
			Handler:        handler,
			MaxHeaderBytes: 1 << 20,
			ReadTimeout:    time.Second * 10,
			WriteTimeout:   time.Second * 10,
			TLSConfig:      manager.TLSConfig(),
		},
		Manager: manager,
	}
}

func (s *Server) Run() error {
	go http.ListenAndServe(":80", s.Manager.HTTPHandler(nil))
	logrus.Printf("Server started on port %s", s.HTTPServer.Addr)
	return s.HTTPServer.ListenAndServeTLS("", "")
}
