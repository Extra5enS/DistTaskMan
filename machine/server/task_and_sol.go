package server

type Task struct {
	Req  string
	Args []string

	Solution chan Sol
}

type Sol string
