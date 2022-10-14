package server

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	"github.com/achelabov/systat/server/widgets"
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

	wdgts := widgets.NewWidgets()

	ticker := time.NewTicker(time.Second)
	for {
		var battBuffer bytes.Buffer
		for i, v := range wdgts.GetBatteries() {
			battBuffer.WriteString(fmt.Sprintln("batt id: ", i, "load: ", v.PercentFull, "state: ", v.State))
		}
		conn.Write(battBuffer.Bytes())

		var cpuBuffer bytes.Buffer
		for i, v := range wdgts.GetCpus() {
			cpuBuffer.WriteString(fmt.Sprintln("cpu id: ", i, "load: ", v.CpuLoad))
		}
		conn.Write(cpuBuffer.Bytes())

		<-ticker.C
	}
}
