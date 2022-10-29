package client

import (
	"context"
	"io"
	"log"

	pb "github.com/achelabov/systat/proto"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type grpcClient struct {
	conn   *grpc.ClientConn
	client pb.StatsServiceClient
}

func NewClient() *grpcClient {
	return &grpcClient{}
}

func (c *grpcClient) Dial(address string) error {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	c.conn = conn

	return err
}

func (c *grpcClient) Start() {
	c.client = pb.NewStatsServiceClient(c.conn)
}

func (c *grpcClient) Close() {
	c.conn.Close()
}

func (c *grpcClient) Receive() <-chan *pb.StatsResponse {
	out := make(chan *pb.StatsResponse)
	stream, err := c.client.GetStats(context.Background(), new(emptypb.Empty))
	if err != nil {
		log.Fatalf("open stream error %v", err)
	}

	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				log.Println("__EOF__")
				return
			}
			if err != nil {
				log.Fatalf("can not receive %v", err)
			}

			out <- resp
		}
	}()

	return out
}
