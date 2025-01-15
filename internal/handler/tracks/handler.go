package tracks

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/rdy24/spotify-catalog/internal/middleware"
	"github.com/rdy24/spotify-catalog/internal/models/spotify"
	"github.com/rdy24/spotify-catalog/internal/models/trackactivities"
)

//go:generate mockgen -source=handler.go -destination=handler_mock_test.go -package=tracks
type service interface {
	Search(ctx context.Context, query string, pageSize, pageIndex int, userID uint) (*spotify.SearchResponse, error)
	UpsertTrackActivities(ctx context.Context, userID uint, request trackactivities.TrackActivityRequest) error
	GetRecommendation(ctx context.Context, userID uint, limit int, trackID string) (*spotify.RecommendationResponse, error)
}

type Handler struct {
	*gin.Engine
	service service
}

func NewHandler(api *gin.Engine, service service) *Handler {
	return &Handler{
		api,
		service,
	}
}

func (h *Handler) RegisterRoutes() {
	route := h.Group("/tracks")
	route.Use(middleware.AuthMiddleware())
	route.GET("/search", h.Search)
	route.POST("/track-activities", h.UpsertTrackActivities)
	route.GET("/recommendation", h.GetRecommendation)
}
