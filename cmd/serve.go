package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve app on the server",
	Long:  "Application will be served on host and port defined on config.yml",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Application will up using this command, disabled for now!!!")
	},
}
