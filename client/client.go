package client

type ChatClient interface {
	Dial(address string) error
	Start()
	Close()
	Incoming() chan string
}
