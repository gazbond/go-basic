package server

import (
	"github.com/gin-gonic/gin"
	"github.com/gookit/color"
)

// Start server
func Start() {

	// Gin engine
	engine := gin.New()

	// Config global middleware
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())

	// Serve static files
	engine.Static("/", "./static/")

	// Start server
	color.Yellow.Println("Server starting on: ", ServerUrl)
	engine.Run(ServerUrl)
}
