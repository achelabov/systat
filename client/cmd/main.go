package main

import (
	"github.com/achelabov/systat/client"
)

func main() {
	c := client.NewClient()
	c.Dial()
	c.Start()
}
