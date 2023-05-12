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

	// GET index.html
	engine.GET("/", func(c *gin.Context) {
		// Response
		c.File("./static/index.html")
	})

	// GET favicon.ico
	engine.GET("/favicon.ico", func(c *gin.Context) {
		// Response
		c.File("./static/favicon.ico")
	})

	// Serve static files
	engine.Static("/assets", "./static/assets")

	// Start server
	color.Yellow.Println("Server starting on: ", ServerUrl)
	engine.Run(ServerUrl)
}
