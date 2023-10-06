package cmd

import (
	"blog-system/pkg/bootstrap"
	"github.com/spf13/cobra"
	"log"
)

func init() {
	rootCmd.AddCommand(migrateCmd)
}

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Table migration",
	Long:  "Table migrate",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Migrating database tables")
		migrate()
	},
}

func migrate() {
	bootstrap.Migrate()
}
