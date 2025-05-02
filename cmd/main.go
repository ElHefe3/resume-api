package main

import (
	"time"

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
		AllowOrigins: []string{
			"http://192.168.1.157:4173", // SPA dev host
			// add http://<temp‑server>:<port> if you host UI elsewhere
		},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,            // only if your fetch() sends cookies / auth header
		MaxAge:           12 * time.Hour,  // cache pre‑flight
	}))

	// routes
	r.GET("/directories", handler.RetrieveFilesDirectories)
	r.GET("/file", handler.ServeMarkdownPage)

	r.Run(":8900")
}
