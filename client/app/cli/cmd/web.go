package cmd

import (
	"log"
	"net/url"
	"os"
	"os/signal"

	"github.com/achelabov/systat/client"
	httpServer "github.com/achelabov/systat/client/app/web/server"
	cfg "github.com/achelabov/systat/client/config"
	pb "github.com/achelabov/systat/proto"
	"github.com/gorilla/websocket"
	"github.com/spf13/cobra"
)

var webCmd = &cobra.Command{
	Use:   "web",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		app := httpServer.NewApp()

		done := make(chan struct{})
		go func() {
			app.Run("8080")
			done <- struct{}{}
		}()

		cfg.Init()

		c := client.NewClient()
		c.Dial(os.Getenv("HOST") + ":" + os.Getenv("PORT"))
		c.Start()

		host := os.Getenv("HOST") + ":" + "8080"
		url := url.URL{Scheme: "ws", Host: host, Path: "/echo"}
		wsClient, _, err := websocket.DefaultDialer.Dial(url.String(), nil)
		if err != nil {
			log.Println("ws dial error: ", err)
			return
		}

		statsCh := make(chan *pb.StatsResponse, 1)

		go func() {
			defer close(statsCh)

			for stats := range c.Receive() {
				statsCh <- stats
			}
		}()

		interrupt := make(chan os.Signal, 1)
		signal.Notify(interrupt, os.Interrupt)

		for {
			select {
			case stats := <-statsCh:
				//log.Println(stats)
				err := wsClient.WriteMessage(websocket.TextMessage, []byte(stats.String()))
				if err != nil {
					log.Println("ws write error: ", err)
					return
				}
			case <-interrupt:
				log.Println("interrupt")

				err := wsClient.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
				if err != nil {
					log.Println("write close:", err)
					return
				}
			case <-done:
				return
			}
		}

		<-done
		c.Close()
	},
}

func init() {
	rootCmd.AddCommand(webCmd)
}
