package linkage

import (
	"github.com/gin-gonic/gin"
)

type LinkageRoute struct {
	method   string
	endpoint string
	group    string
	handler  gin.HandlerFunc
}

// getRoutes returns a slice of LinkageRoutes representing available request routes
func getRoutes() []LinkageRoute {
	var routes = []LinkageRoute{
		LinkageRoute{
			method:   "GET",
			endpoint: "/",
			group:    "/",
			handler:  h.browserRequest,
		},
		LinkageRoute{
			method:   "GET",
			endpoint: "/links",
			group:    "/api",
			handler:  h.getLinks,
		},
		LinkageRoute{
			method:   "GET",
			endpoint: "/link/:stub",
			group:    "/api",
			handler:  h.getLink,
		},
		LinkageRoute{
			method:   "POST",
			endpoint: "/link/:stub",
			group:    "/api",
			handler:  h.addLink,
		},
		LinkageRoute{
			method:   "DELETE",
			endpoint: "/link/:stub",
			group:    "/api",
			handler:  h.deleteLink,
		},
		LinkageRoute{
			method:   "GET",
			endpoint: "/search/all",
			group:    "/api",
			handler:  h.getSearchEngines,
		},
		LinkageRoute{
			method:   "GET",
			endpoint: "/search/current",
			group:    "/api",
			handler:  h.getCurrentSearchEngine,
		},
		LinkageRoute{
			method:   "POST",
			endpoint: "/search",
			group:    "/api",
			handler:  h.addSearchEngine,
		},
	}
	return routes
}

func setupRoutes(router *gin.Engine, routes []LinkageRoute) {
	for _, route := range routes {

	}
}
