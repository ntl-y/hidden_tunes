package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) audioPlay(c *gin.Context) {
	audio, err := h.service.GetRandomAudio()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, audio)
}
