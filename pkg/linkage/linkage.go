package linkage

import (
	"github.com/gin-gonic/gin"
	"github.com/irlevesque/linkage/pkg/handlers"
	"github.com/irlevesque/linkage/pkg/routes"
)

type Linkage struct {
	App  *gin.Engine
	Host string
}

// New creates a new Linkage instance with the specified address and port.
// Defaults to localhost:8080
func New(address string, port string) *Linkage {
	app := new(Linkage)
	if address == "" {
		address = "localhost"
	}
	if port == "" {
		port = "8080"
	}
	app.Host = address + ":" + port
	app.App = gin.Default()
	app.addRoutes()

	return app
}

// Serve starts the Linkage service
func (l *Linkage) Serve() {
	l.App.Run(l.Host)
}

func (l *Linkage) addRoutes() {
	for _, r := range routes.GetBrowserRoutes() {
		l.App.Handle(r.Method, r.Endpoint, r.Handler)
	}

	for _, r := range routes.GetApiRoutes() {
		l.App.Handle(r.Method, r.Endpoint, r.Handler)
	}

	reflect := routes.GetReflection()
	l.App.Handle("GET", routes.ApiBasePath, handlers.GetRoutesHandler(reflect))
}
