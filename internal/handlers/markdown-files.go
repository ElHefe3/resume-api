package handler

import (
	"log"
	"net/http"

	"github.com/ElHefe3/resume-api/internal/services/nextcloud"
	"github.com/gin-gonic/gin"
	"github.com/gomarkdown/markdown"
)

func ServeMarkdown(c *gin.Context, filePath string) {
	content, _, err := nextcloud.RetrieveFile(filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	html := markdown.ToHTML(content, nil, nil)

	log.Println(html)

	c.Data(http.StatusOK, "text/html; charset=utf-8", html)
}
