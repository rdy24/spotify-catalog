package memberships

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rdy24/spotify-catalog/internal/models/memberships"
)

func (h *Handler) Login(c *gin.Context) {
	var req memberships.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request",
			"error":   err.Error(),
		})
		return
	}

	accessToken, err := h.service.Login(req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, memberships.LoginResponse{
		AccessToken: accessToken,
	})
}
