package cmd

import (
	httpServer "github.com/achelabov/systat/client/app/web/server"
	"github.com/spf13/cobra"
)

var webCmd = &cobra.Command{
	Use:   "web",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		app := httpServer.NewApp()
		app.Run("1234")
	},
}

func init() {
	rootCmd.AddCommand(webCmd)
}
