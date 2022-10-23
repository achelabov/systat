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
	conn *grpc.ClientConn
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
	client := pb.NewStatsClient(c.conn)
	stream, err := client.GetCpus(context.Background(), new(emptypb.Empty))
	if err != nil {
		log.Fatalf("openn stream error %v", err)
	}

	//ctx := stream.Context()
	done := make(chan bool)

	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				done <- true //close(done)
				return
			}
			if err != nil {
				log.Fatalf("can not receive %v", err)
			}
			log.Println("--------------------------------------------------")
			for i := 0; i < len(resp.Cpus); i++ {
				log.Printf("Resp received: cpus %d is %s ", i, resp.Cpus[i])
			}
			log.Println("--------------------------------------------------")
		}
	}()

	<-done
	log.Printf("finished")
}

func (c *grpcClient) Close() {
	c.conn.Close()
}
