package handlers

import (
	"errors"
	"net/http"
	"strings"

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
			Name:      "google",
			QueryURL:  "https://www.google.com/search?q=",
			IsDefault: true,
		},
		{
			Name:      "bing",
			QueryURL:  "https://www.bing.com/search?q=",
			IsDefault: false,
		},
	}
}

var SearchEngines []*SearchEngine

// SearchEngineConfigWriter defines the interface for persisting SEs to disk
type SearchEngineConfigWriter interface {
	SaveSearchEngines(se []*SearchEngine) error
}

var searchEngineConfigWriter SearchEngineConfigWriter

// SetSearchEngineConfigWriter sets the writer for SEs
func SetSearchEngineConfigWriter(writer SearchEngineConfigWriter) {
	searchEngineConfigWriter = writer
}

// GetSearchEngines returns all search engines.
func GetSearchEngines(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, SearchEngines)
}

// SearchEngineQueryURL returns a URL for redirecting to the default search engine
func SearchEngineQueryURL(term string) (error, string) {
	for _, se := range SearchEngines {
		if se.IsDefault {
			return nil, se.QueryURL + term
		}
	}
	// If none marked as default, return empty with 404 to signal no default present.
	return errors.New("No default search engine configured"), ""
}

// getDefaultSearchEngine returns the current default search engine.
func GetDefaultSearchEngine(c *gin.Context) {
	for _, se := range SearchEngines {
		if se.IsDefault {
			c.IndentedJSON(http.StatusOK, se)
			return
		}
	}
	// If none marked as default, return empty with 404 to signal no default present.
	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "No default search engine configured"})
}

func setDefaultSearchEngine(se *SearchEngine, searchEngines []*SearchEngine) error {
	for i := range searchEngines {
		searchEngines[i].IsDefault = false
	}

	se.IsDefault = true

	if err := searchEngineConfigWriter.SaveSearchEngines(SearchEngines); err != nil {
		return err
	}

	return nil
}

func SetDefaultSearchEngine(c *gin.Context) {
	d := strings.ToLower(c.Param("name"))
	for _, se := range SearchEngines {
		if se.Name == d {
			if err := setDefaultSearchEngine(se, SearchEngines); err != nil {
				c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to update config: " + err.Error()})
			}
			c.IndentedJSON(http.StatusAccepted, se)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Search engine not found"})
}

// addSearchEngine adds a new search engine
func AddSearchEngine(c *gin.Context) {
	var newSearchEngine SearchEngine
	if err := c.ShouldBind(&newSearchEngine); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error parsing input": err.Error()})
		return
	}

	for _, e := range SearchEngines {
		if e.Name == newSearchEngine.Name {
			c.IndentedJSON(http.StatusConflict, gin.H{"error": "Search engine already exists", "searchEngine": e})
			return
		}
	}

	SearchEngines = append(SearchEngines, &newSearchEngine)

	if newSearchEngine.IsDefault {
		setDefaultSearchEngine(&newSearchEngine, SearchEngines)
	}

	if err := searchEngineConfigWriter.SaveSearchEngines(SearchEngines); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error saving to disk: ": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, newSearchEngine)
}

func DeleteSearchEngine(c *gin.Context) {
	name := strings.ToLower(c.Param("name"))
	for i, se := range SearchEngines {
		if se.Name == name {
			SearchEngines = append(SearchEngines[:i], SearchEngines[i+1:]...)
			if err := searchEngineConfigWriter.SaveSearchEngines(SearchEngines); err != nil {
				c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.IndentedJSON(http.StatusOK, gin.H{"message": "Search engine deleted", "searchEngine": se})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Search engine not found"})
}
