package handler

import (
	"net/http"

	"github.com/ElHefe3/resume-api/internal/services/nextcloud"
	"github.com/gin-gonic/gin"
)

func ServeHTML(c *gin.Context, filePath string) {
	html, _, err := nextcloud.RetrieveFile(filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Data(http.StatusOK, "text/html; charset=utf-8", html)
}
