package handler

import (
	"log"
	"net/http"

	"github.com/ElHefe3/resume-api/pkg/utils"
	"github.com/gin-gonic/gin"
)

func FileHandler(c *gin.Context) {
    filePath := c.Query("path")
    if filePath == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "path query param is required"})
        return
    }

	fileType := utils.GetFileType(filePath);

	log.Println("🤖: ", fileType)

    switch fileType {
    case "markdown":
        ServeMarkdown(c, filePath)
    case "html":
        ServeHTML(c, filePath)
    default:
        c.JSON(http.StatusUnsupportedMediaType, gin.H{"error": "unsupported file type"})
    }
}
