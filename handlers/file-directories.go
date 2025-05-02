package handler

import (
	"log"
	"net/http"

	"github.com/ElHefe3/resume-api/pkg/nextcloud"
	"github.com/gin-gonic/gin"
)

func RetrieveFilesDirectories(c *gin.Context) {
	files, err := nextcloud.RetrieveFilesDirectories()
	if err != nil {
		log.Println("Error retrieving pages:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"files": files,
	})
}
