package cmd

import (
	"github.com/spf13/cobra"
)

var webCmd = &cobra.Command{
	Use:   "",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCmd.AddCommand(webCmd)
}
