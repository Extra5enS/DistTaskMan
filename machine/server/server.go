package server

import (
	"bufio"
	"encoding/json"
	"io"
	"net"
)

type Server struct {
	conn  *net.Conn
	tasks chan Task
}

func NewServer(address string, tasks chan Task) (Server, error) {
	ln, _ := net.Listen("tcp", address)
	conn, _ := ln.Accept()

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
			line, _, err := connReader.ReadLine()
			if err != nil {
				(*m.conn).Write([]byte(`{"error": "Can not read data"}`))
			}
			err = json.Unmarshal(line, json_map)
			if err != nil {
				(*m.conn).Write([]byte(`{"error": "Not a JSON data"}`))
			}
			if name_inter, ok := json_map["name"]; !ok {
				(*m.conn).Write([]byte(`{"error": "No name field in JSON"}`))
			} else if name, ok := name_inter.(string); !ok {
				(*m.conn).Write([]byte(`{"error": "Name field value is now a string"}`))
			} else if args_inter, ok := json_map["args"]; !ok {
				(*m.conn).Write([]byte(`{"error": "No args field in JSON"}`))
			} else if args, ok := args_inter.(map[string]interface{}); !ok {
				(*m.conn).Write([]byte(`{"error": "Args field value is now a string"}`))
			} else {
				handles[name](*m.conn, args, m.tasks)
			}
		}
	}(end)
}

var handles = map[string]func(conn net.Conn, args map[string]interface{}, tasks chan Task){
	"hello": func(conn net.Conn, args map[string]interface{}, tasks chan Task) {
		conn.Write([]byte(`{"respond":"Hello, world!}"`))
	},
	"master_req": func(conn net.Conn, args map[string]interface{}, tasks chan Task) {
		if address_inter, ok := args["address"]; !ok {
			(conn).Write([]byte(`{"error": "No address field in JSON"}`))
		} else if address, ok := address_inter.(string); !ok {
			(conn).Write([]byte(`{"error": "Address field value is now a string"}`))
		} else if num_inter, ok := args["args"]; !ok {
			(conn).Write([]byte(`{"error": "No num field in JSON"}`))
		} else if num, ok := num_inter.(string); !ok {
			(conn).Write([]byte(`{"error": "Num field value is now a string"}`))
		} else {
			sols := make(chan Sol)
			task := Task{
				Req:      "master_req",
				Args:     []string{address, num},
				Solution: sols,
			}
			tasks <- task

			sol := <-sols
			io.WriteString(conn, string(sol))
		}
	},
}
