package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// searchEngine represents a fallback search engine configuration.
type SearchEngine struct {
	Name      string `json:"name" form:"name" binding:"required"`
	QueryURL  string `json:"query_url" form:"query_url" binding:"required"`
	IsDefault bool   `json:"default" form:"default"`
}

// NewSearchEngines returns a slice of available search engines.
func NewSearchEngines() []*SearchEngine {
	return []*SearchEngine{
		{
			Name:      "Google",
			QueryURL:  "https://www.google.com/search?q=",
			IsDefault: true,
		},
		{
			Name:      "Bing",
			QueryURL:  "https://www.bing.com/search?q=",
			IsDefault: false,
		},
	}
}

var searchEngines []*SearchEngine = NewSearchEngines()

// getSearchEngines returns all search engines.
func GetSearchEngines(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, searchEngines)
}

func SearchEngineQueryURL(term string) (error, string) {
	for _, se := range searchEngines {
		if se.IsDefault {
			return nil, se.QueryURL + term
		}
	}
	// If none marked as default, return empty with 404 to signal no default present.
	return errors.New("No default search engine configured"), ""
}

// getCurrentSearchEngine returns the currently-default search engine.
func GetCurrentSearchEngine(c *gin.Context) {
	for _, se := range searchEngines {
		if se.IsDefault {
			c.IndentedJSON(http.StatusOK, se)
			return
		}
	}
	// If none marked as default, return empty with 404 to signal no default present.
	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "No default search engine configured"})
}

func setDefaultSearchEngine(se *SearchEngine, searchEngines []*SearchEngine) bool {
	for i := range searchEngines {
		searchEngines[i].IsDefault = false
	}

	se.IsDefault = true
	return true
}

// addSearchEngine adds a new search engine
func AddSearchEngine(c *gin.Context) {
	newSearchEngine := SearchEngine{Name: c.Param("name"), QueryURL: c.Param("query_url")}
	c.IndentedJSON(http.StatusAccepted, gin.H{"default": c.Param("default")})

	for _, e := range searchEngines {
		if e.Name == newSearchEngine.Name {
			c.IndentedJSON(http.StatusConflict, gin.H{"error": "Search engine already exists", "searchEngine": e})
			return
		}
	}

	isDefault, err := strconv.ParseBool(c.Param("default"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error parsing default": err.Error()})
		return
	}

	if err := c.ShouldBind(&newSearchEngine); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	searchEngines = append(searchEngines, &newSearchEngine)

	if isDefault {
		setDefaultSearchEngine(&newSearchEngine, searchEngines)
	}

	c.IndentedJSON(http.StatusCreated, newSearchEngine)
}
