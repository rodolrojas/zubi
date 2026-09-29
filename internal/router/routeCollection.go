package router

import (
	"github.com/sirupsen/logrus"
	"rodolrojas.com/zubi/internal/common"
	"rodolrojas.com/zubi/internal/common/structs"
)

type RouteCollection struct {
	Routes []Route
}

func NewRouteCollection(server *structs.BaseServerConfig) *RouteCollection {
	var collection = &RouteCollection{
		Routes: []Route{},
	}
	var serviceName, versionName string;
	for i, service := range server.Services {
		serviceName = i
		for j, version := range service.Versions {
			versionName = j
			for _, endpoint := range version.Endpoints {
				route, err := BuildRoute(serviceName, versionName, endpoint)
				if err != nil {
					continue
				}
				collection.AddRoute(route)
			}
		}
	}
	collection.MapCacheInvalidations()
	common.DebugStruct(collection)
	return collection
}

func (rc *RouteCollection) AddRoute(route Route) {
	rc.Routes = append(rc.Routes, route)
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