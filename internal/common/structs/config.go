package structs

import "time"

type BaseEndpointParamsConfig struct {
	Name string `yaml:"name"`
	Type string `yaml:"type"`
	Required bool `yaml:"required"`
}
type BaseCacheConfig struct {
	Enabled bool `yaml:"enabled"`
	TTL time.Duration `yaml:"ttl"`
	Invalidates []BaseCacheInvalidationConfig `yaml:"invalidates"`
}
type BaseCacheInvalidationConfig struct {
	EndpointKey string `yaml:"endpoint"`
	Params map[string]string `yaml:"params"`
	Wide bool `yaml:"wide"`
}
type BaseEndpointConfig struct {
	Name string `yaml:"name"`
	Route string `yaml:"route"`
	Method string `yaml:"method"`
	Public bool `yaml:"public"`
	Params []BaseEndpointParamsConfig `yaml:"params"`
	UpstreamPath string `yaml:"upstream_path"`
	Response BaseEndpointResponseConfig `yaml:"response"`
	Request BaseEndpointRequestConfig `yaml:"request"`
	Cache BaseCacheConfig `yaml:"cache"`
}
type BaseEndpointResponseConfig struct {
	ContentType string `yaml:"content_type"`
	ExpectedSuccessStatus int `yaml:"expected_success_status"`
}
type BaseEndpointRequestConfig struct {
	AllowedContentTypes []string `yaml:"allowed_content_types"`
}
type HealthCheckConfig struct {
	BaseEndpointConfig
	Enabled bool `yaml:"enabled"`
}	
type ServiceVersionConfig struct {
	Version string `yaml:"version"`
	URL     string `yaml:"url"`
	Public  bool `yaml:"public"`
	Health  HealthCheckConfig `yaml:"health"`
	Endpoints []BaseEndpointConfig `yaml:"endpoints"`
}
type BaseServiceConfig struct {
	Name string `yaml:"name"`
	BasePath string `yaml:"base_path"`
	Versions map[string]ServiceVersionConfig `yaml:"versions"`
}
type BaseServerConfig struct{
	Port int `yaml:"port"`
	ReadTimeout time.Duration `yaml:"read_timeout"`
	WriteTimeout time.Duration `yaml:"write_timeout"`
	IdleTimeout time.Duration `yaml:"idle_timeout"`
	Services map[string]BaseServiceConfig `yaml:"services"`
	RateLimit RateLimitConfig `yaml:"rate_limit"`
}
type RateLimitConfig struct {
	Enabled bool `yaml:"enabled"`
	MaxTokens int `yaml:"max_tokens"`
	TokenRefillRate int `yaml:"token_refill_rate"`
	TokenRefillIntervalSeconds int `yaml:"token_refill_interval_seconds"`
}
type BaseConfig struct {
	Server BaseServerConfig `yaml:"server"`
}