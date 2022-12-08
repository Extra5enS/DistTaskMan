package server

import (
	"bufio"
	"encoding/json"
	"io"
	"net"
	"net/http"
)

type Server struct {
	conn  *net.Conn
	tasks chan Task
}

func NewServer(address string, tasks chan Task) (Server, error) {
	ln, _ := net.Listen("tcp", address)
	conn, _ := ln.Accept()

	mux := http.NewServeMux()
	for s, f := range handleGen {
		mux.HandleFunc(s, f(tasks))
	}
	return Server{
		conn:  &conn,
		tasks: tasks,
	}, nil
}

func (m Server) Run(end chan error) {
	go func(end chan error) {
		for {
			json_map := make(map[string]interface{})
			connReader := bufio.NewReader(*m.conn)
			line, isPrefix, err := connReader.ReadLine()
			err = json.Unmarshal(line, json_map)

			handles[args[0]](*m.conn, args[1:], m.tasks)
		}
	}(end)
}

var handles = map[string]func(conn net.Conn, args []string, tasks chan Task){
	"hello": func(conn net.Conn, args []string, tasks chan Task) {
		conn.Write([]byte("Hello, world!\n"))
	},
	"master_req": func(conn net.Conn, args []string, tasks chan Task) {
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
		io.WriteString(conn, string(sol))
	},
}
