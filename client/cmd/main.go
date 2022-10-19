package main

import (
	"fmt"
	"log"
	"os"

	"github.com/achelabov/systat/client"
	cfg "github.com/achelabov/systat/client/config"
)

func main() {
	cfg.Init()

	c := client.NewClient()
	c.Dial(os.Getenv("HOST") + ":" + os.Getenv("PORT"))
	c.Start()
	for {
		select {
		case err := <-c.Error():
			log.Fatal(err)
		case msg := <-c.Incoming():
			fmt.Println(msg)
		}
	}
}
