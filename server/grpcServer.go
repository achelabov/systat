package server

import (
	"log"
	"net"

	pb "github.com/achelabov/systat/proto"
	"google.golang.org/grpc"
)

type server struct {
	srv      *grpc.Server
	listener net.Listener
}

func NewServer() *server {
	return &server{
		srv: grpc.NewServer(),
	}
}

func (s *server) Listen(address string) error {
	listener, err := net.Listen("tcp", address)
	s.listener = listener

	log.Println("listening on: ", address)

	return err
}

func (s *server) Start() {
	pb.RegisterStatsServer(s.srv, pb.UnimplementedStatsServer{})

	log.Println("start server")

	if err := s.srv.Serve(s.listener); err != nil {
		log.Println("failed to serve: ", err)
	}
}

func (s *server) Close() {
	s.listener.Close()
}

//TODO
func (s *server) GetBatteries() error {

	return nil
}
