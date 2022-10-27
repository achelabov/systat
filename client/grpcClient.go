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

func (c *grpcClient) Recieve() {
	//cpusStream, err := c.client.GetCpus(context.Background(), new(emptypb.Empty))
	//if err != nil {
	//	log.Fatalf("open cpus stream error %v", err)
	//}
	//battsStream, err := c.client.GetBatteries(context.Background(), new(emptypb.Empty))
	//if err != nil {
	//	log.Fatalf("open batts stream error %v", err)
	//}

	stream, err := c.client.GetStats(context.Background(), new(emptypb.Empty))
	if err != nil {
		log.Fatalf("open stream error %v", err)
	}

	//ctx := stream.Context()
	done := make(chan struct{})

	getStats := func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				log.Println("_EOF_")
				done <- struct{}{} //close(done)
				return
			}
			if err != nil {
				log.Fatalf("can not receive %v", err)
			}
			log.Println("--------------------------------------------------")
			for i := 0; i < len(resp.Cpus.Cpus); i++ {
				log.Printf("cpus %d load is %s ", i, resp.Cpus.Cpus[i])
			}
			log.Println("average cpus load: ", resp.Cpus.AverageLoad)
			log.Println("--------------------------------------------------")

			log.Println("--------------------------------------------------")
			for i := 0; i < len(resp.Batteries.Batteries); i++ {
				log.Printf("batt %d load is %f, state is %s",
					i, resp.Batteries.Batteries[i].BatteryLoad, resp.Batteries.Batteries[i].State)
			}
			log.Println("--------------------------------------------------")
		}
	}

	//getBatts := func(done chan<- struct{}) {
	//	for {
	//		resp, err := battsStream.Recv()
	//		if err == io.EOF {
	//			done <- struct{}{} //close(done)
	//			return
	//		}
	//		if err != nil {
	//			log.Fatalf("can not receive %v", err)
	//		}
	//		log.Println("--------------------------------------------------")
	//		for i := 0; i < len(resp.Batteries); i++ {
	//			log.Printf("batt %d load is %f, state is %s",
	//				i, resp.Batteries[i].BatteryLoad, resp.Batteries[i].State)
	//		}
	//		log.Println("--------------------------------------------------")
	//	}
	//}

	go getStats()
	//go getBatts(done)

	//<-done
	<-done
	log.Printf("finished")
}
