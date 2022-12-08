package server

type Server interface {
	Run(end chan error)
}
