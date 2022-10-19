package client

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type TCPClient struct {
	conn     net.Conn
	name     string
	incoming chan string
	error    chan error
}

func NewClient() *TCPClient {
	return &TCPClient{
		incoming: make(chan string, 5),
	}
}

func (c *TCPClient) Dial(address string) error {
	addr, _ := net.ResolveTCPAddr("tcp", address)
	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		log.Fatal(err)
	}

	c.conn = conn

	return err
}

func (c *TCPClient) Start() {
	for {
		data, err := bufio.NewReader(c.conn).ReadString('\n')
		if err != nil {
			c.error <- err
			return
		}
		c.incoming <- data
		fmt.Println(<-c.incoming)
	}
}

func (c *TCPClient) Incoming() chan string {
	return c.incoming
}

func (c *TCPClient) Error() chan error {
	return c.error
}

func (c *TCPClient) Close() {
	c.conn.Close()
}
