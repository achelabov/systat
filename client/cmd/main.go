package main

import (
	"os"

	"github.com/achelabov/systat/client"
	cfg "github.com/achelabov/systat/client/config"
)

func main() {
	cfg.Init()

	c := client.NewClient()
	c.Dial(os.Getenv("HOST") + ":" + os.Getenv("PORT"))
	c.Start()
	c.Close()
}
