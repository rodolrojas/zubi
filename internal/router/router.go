package router

import (
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
	return r, nil
}

func (r *Router) GetHandler() *any {
	return &r.handler
}