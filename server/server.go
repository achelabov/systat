package server

import "net"

type Server interface {
	Listen(address string) error
	Broadcast(conn net.Conn) error
	Start()
	Close()
}
