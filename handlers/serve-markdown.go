package handler

import (
	"net/http"
	"strings"

	"github.com/ElHefe3/resume-api/pkg/nextcloud"
	"github.com/gin-gonic/gin"
	"github.com/gomarkdown/markdown"
)

func ServeMarkdownPage(c *gin.Context) {
	filePath := c.Query("path")
	if filePath == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "path query param is required"})
		return
	}

	if !strings.HasSuffix(filePath, ".md") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "only .md files supported"})
		return
	}

	content, _, err := nextcloud.RetrieveFile(filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	html := markdown.ToHTML(content, nil, nil)

	c.Data(http.StatusOK, "text/html; charset=utf-8", html)
}
