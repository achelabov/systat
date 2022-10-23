package client

type ChatClient interface {
	Dial(address string) error
	Start()
	Close()
}
