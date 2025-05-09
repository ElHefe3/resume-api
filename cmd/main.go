package main

import (
	"fmt"

	"github.com/ElHefe3/resume-api/internal/config"
	"github.com/ElHefe3/resume-api/internal/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.Load()

	r := gin.Default()

	r.Use(func(c *gin.Context) {
		fmt.Printf(">> %s %s  Origin=%s\n",
			c.Request.Method, c.Request.URL.Path, c.Request.Header.Get("Origin"))
		c.Next()
	})
	

	// 🔑  CORS policy
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{config.Cfg.FeUrl},
		AllowMethods:     []string{"GET", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))

	// routes
	routes.RegisterRoutes(r)

	r.Run(":44000")
}
