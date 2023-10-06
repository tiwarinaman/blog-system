package cmd

import (
	"blog-system/pkg/bootstrap"
	"github.com/spf13/cobra"
	"log"
)

func init() {
	rootCmd.AddCommand(seedCmd)
}

var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "Database seeder",
	Long:  "Database seeeder",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Database seeding")
		seed()
	},
}

func seed() {
	bootstrap.Seed()
}
