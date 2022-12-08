package server

import (
	"io"
	"net/http"
)

type Server struct {
	server *http.Server
}

func NewServer(address string, tasks chan Task, sols chan Sol) (Server, error) {
	mux := http.NewServeMux()
	for s, f := range handleGen {
		mux.HandleFunc(s, f(tasks, sols))
	}
	return Server{
		server: &http.Server{
			Addr:    address,
			Handler: mux,
		},
	}, nil
}

func (m Server) Run(end chan error) {
	go func(end chan error) {
		err := m.server.ListenAndServe()
		end <- err
	}(end)
}

var handleGen = map[string]func(tasks chan Task, sols chan Sol) func(w http.ResponseWriter, r *http.Request){
	"/hello": func(tasks chan Task, sols chan Sol) func(w http.ResponseWriter, r *http.Request) {
		return func(w http.ResponseWriter, r *http.Request) {
			io.WriteString(w, "Hello, world!\n")
		}
	},
}
