package client

import (
	"log"
	"net"
)

type TCPClient struct {
	conn     net.Conn
	name     string
	incoming chan string
}

func NewClient() *TCPClient {
	return &TCPClient{
		incoming: make(chan string),
	}
}

func (c *TCPClient) Dial() error {
	addr, _ := net.ResolveTCPAddr("tcp", ":1337")
	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		log.Fatal(err)
	}

	c.conn = conn

	return err
}

//TODO
func (c *TCPClient) Start() {

}

func (c *TCPClient) Close() {
	c.conn.Close()
}