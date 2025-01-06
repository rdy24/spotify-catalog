package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rdy24/spotify-catalog/internal/configs"
	membershipsHandler "github.com/rdy24/spotify-catalog/internal/handler/memberships"
	"github.com/rdy24/spotify-catalog/internal/models/memberships"
	membershipRepository "github.com/rdy24/spotify-catalog/internal/repository/memberships"
	membershipSvc "github.com/rdy24/spotify-catalog/internal/service/memberships"
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

	db.AutoMigrate(&memberships.User{})

	r := gin.Default()

	membershipRepo := membershipRepository.NewRepository(db)

	membershipService := membershipSvc.NewService(cfg, membershipRepo)

	membershipHandler := membershipsHandler.NewHandler(r, membershipService)

	membershipHandler.RegisterRoutes()

	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	r.Run(cfg.Service.Port)
}
