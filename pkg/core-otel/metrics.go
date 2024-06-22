package coreotel

import (
	"context"

	"go.opentelemetry.io/otel/exporters/prometheus"
	sdkmetrics "go.opentelemetry.io/otel/sdk/metric"
	"go.uber.org/zap"
)

func InitMetrics(ctx context.Context, log *zap.Logger) (*sdkmetrics.MeterProvider, error) {
	exporter, err := prometheus.New()
	if err != nil {
		log.Sugar().Warnf("warn: Failed to create metrics exporter: %v", err)
	}

	meterProvider := sdkmetrics.NewMeterProvider(
		sdkmetrics.WithReader(exporter),
	)

	return meterProvider, nil
}
