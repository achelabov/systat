package server

type Server interface {
	Listen(address string) error
	Start()
	Close()
}
