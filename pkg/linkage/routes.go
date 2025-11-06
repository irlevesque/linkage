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
			endpoint: "/links",
			handler:  handlers.GetLinks,
		},
		LinkageRoute{
			method:   "GET",
			endpoint: "/link/:stub",
			handler:  handlers.GetLink,
		},
		LinkageRoute{
			method:   "POST",
			endpoint: "/link/:stub",
			handler:  handlers.AddLink,
		},
		LinkageRoute{
			method:   "DELETE",
			endpoint: "/link/:stub",
			handler:  handlers.DeleteLink,
		},
		LinkageRoute{
			method:   "GET",
			endpoint: "/search/all",
			handler:  handlers.GetSearchEngines,
		},
		LinkageRoute{
			method:   "GET",
			endpoint: "/search/current",
			handler:  handlers.GetCurrentSearchEngine,
		},
		LinkageRoute{
			method:   "POST",
			endpoint: "/search",
			handler:  handlers.AddSearchEngine,
		},
	}
	return routes
}
