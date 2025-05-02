package main

import (
	"github.com/ElHefe3/resume-api/config"
	handler "github.com/ElHefe3/resume-api/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	config.Load()

	r := gin.Default()

	r.GET("/directories", handler.RetrieveFilesDirectories)
	r.GET("/file", handler.ServeMarkdownPage)

	r.Run(":8900")
}
