package server

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net"
)

type Server struct {
	address string
	tasks   chan Task
}

func NewServer(address string, tasks chan Task) (Server, error) {
	return Server{address: address, tasks: tasks}, nil
}

func (m Server) Run(end chan error) {
	go func(end chan error) {
		ln, err := net.Listen("tcp", m.address)
		if err != nil {
			fmt.Println(err)
			return
		}
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		for {
			json_map := make(map[string]interface{})

			connReader := bufio.NewReader(conn)
			fmt.Println("listen")

			line, _, err := connReader.ReadLine()
			fmt.Printf("listen: %s\n", line)
			if err != nil {
				fmt.Println(`{"error": "Can not read data"}` + "\n")
				(conn).Write([]byte(`{"error": "Can not read data"}` + "\n"))
			}
			err = json.Unmarshal(line, &json_map)
			if err != nil {
				fmt.Println(`{"error": "Not a JSON data"}\n`)
				(conn).Write([]byte(`{"error": "Not a JSON data"}` + "\n"))
			}
			if name_inter, ok := json_map["name"]; !ok {
				fmt.Println(`{"error": "No name field in JSON"}\n`)
				(conn).Write([]byte(`{"error": "No name field in JSON"}` + "\n"))
			} else if name, ok := name_inter.(string); !ok {
				fmt.Println(`{"error": "Name field value is now a string"}` + "\n")
				(conn).Write([]byte(`{"error": "Name field value is now a string"}` + "\n"))
			} else if args_inter, ok := json_map["args"]; !ok {
				fmt.Println(`{"error": "No args field in JSON"}` + "\n")
				(conn).Write([]byte(`{"error": "No args field in JSON"}` + "\n"))
			} else if args, ok := args_inter.(map[string]interface{}); !ok {
				fmt.Println(`{"error": "Args field value is now a string"}` + "\n")
				(conn).Write([]byte(`{"error": "Args field value is now a string"}` + "\n"))
			} else {
				fmt.Println(name, args)
				handles[name](conn, args, m.tasks)
			}
		}
	}(end)
}

var handles = map[string]func(conn net.Conn, args map[string]interface{}, tasks chan Task){
	"hello": func(conn net.Conn, args map[string]interface{}, tasks chan Task) {
		conn.Write([]byte(`{"respond":"Hello, world!}"` + "\n"))
	},
	"master_req": func(conn net.Conn, args map[string]interface{}, tasks chan Task) {
		if address_inter, ok := args["address"]; !ok {
			(conn).Write([]byte(`{"error": "No address field in JSON"}` + "\n"))
		} else if address, ok := address_inter.(string); !ok {
			(conn).Write([]byte(`{"error": "Address field value is now a string"}` + "\n"))
		} else if num_inter, ok := args["num"]; !ok {
			(conn).Write([]byte(`{"error": "No num field in JSON"}` + "\n"))
		} else if num, ok := num_inter.(string); !ok {
			(conn).Write([]byte(`{"error": "Num field value is now a string"}` + "\n"))
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
