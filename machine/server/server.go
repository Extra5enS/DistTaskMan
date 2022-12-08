package server

import (
	"io"
	"net/http"
)

type Server struct {
	server *http.Server
}

func NewServer(address string, tasks chan Task) (Server, error) {
	mux := http.NewServeMux()
	for s, f := range handleGen {
		mux.HandleFunc(s, f(tasks))
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

var handleGen = map[string]func(tasks chan Task) func(w http.ResponseWriter, r *http.Request){
	"/hello": func(tasks chan Task) func(w http.ResponseWriter, r *http.Request) {
		return func(w http.ResponseWriter, r *http.Request) {
			io.WriteString(w, "Hello, world!\n")
		}
	},
	"/master_req": func(tasks chan Task) func(w http.ResponseWriter, r *http.Request) {
		return func(w http.ResponseWriter, r *http.Request) {
			if !r.URL.Query().Has("address") {
				io.WriteString(w, "error: no address")
				return
			}
			address := r.URL.Query().Get("address")
			if !r.URL.Query().Has("num") {
				io.WriteString(w, "error: no num")
				return
			}
			num := r.URL.Query().Get("num")
			sols := make(chan Sol)
			task := Task{
				Req:      "master_req",
				Args:     []string{address, num},
				Solution: sols,
			}
			tasks <- task

			sol := <-sols
			io.WriteString(w, string(sol))
		}
	},
}
