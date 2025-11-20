package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Link struct {
	Stub        string    `json:"stub" form:"stub" binding:"required"`
	URL         string    `json:"url" form:"url" binding:"required"`
	Description string    `json:"description,omitempty" form:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

// LinkConfigWriter defines the interface for updating configuration files.
type LinkConfigWriter interface {
	SaveLinks(links []*Link) error
}

var linkConfigWriter LinkConfigWriter

// LinkConfigWriter sets the configuration updater for links.
func SetLinkConfigWriter(writer LinkConfigWriter) {
	linkConfigWriter = writer
}

var bday = time.Date(2025, time.November, 3, 0, 0, 0, 0, time.UTC)

func NewLinks() []*Link {
	return []*Link{
		{
			Stub:        "gh",
			URL:         "https://github.com",
			Description: "GitHub",
			CreatedAt:   bday,
		},
		{
			Stub:        "linkage",
			URL:         "%s:%s",
			Description: "Linkage",
			CreatedAt:   bday,
		},
	}
}

var Links []*Link

// TODO: may want to refactor this to use a map if O(n) becomes too slow
func findLink(stub string) (int, *Link) {
	for i, l := range Links {
		if l.Stub == stub {
			return i, l
		}
	}
	return -1, nil
}

func GetLinks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Links)
}

func GetLink(c *gin.Context) {
	reqLink := c.Param("stub")
	_, l := findLink(reqLink)
	if l == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Link not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, l)
}

func AddLink(c *gin.Context) {
	var newLink = Link{Stub: c.Param("stub")}
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
	Links = append(Links, &newLink)
	if err := linkConfigWriter.SaveLinks(Links); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to update config: " + err.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated, newLink)
}

func DeleteLink(c *gin.Context) {
	reqLink := c.Param("stub")
	i, l := findLink(reqLink)
	if l == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Link not found"})
		return
	}
	Links = append(Links[:i], Links[i+1:]...)
	if err := linkConfigWriter.SaveLinks(Links); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to update config: " + err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Link deleted", "link": l})
}
