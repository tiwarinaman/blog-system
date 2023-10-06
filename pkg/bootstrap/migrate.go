package bootstrap

import (
	"blog-system/internal/database/migration"
	"blog-system/pkg/config"
	"blog-system/pkg/database"
)

func Migrate() {
	config.Set()

	database.Connect()

	migration.Migrate()
}
