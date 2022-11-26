package main

import (
	"exercise/internal/app/database"
	"exercise/internal/app/exercise/handler"
	userHandler "exercise/internal/app/user/handler"
	"exercise/internal/pkg/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Main function
func main() {
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]string{
			"message": "Hello world",
		})
	})

	// Define db connection
	db := database.NewConnDatabase()
	exerciseHandler := handler.NewExerciseHandler(db)
	userHandler := userHandler.NewUserHandler(db)

	// GET request
	r.GET("/exercises/:id", middleware.WithAuh(), exerciseHandler.GetExerciseByID)
	r.GET("/exercises/:id/score", middleware.WithAuh(), exerciseHandler.GetScore)

	// POST request
	r.POST("/exercises", middleware.WithAuh(), exerciseHandler.CreateExercise)
	r.POST("/exercises/:id/questions", middleware.WithAuh(), exerciseHandler.CreateQuestion)
	r.POST("/exercises/:id/questions/:qid/answer", middleware.WithAuh(), exerciseHandler.CreateAnswer)
	r.POST("/register", userHandler.Register)

	// Define localhost port
	r.Run(":1234")
}
