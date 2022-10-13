package server

import (
	"fmt"
	"log"
	"net"
	"sync"
)

type TCPServer struct {
	listener net.Listener
	clients  []*client
	mutex    *sync.Mutex
}

type client struct {
	conn net.Conn
	name string
}

func NewTCPServer() *TCPServer {
	return &TCPServer{
		mutex: &sync.Mutex{},
	}
}

func (s *TCPServer) Listen(address string) error {
	listener, err := net.Listen("tcp", address)
	s.listener = listener

	fmt.Println("listening on: ", address)

	return err
}

func (s *TCPServer) Start() {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			log.Fatal("tcp server accept error: ", err)
		}

		go s.Broadcast(conn)
	}
}

func (s *TCPServer) Close() {
	s.listener.Close()
}

func (s *TCPServer) Broadcast(conn net.Conn) {
	defer conn.Close()

	//	for {
	//		TODO: send stream to client
	//	}
}