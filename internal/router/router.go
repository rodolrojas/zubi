package router

import (
	"github.com/sirupsen/logrus"
	"rodolrojas.com/zubi/internal/common/structs"
)

type Router struct {
	routes RouteCollection
	handler any
}

func NewRouter(services *structs.BaseServerConfig) (*Router, error) {
	var r = &Router{
		routes: *NewRouteCollection(services),
		handler: nil,
	}

	logrus.Infof("%d routes mapped and added to gateway",int(len(r.routes.Routes)))
	logrus.Infof("%d health checks detected",int(len(r.routes.HealthChecks)))
	r.routes.ListHealthChecks()
	r.routes.ListRoutes()
	return r, nil
}

func (r *Router) GetHandler() *any {
	return &r.handler
}