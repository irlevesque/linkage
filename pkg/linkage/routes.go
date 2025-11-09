package linkage

import (
	"github.com/gin-gonic/gin"
	"github.com/irlevesque/linkage/pkg/handlers"
)

type LinkageRoute struct {
	method   string
	endpoint string
	handler  gin.HandlerFunc
}

var apiBasePath = "/api"

// GetRoutes returns a slice of LinkageRoutes representing available request routes
func GetRoutes() []LinkageRoute {
	var routes = []LinkageRoute{
		LinkageRoute{
			method:   "GET",
			endpoint: "/",
			handler:  handlers.BrowserRequest,
		},
		LinkageRoute{
			method:   "GET",
			endpoint: apiBasePath + "/links",
			handler:  handlers.GetLinks,
		},
		LinkageRoute{
			method:   "GET",
			endpoint: apiBasePath + "/link/:stub",
			handler:  handlers.GetLink,
		},
		LinkageRoute{
			method:   "POST",
			endpoint: apiBasePath + "/link/:stub",
			handler:  handlers.AddLink,
		},
		LinkageRoute{
			method:   "DELETE",
			endpoint: apiBasePath + "/link/:stub",
			handler:  handlers.DeleteLink,
		},
		LinkageRoute{
			method:   "GET",
			endpoint: apiBasePath + "/search/all",
			handler:  handlers.GetSearchEngines,
		},
		LinkageRoute{
			method:   "GET",
			endpoint: apiBasePath + "/search/current",
			handler:  handlers.GetCurrentSearchEngine,
		},
		LinkageRoute{
			method:   "POST",
			endpoint: apiBasePath + "/search",
			handler:  handlers.AddSearchEngine,
		},
	}
	return routes
}
