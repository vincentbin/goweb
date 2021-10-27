package server

import (
	"html/template"
	"net/http"
)

type Handler func(*Context)

type Server struct {
	*Group
	router        *Router
	HTMLTemplates *template.Template
}

func InitServer() *Server {
	server := &Server{
		router: InitRouter(),
	}
	server.Group = &Group{server: server}
	return server
}

func (s *Server) Get(url string, handler Handler) {
	s.router.addRouter("GET", url, handler)
}

func (s *Server) Post(url string, handler Handler) {
	s.router.addRouter("POST", url, handler)
}

func (s *Server) LoadTemplate(pattern string) {
	s.HTMLTemplates = template.Must(template.New("").ParseGlob(pattern))
}

func (s *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	context := InitContext(w, req, s)
	s.router.handle(context)
}

func (s *Server) Run(address string) error {
	return http.ListenAndServe(address, s)
}
