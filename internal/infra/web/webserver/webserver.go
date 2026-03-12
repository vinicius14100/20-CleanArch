package webserver

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type HandlerEntry struct {
	Path    string
	Method  string
	Handler http.HandlerFunc
}

type WebServer struct {
	Router         chi.Router
	Handlers       map[string]http.HandlerFunc
	MethodHandlers []HandlerEntry
	WebServerPort  string
}

func NewWebServer(serverPort string) *WebServer {
	return &WebServer{
		Router:        chi.NewRouter(),
		Handlers:      make(map[string]http.HandlerFunc),
		WebServerPort: serverPort,
	}
}

func (s *WebServer) AddHandler(path string, handler http.HandlerFunc) {
	s.Handlers[path] = handler
}

func (s *WebServer) AddMethodHandler(method, path string, handler http.HandlerFunc) {
	s.MethodHandlers = append(s.MethodHandlers, HandlerEntry{Path: path, Method: method, Handler: handler})
}

// loop through the handlers and add them to the router
// register middeleware logger
// start the server
func (s *WebServer) Start() {
	s.Router.Use(middleware.Logger)
	for path, handler := range s.Handlers {
		s.Router.Handle(path, handler)
	}
	for _, entry := range s.MethodHandlers {
		switch entry.Method {
		case "GET":
			s.Router.Get(entry.Path, entry.Handler)
		case "POST":
			s.Router.Post(entry.Path, entry.Handler)
		}
	}
	http.ListenAndServe(s.WebServerPort, s.Router)
}
