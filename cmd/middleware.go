package cmd

import (
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func MiddlewareCORS() gin.HandlerFunc {
	log.Println("Middleware CORS Aktif") // Logging untuk memastikan middleware berjalan
	return cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3039"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
}
