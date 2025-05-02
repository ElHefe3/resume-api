package handler

import (
	"log"
	"net/http"
	"path"
	"strings"

	"github.com/ElHefe3/resume-api/pkg/nextcloud"
	"github.com/gin-gonic/gin"
)

func RetrieveFilesDirectories(c *gin.Context) {
	fullPaths, err := nextcloud.RetrieveFilesDirectories()
	if err != nil {
		log.Println("Error retrieving pages:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var files []string
	for _, p := range fullPaths {
		if strings.HasSuffix(p, "/") || path.Ext(p) != ".md" {
			continue
		}

		files = append(files, path.Base(p))
	}

	c.JSON(http.StatusOK, gin.H{
		"files": files,
	})
}
