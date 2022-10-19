package main

import (
	"log"
	"os"

	"github.com/achelabov/systat/server"
	cfg "github.com/achelabov/systat/server/config"
)

func main() {
	cfg.Init()

	s := server.NewServer()

	err := s.Listen(os.Getenv("HOST") + ":" + os.Getenv("PORT"))
	if err != nil {
		log.Fatal("tcp server listener error:", err)
	}

	s.Start()
	s.Close()
}
