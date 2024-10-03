package handler

import (
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func downloadFile(url string, dest string) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	out, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, response.Body)
	return err
}

func (h *Handler) playAudio(c *gin.Context) {
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
	destPath := filepath.Join("web", "music", audio.Name+".mp3")
	err = downloadFile(audio.AudioDownload, destPath)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, audio)
}
