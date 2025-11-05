package linkage

import (
	"github.com/gin-gonic/gin"
)

type Linkage struct {
	App  *gin.Engine
	Host string
}

func (l *Linkage) Configure() {
	l.App = gin.Default()
	l.App.Use(gin.Logger())
	l.App.Use(gin.Recovery())

	routes := getRoutes()
	println(routes)
}

func New(address string) *Linkage {
	app := new(Linkage)
	app.Host = address

	return &Linkage{
		App:  app.App,
		Host: address,
	}
}
