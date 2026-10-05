package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"rodolrojas.com/zubi/internal/common/structs"
	"rodolrojas.com/zubi/internal/pipeline"
)

type RouteCollection struct {
	Routes []Route
	HealthChecks []Route
}

func NewRouteCollection(server *structs.BaseServerConfig) *RouteCollection {
	var collection = &RouteCollection{
		Routes: []Route{},
		HealthChecks: []Route{},
	}
	for i, service := range server.Services {
		service.Name = i;
		for j, version := range service.Versions {
			version.Version = j;
			if version.Health.Enabled {
				version.Health.Name = "health"
				route, err := BuildRoute(&service, &version, &version.Health)
				if err != nil {
					continue
				}
				collection.HealthChecks = append(collection.HealthChecks, route)
			}
			for _, endpoint := range version.Endpoints {
				route, err := BuildRoute(&service, &version, &endpoint)
				if err != nil {
					continue
				}
				collection.Routes = append(collection.Routes, route)
			}
		}
	}

	collection.MapCacheInvalidations()
	logrus.Infof("%d routes mapped and added to gateway",int(len(collection.Routes)))
	logrus.Infof("%d health checks detected",int(len(collection.HealthChecks)))
	collection.ListHealthChecks()
	collection.ListRoutes()

	return collection
}

func (rc *RouteCollection) FindRoute(key string) *Route {
	for _, route := range rc.Routes {
		if route.Key == key {
			return &route
		}
	}
	return nil
}

func (rc *RouteCollection) MapCacheInvalidations() {
	for x, route := range rc.Routes {
		if route.CacheInvalidates != nil {
			for _, invalidation := range route.CacheInvalidates {
				// TODO: Cross-version invalidation mapping
				var key string = BuildRouteKey(route.Service, route.Version, invalidation.EndpointKey)
				var targetRoute = rc.FindRoute(key)
				if targetRoute != nil {
					logrus.Debugf("Mapping cache invalidation for route %s to target %s", route.Key, targetRoute.Key)
					rc.Routes[x].CacheInvalidationTargets = append(rc.Routes[x].CacheInvalidationTargets, *targetRoute)
				} else {
					logrus.Warnf("Target route %s not found for cache invalidation in route %s", key, route.Key)
				}
			}
		}
	}
}

func (rc *RouteCollection) ListRoutes() {
	var output = "Routes mapped: \n"
	for _, route := range rc.Routes {
		var routeInfo = fmt.Sprintf("\t\t[%s]\t%s => %s \n", route.Method, route.Path, route.Upstream)
		output = fmt.Sprintf("%s%s", output, routeInfo)
	}
	logrus.Infof("%s", output)
}

func (rc *RouteCollection) ListHealthChecks() {
	var output = "Health Checks: \n"
	for _, route := range rc.HealthChecks {
		var routeInfo = fmt.Sprintf("\t\t%s => %s \n", route.Key, route.Upstream)
		output = fmt.Sprintf("%s%s", output, routeInfo)
	}
	logrus.Infof("%s", output)
}

func (rc *RouteCollection) setupRoutes(handler *mux.Router) {
	for _, route := range rc.Routes {
		handler.HandleFunc(route.Path, NewRouteHandler(&route)).Methods(route.Method)
	}
	for _, healthCheck := range rc.HealthChecks {
		handler.HandleFunc(healthCheck.Path, NewRouteHandler(&healthCheck)).Methods(healthCheck.Method)
	}
}

func NewRouteHandler(route *Route) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		pipeline.RunRequestPipeline(w, r)
	}
}