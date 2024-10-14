package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) playAudio(c *gin.Context) {
	c.HTML(http.StatusOK, "play.html", nil)
}

func (h *Handler) home(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func (h *Handler) getRandomAudio(c *gin.Context) {
	audio, err := h.service.GetRandomAudio()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, audio)
}
