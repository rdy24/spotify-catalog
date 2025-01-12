package tracks

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rdy24/spotify-catalog/internal/models/trackactivities"
)

func (h *Handler) UpsertTrackActivities(c *gin.Context) {
	ctx := c.Request.Context()

	var req trackactivities.TrackActivityRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.GetUint("userID")
	err := h.service.UpsertTrackActivities(ctx, userID, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}
