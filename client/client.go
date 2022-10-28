package client

import pb "github.com/achelabov/systat/proto"

type ChatClient interface {
	Dial(address string) error
	Start()
	Recieve() <-chan *pb.StatsResponse
	Close()
}
