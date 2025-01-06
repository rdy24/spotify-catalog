package memberships

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rdy24/spotify-catalog/internal/models/memberships"
)

func (h *Handler) SignUp(c *gin.Context) {
	var (
		req memberships.SignUpRequest
	)

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request",
			"error":   err.Error(),
		})
		return
	}

	err := h.service.SignUp(req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to sign up user",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "successfully signed up user",
	})
}
