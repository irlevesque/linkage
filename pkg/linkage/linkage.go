package linkage

import (
	"github.com/gin-gonic/gin"
)

type Linkage struct {
	App  *gin.Engine
	Host string
}

func New(address string) *Linkage {
	app := new(Linkage)
	if address == "" || address == ":" {
		address = "localhost:8080"
	}
	app.Host = address
	app.App = gin.Default()
	app.addRoutes()

	return app
}

func (l *Linkage) Serve() {
	l.App.Run(l.Host)
}

func (l *Linkage) addRoutes() {
	routes := GetRoutes()
	for _, route := range routes {
		l.App.Handle(route.method, route.endpoint, route.handler)
	}
}
