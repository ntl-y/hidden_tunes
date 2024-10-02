package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) audioPlay(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"random_audio": "audio",
	})
}
