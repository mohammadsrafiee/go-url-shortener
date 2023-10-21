package handlers

import (
	"github.com/gin-gonic/gin"
)

func NewShorterConfigRoutes(engine *gin.Engine) {
	// Define your API routes
	engine.GET("/api/v1/shortener-config", GetShortenerConfigs)
	engine.GET("/api/v1/items/:id", GetItem)
	engine.POST("/api/v1/items", CreateItem)
	engine.PUT("/api/v1/items/:id", UpdateItem)
	engine.DELETE("/api/v1/items/:id", DeleteItem)
}
