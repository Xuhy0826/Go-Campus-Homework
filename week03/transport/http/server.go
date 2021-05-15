package http

import (
	"context"
	"errors"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type ServerOption func(*Server)

// Option模式进行配置：

// Network with server network.
func Network(network string) ServerOption {
	return func(s *Server) {
		s.network = network
	}
}

// Address with server address.
func Address(addr string) ServerOption {
	return func(s *Server) {
		s.address = addr
	}
}

// Timeout with server timeout.
func Timeout(timeout time.Duration) ServerOption {
	return func(s *Server) {
		s.timeout = timeout
	}
}


// *************************************************************

// Server HTTP server wrapper.
type Server struct {
	*http.Server

	lis     net.Listener
	network string
	address string
	timeout time.Duration
	router  *mux.Router
	log     *log.Logger
}

func NewServer(opts ...ServerOption) *Server{
	srv := &Server{
		network: "tcp",
		address: ":0",
		timeout: time.Second,
		log:     log.Default(),
	}
	for _, o := range opts {
		o(srv)
	}
	srv.router = mux.NewRouter()
	srv.Server = &http.Server{Handler: srv}
	return srv
}

// Server implement Handler interface
// ServeHTTP should write reply headers and data to the ResponseWriter and then return.
func (s *Server) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	ctx, cancel := context.WithTimeout(req.Context(), s.timeout)
	defer cancel()

	s.router.ServeHTTP(res, req.WithContext(ctx))
}

// Handle registers a new route with a matcher for the URL path.
func (s *Server) Handle(path string, h http.Handler) {
	s.router.Handle(path, h)
}

// HandlePrefix registers a new route with a matcher for the URL path prefix.
func (s *Server) HandlePrefix(prefix string, h http.Handler) {
	s.router.PathPrefix(prefix).Handler(h)
}

// HandleFunc registers a new route with a matcher for the URL path.
func (s *Server) HandleFunc(path string, h http.HandlerFunc) {
	s.router.HandleFunc(path, h)
}


// Start start the HTTP server.
func (s *Server) Start() error {
	lis, err := net.Listen(s.network, s.address)
	if err != nil {
		return err
	}
	s.lis = lis
	s.log.Printf("[HTTP] server listening on: %s\n", lis.Addr().String())
	if err := s.Serve(lis); !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

// Stop stop the HTTP server.
func (s *Server) Stop() error {
	s.log.Printf("[HTTP] server stopping\n")
	return s.Shutdown(context.Background())
}
