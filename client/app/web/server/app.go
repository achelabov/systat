package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	handler "github.com/achelabov/systat/client/app/web/controller"
)

type app struct {
	httpServer *http.Server
}

func NewApp() *app {
	return &app{}
}

func (a *app) Run(port string) error {
	router := initRouter()

	a.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := a.httpServer.ListenAndServe(); err != nil {
			log.Fatalf("Failed to listen and serve: %+v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return a.httpServer.Shutdown(ctx)
}

func initRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handler.GetStatistics)

	return mux
}
