package server

import (
	"io"
	"net/http"
)

type masterServer struct {
	server *http.Server
}

func NewMasterServer(address string, tasks chan Task, sols chan Sol) (masterServer, error) {
	mux := http.NewServeMux()
	for s, f := range masterHandleGen {
		mux.HandleFunc(s, f(tasks, sols))
	}
	return masterServer{
		server: &http.Server{
			Addr:    address,
			Handler: mux,
		},
	}, nil
}

func (m *masterServer) Run(end chan error) {
	go func(end chan error) {
		err := m.server.ListenAndServe()
		end <- err
	}(end)
}

var masterHandleGen = map[string]func(tasks chan Task, sols chan Sol) func(w http.ResponseWriter, r *http.Request){
	"/hello": func(tasks chan Task, sols chan Sol) func(w http.ResponseWriter, r *http.Request) {
		return func(w http.ResponseWriter, r *http.Request) {
			io.WriteString(w, "Hello, world!\n")
		}
	},
}
