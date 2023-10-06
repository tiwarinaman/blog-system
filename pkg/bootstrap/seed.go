package bootstrap

import (
	"blog-system/internal/database/seeder"
	"blog-system/pkg/config"
	"blog-system/pkg/database"
)

func Seed() {
	config.Set()

	database.Connect()

	seeder.Seed()
}
