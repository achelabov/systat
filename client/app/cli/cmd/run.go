package cmd

import (
	"fmt"
	"os"

	"github.com/achelabov/systat/client"
	cfg "github.com/achelabov/systat/client/config"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cfg.Init()

		c := client.NewClient()
		c.Dial(os.Getenv("HOST") + ":" + os.Getenv("PORT"))
		c.Start()

		for stats := range c.Receive() {
			fmt.Println(stats)
		}

		c.Close()
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
