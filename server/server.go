package server

import (
	//"github.com/gobwas/ws"
	"fmt"

	//"log"
	"net/http"
)

type Server struct {
	manager *Manager
	mux     *http.ServeMux
	address string
	port    string
}

func New(address string, port string) *Server {
	return &Server{
		NewManager(),
		http.NewServeMux(),
		address,
		port,
	}
}

func (s *Server) SetupRoutes() {
	//static serving
	s.mux.Handle("/", http.FileServer(http.Dir("F:\\vscProjects\\FrontendProjects\\carcassonne-client")))

	//websocket handler
	s.mux.HandleFunc("/ws", s.manager.ServeWS)

}

func (s *Server) StartServing() error {

	return http.ListenAndServe(fmt.Sprintf("%s:%s", s.address, s.port), s.mux)
}
