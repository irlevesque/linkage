package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Query Endpoints
	router.GET("/", browserRequest)

	// API Endpoints
	{
		api := router.Group("/api")
		// links
		api.GET("/links", getLinks)
		api.GET("/link/:stub", getLink)
		api.POST("/link/:stub", addLink)
		api.DELETE("/link/:stub", deleteLink)

		// search handlers
		api.GET("/search/all", getSearchEngines)
		api.GET("/search/current", getCurrentSearchEngine)
		api.POST("/search", addSearchEngine)
	}

	router.Run(":8080")
}
