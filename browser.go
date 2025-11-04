package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func browserRequest(c *gin.Context) {
	term := c.Query("s")

	_, l := findLink(term)
	if l != nil {
		c.Redirect(http.StatusMovedPermanently, l.URL)
		return
	}

	err, fallback := searchEngineQueryURL(term)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error: %s", err)
		return
	}
	c.Redirect(http.StatusMovedPermanently, fallback)
}
