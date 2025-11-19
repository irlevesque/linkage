package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/irlevesque/linkage/pkg/handlers"
)

type LinkageRoute struct {
	Method   string
	Endpoint string
	Handler  gin.HandlerFunc
}

var ApiBasePath = "/api"

// GetReflection
func GetReflection() handlers.RoutesReflection {
	var rr handlers.RoutesReflection
	for _, r := range GetApiRoutes() {
		rr.Register(r.Method, r.Endpoint)
	}
	return rr
}

// GetApiRoutes returns a slice of LinkageRoutes representing available API routes
func GetApiRoutes() []LinkageRoute {
	return []LinkageRoute{
		LinkageRoute{
			Method:   "GET",
			Endpoint: ApiBasePath + "/links",
			Handler:  handlers.GetLinks,
		},
		LinkageRoute{
			Method:   "GET",
			Endpoint: ApiBasePath + "/link/:stub",
			Handler:  handlers.GetLink,
		},
		LinkageRoute{
			Method:   "POST",
			Endpoint: ApiBasePath + "/link/:stub",
			Handler:  handlers.AddLink,
		},
		LinkageRoute{
			Method:   "DELETE",
			Endpoint: ApiBasePath + "/link/:stub",
			Handler:  handlers.DeleteLink,
		},
		LinkageRoute{
			Method:   "GET",
			Endpoint: ApiBasePath + "/search",
			Handler:  handlers.GetSearchEngines,
		},
		LinkageRoute{
			Method:   "POST",
			Endpoint: ApiBasePath + "/search",
			Handler:  handlers.AddSearchEngine,
		},
		LinkageRoute{
			Method:   "DELETE",
			Endpoint: ApiBasePath + "/search/:name",
			Handler:  handlers.DeleteSearchEngine,
		},
		LinkageRoute{
			Method:   "GET",
			Endpoint: ApiBasePath + "/search/default",
			Handler:  handlers.GetDefaultSearchEngine,
		},
		LinkageRoute{
			Method:   "POST",
			Endpoint: ApiBasePath + "/search/default/:name",
			Handler:  handlers.SetDefaultSearchEngine,
		},
	}
}

// GetBrowserRoutes returns a slice of LinkageRoutes representing available Browser routes
func GetBrowserRoutes() []LinkageRoute {
	return []LinkageRoute{
		LinkageRoute{
			Method:   "GET",
			Endpoint: "/",
			Handler:  handlers.BrowserRequest,
		},
	}
}
