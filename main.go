package main

import (
	"github.com/gin-gonic/gin"
	"rest-api.com/db"
	"rest-api.com/handler"
	middlewares "rest-api.com/middleswares"
)

func main() {
	db.ConnectDatabase()

	server := gin.Default() // Creates a Default http server for use

	//------------------------ Events Routes -----------------------------//

	server.GET("/", handler.GetEvents)
	server.GET("/events/:id", handler.GetEvent)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)

	authenticated.POST("/events", handler.CreateEvent)
	authenticated.PUT("/events/:id", handler.UpdateEvent)
	authenticated.DELETE("/events/:id", handler.RemoveEvent)

	//------------------------- User Routes --------------------------------//

	server.POST("/signup", handler.CreateUser)
	server.POST("/login", handler.Login)

	err := server.Run(":8080") // in development -> Localhost. Essentially the address of the server
	if err != nil {
		return
	}
}
