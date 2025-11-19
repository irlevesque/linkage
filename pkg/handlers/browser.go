package handlers

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func BrowserRequest(c *gin.Context) {
	term := c.Query("s")

	if term == "" {
		c.String(http.StatusBadRequest, "Missing search term")
		return
	}

	// isolate the search term from the query string
	var stub, args string
	if strings.Contains(term, "/") {
		stub = term[:strings.Index(term, "/")]
		args = term[strings.Index(term, "/"):]
		log.Printf("Evaluating stub %s and more %s\n", stub, args)
	} else {
		stub = term
		args = ""
	}

	_, l := findLink(stub)
	if l != nil {
		c.Redirect(http.StatusMovedPermanently, l.URL+args)
		return
	}

	err, fallback := SearchEngineQueryURL(term)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error: %s", err)
		return
	}
	c.Redirect(http.StatusMovedPermanently, fallback)
}
