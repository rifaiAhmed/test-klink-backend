package cmd

import (
	"backend-test/helpers"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ServeHTTP menjalankan server HTTP
func ServeHTTP() {
	dependency := dependencyInject()

	r := gin.Default()
	r.Use(MiddlewareCORS())

	// Healthcheck
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "OK"})
	})

	member := r.Group("/members")
	{
		member.POST("", dependency.MemberAPI.CreateMember)
		member.GET("/:id", dependency.MemberAPI.GetMemberByID)
		// member.PUT("/:id", dependency.MemberAPI.UpdateMember)
		member.DELETE("/:id", dependency.MemberAPI.DeleteMember)
		member.GET("", dependency.MemberAPI.GetAllMembers)
	}

	port := helpers.GetEnv("PORT", "8080")
	log.Printf("Server running on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
