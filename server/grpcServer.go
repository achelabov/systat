package server

import (
	"log"
	"net"

	pb "github.com/achelabov/systat/proto"
	"google.golang.org/grpc"
)

type grpcServer struct {
	srv      *grpc.Server
	listener net.Listener
}

func NewGrpcServer() *grpcServer {
	return &grpcServer{
		srv: grpc.NewServer(),
	}
}

func (s *grpcServer) Listen(address string) error {
	listener, err := net.Listen("tcp", address)
	s.listener = listener

	log.Println("listening on: ", address)

	return err
}

func (s *grpcServer) Start() {
	pb.RegisterStatsServiceServer(s.srv, &statsServer{})

	log.Println("start server")

	if err := s.srv.Serve(s.listener); err != nil {
		log.Println("failed to serve: ", err)
	}
}

func (s *grpcServer) Close() {
	s.listener.Close()
	s.srv.Stop()
}
