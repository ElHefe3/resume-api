package routes

import (
	handler "github.com/ElHefe3/resume-api/internal/handlers"
	"github.com/ElHefe3/resume-api/internal/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
    r.Use(middleware.AuthMiddleware())

    r.GET("/file", handler.FileHandler)
    r.GET("/directories", handler.RetrieveFilesDirectories)
}
