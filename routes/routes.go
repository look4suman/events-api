package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/look4suman/events-api/middlewares"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEventById)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)

	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)

	server.GET("/users", fetchUsers)
	server.POST("/signup", signup)
	server.POST("/login", login)
}
