package main

import (
	"github.com/ElHefe3/resume-api/config"
	handler "github.com/ElHefe3/resume-api/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.Load()

	r := gin.Default()

	// 🔑  CORS policy
	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,          // OR: AllowOrigins: []string{"*"}
		AllowMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: false,        // must be false when AllowAllOrigins is true
	}))

	// routes
	r.GET("/directories", handler.RetrieveFilesDirectories)
	r.GET("/file", handler.ServeMarkdownPage)

	r.Run(":8900")
}
