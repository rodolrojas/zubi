package router

import (
	"fmt"

	"rodolrojas.com/zubi/internal/common/structs"
)

type Route struct {
	Key string
	Service string
	Version string
	Method string
	Path   string
	Public bool
	Upstream string
	Params []structs.BaseEndpointParamsConfig
	CacheEnabled bool
	CacheTTL int
	CacheInvalidates []structs.BaseCacheInvalidationConfig
	CacheInvalidationTargets []Route
}

func BuildRouteKey(service, version, endpointName string) string {
	return fmt.Sprintf("%s:%s:%s", service, version, endpointName)
}

func BuildRoute(service, version string , endpoint structs.BaseEndpointConfig) (Route, error) {
	var basePath = "/api/%s/%s%s" // /api/{service}/{version}{endpoint}
	basePath = fmt.Sprintf(basePath, service, version, endpoint.Route)
	var r = Route{
		Key: BuildRouteKey(service, version, endpoint.Name),
		Service: service,
		Version: version,
		Method: endpoint.Method,
		Path: basePath,
		Public: endpoint.Public,
		Upstream: endpoint.UpstreamPath,
		Params: endpoint.Params,
		CacheEnabled: endpoint.Cache.Enabled,
		CacheTTL: int(endpoint.Cache.TTL.Seconds()),
		CacheInvalidates: endpoint.Cache.Invalidates,
	}
	return r, nil
}