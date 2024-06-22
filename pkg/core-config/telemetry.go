package coreconfig

import coreutils "github.com/amosehiguese/dc/pkg/core-utils"

// telemetry config represents the configuration for
// enabling telemetry
type telemetryConfig struct {
	Otel_enabled  string
	Otel_endpoint string
}

func setTelemetryConfig() *telemetryConfig {
	var t telemetryConfig
	coreutils.MustMapEnv(&t.Otel_enabled, "OTEL_ENABLED")
	coreutils.MustMapEnv(&t.Otel_endpoint, "OTEL_ENDPOINT")

	return &t
}
