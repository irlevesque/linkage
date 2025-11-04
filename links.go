package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type link struct {
	Stub        string    `json:"stub" form:"stub" binding:"required"`
	URL         string    `json:"url" form:"url" binding:"required"`
	Description string    `json:"description" form:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

var bday = time.Date(2025, time.November, 3, 0, 0, 0, 0, time.UTC)

// links defines the default slice of links
var links = []link{
	{
		Stub:        "gh",
		URL:         "https://github.com",
		Description: "GitHub",
		CreatedAt:   bday,
	},
	{
		Stub:        "linkage",
		URL:         "https://github.com/irlevesque/linkage",
		Description: "Linkage",
		CreatedAt:   bday,
	},
}

func findLink(stub string) (int, *link) {
	for i, l := range links {
		if l.Stub == stub {
			return i, &l
		}
	}
	return -1, nil
}

// TODO: may want to refactor this if O(n) becomes too slow
func deleteLink(c *gin.Context) {
	reqLink := c.Param("stub")
	i, l := findLink(reqLink)
	if l == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Link not found"})
		return
	}
	links = append(links[:i], links[i+1:]...)
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Link deleted", "link": l})
}

func getLinks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, links)
}

func getLink(c *gin.Context) {
	reqLink := c.Param("stub")
	_, l := findLink(reqLink)
	if l == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Link not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, l)
}

func addLink(c *gin.Context) {
	var newLink = link{Stub: c.Param("stub")}
	_, l := findLink(newLink.Stub)
	if l != nil {
		c.IndentedJSON(http.StatusConflict, gin.H{"error": "Link already exists", "link": l})
		return
	}

	if err := c.ShouldBindJSON(&newLink); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newLink.CreatedAt = time.Now()
	links = append(links, newLink)
	c.IndentedJSON(http.StatusCreated, newLink)
}
