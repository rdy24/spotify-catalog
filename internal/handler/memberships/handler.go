package memberships

import (
	"github.com/gin-gonic/gin"
	"github.com/rdy24/spotify-catalog/internal/models/memberships"
)

//go:generate mockgen -source=handler.go -destination=handler_mock_test.go -package=memberships
type service interface {
	SignUp(request memberships.SignUpRequest) error
	Login(request memberships.LoginRequest) (string, error)
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
	route := h.Group("/memberships")

	route.POST("/sign-up", h.SignUp)
	route.POST("/login", h.Login)
}
