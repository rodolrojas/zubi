package router

import (
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
	"rodolrojas.com/zubi/internal/common/enums"
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

func BuildRoute[T *structs.BaseEndpointConfig | *structs.HealthCheckConfig](
	service *structs.BaseServiceConfig,
	version *structs.ServiceVersionConfig,
	endpoint T,
) (Route, error) {
	var __ep structs.BaseEndpointConfig
	var __method enums.HTTPMethod

	switch v := any(endpoint).(type) {
	case *structs.BaseEndpointConfig:
		__ep = *v
	case *structs.HealthCheckConfig:
		__ep = v.BaseEndpointConfig
	}
	
	if (__ep.Method == "") {
		__method = enums.HTTPMethod("GET")
	} else {
		__method = enums.HTTPMethod(__ep.Method)
		if !__method.Valid() {
			logrus.Fatalf("%s is not an allowed method.", __method)
		}
	}
	var basePath = "/api/%s/%s%s" // /api/{service}/{version}{endpoint}
	basePath = fmt.Sprintf(basePath, service.Name, version.Version, __ep.Route)
	var fullUpstream = strings.TrimRight(version.URL, "/") + "/" + strings.TrimLeft(__ep.UpstreamPath, "/")
	
	var r = Route{
		Key: BuildRouteKey(service.Name, version.Version, __ep.Name),
		Service: service.Name,
		Version: version.Version,
		Method: string(__method),
		Path: basePath,
		Public: __ep.Public,
		Upstream: fullUpstream,
		Params: __ep.Params,
		CacheEnabled: __ep.Cache.Enabled,
		CacheTTL: int(__ep.Cache.TTL.Seconds()),
		CacheInvalidates: __ep.Cache.Invalidates,
	}
	return r, nil
}