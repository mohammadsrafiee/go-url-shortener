package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	shortenerConfig "url-shortener/internal/shortener"
)

var items []shortenerConfig.ShortenerConfig

// @summary This handler return all shortener configs.
// @description Administrator can define shortener config and manipulate them. profiling for this usecase is not exit
// @produce json
// @success 200 {object} ShortenerConfig
// @router /api/shortener-config [get]
func GetShortenerConfigs(c *gin.Context) {
	c.JSON(http.StatusOK, items)
}

func GetItem(c *gin.Context) {
	id := c.Param("id")
	for _, item := range items {
		if id == item.ID {
			c.JSON(http.StatusOK, item)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Item not found"})
}

func CreateItem(c *gin.Context) {
	var item shortenerConfig.ShortenerConfig
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	items = append(items, item)
	c.JSON(http.StatusCreated, item)
}

func UpdateItem(c *gin.Context) {
	id := c.Param("id")
	var updatedItem shortenerConfig.ShortenerConfig
	if err := c.ShouldBindJSON(&updatedItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for i, item := range items {
		if id == item.ID {
			updatedItem.ID = item.ID
			items[i] = updatedItem
			c.JSON(http.StatusOK, updatedItem)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Item not found"})
}

func DeleteItem(c *gin.Context) {
	id := c.Param("id")
	for i, item := range items {
		if id == item.ID {
			items = append(items[:i], items[i+1:]...)
			c.JSON(http.StatusOK, item)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Item not found"})
}
