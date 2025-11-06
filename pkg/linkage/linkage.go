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
	app.Host = address
	app.Configure()

	return app
}

func (l *Linkage) Configure() {
	l.App = gin.Default()
	l.App.Use(gin.Logger())
	l.App.Use(gin.Recovery())

	routes := GetRoutes()
	for _, route := range routes {
		l.App.Handle(route.method, route.endpoint, route.handler)
	}
}

func (l *Linkage) Serve() {
	l.App.Run(l.Host)
}
