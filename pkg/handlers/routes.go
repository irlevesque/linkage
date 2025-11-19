package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RouteInfo represents a simplified view of a route for external consumption.
type RouteInfo struct {
	Method   string
	Endpoint string
}

type RoutesReflection struct {
	Routes []RouteInfo
}

// GetRoutesHandler creates a gin.HandlerFunc that reflects available routes.
func GetRoutesHandler(routes RoutesReflection) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, routes)
	}
}

// RegisterRoutesHandler adds a route to the internal reflection of available routes.
func (r *RoutesReflection) Register(m string, e string) {
	r.Routes = append(r.Routes, RouteInfo{Method: m, Endpoint: e})
}
