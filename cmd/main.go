package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rdy24/spotify-catalog/internal/configs"
	"github.com/rdy24/spotify-catalog/pkg/internalsql"
)

func main() {
	var (
		cfg *configs.Config
	)

	err := configs.Init(
		configs.WithConfigFolders([]string{"./internal/configs"}),
		configs.WithConfigFile("config"),
		configs.WithConfigType("yaml"),
	)

	if err != nil {
		log.Fatalf("failed to initialize config: %v", err)
	}

	cfg = configs.Get()

	db, err := internalsql.Connect(cfg.Database.DataSourceName)

	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	r := gin.Default()
	r.Run(cfg.Service.Port)
}
