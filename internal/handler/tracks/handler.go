package tracks

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/rdy24/spotify-catalog/internal/models/spotify"
)

//go:generate mockgen -source=handler.go -destination=handler_mock_test.go -package=tracks
type service interface {
	Search(ctx context.Context, query string, pageSize, pageIndex int) (*spotify.SearchResponse, error)
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

	route.GET("/search", h.Search)
}
