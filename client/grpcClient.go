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
	cpusStream, err := client.GetCpus(context.Background(), new(emptypb.Empty))
	if err != nil {
		log.Fatalf("openn stream error %v", err)
	}
	battsStream, err := client.GetBatteries(context.Background(), new(emptypb.Empty))
	if err != nil {
		log.Fatalf("openn stream error %v", err)
	}

	//ctx := stream.Context()
	done := make(chan bool)

	getCpus := func(done chan<- bool) {
		for {
			resp, err := cpusStream.Recv()
			if err == io.EOF {
				done <- true //close(done)
				return
			}
			if err != nil {
				log.Fatalf("can not receive %v", err)
			}
			log.Println("--------------------------------------------------")
			for i := 0; i < len(resp.Cpus); i++ {
				log.Printf("cpus %d load is %s ", i, resp.Cpus[i])
			}
			log.Println("average cpus load: ", resp.AverageLoad)
			log.Println("--------------------------------------------------")
		}
	}

	getBatts := func(done chan<- bool) {
		for {
			resp, err := battsStream.Recv()
			if err == io.EOF {
				done <- true //close(done)
				return
			}
			if err != nil {
				log.Fatalf("can not receive %v", err)
			}
			log.Println("--------------------------------------------------")
			for i := 0; i < len(resp.Batteries); i++ {
				log.Printf("batt %d load is %f, state is %s",
					i, resp.Batteries[i].BatteryLoad, resp.Batteries[i].State)
			}
			log.Println("--------------------------------------------------")
		}
	}

	go getCpus(done)
	go getBatts(done)

	<-done
	<-done
	log.Printf("finished")
}

func (c *grpcClient) Close() {
	c.conn.Close()
}
