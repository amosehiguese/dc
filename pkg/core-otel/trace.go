package coreotel

import (
	"context"
	"time"

	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.25.0"
	"go.uber.org/zap"
)

func InitTracing(ctx context.Context, log *zap.Logger, svc string) (*sdktrace.TracerProvider, error) {
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	exporter, err := otlptrace.New(context.Background(), otlptracehttp.NewClient(
		otlptracehttp.WithEndpoint(""), //input endpoint
		otlptracehttp.WithHeaders(headers),
		otlptracehttp.WithInsecure(),
	))
	if err != nil {
		log.Sugar().Warnf("warn: Failed to create trace exporter: %v", err)
	}

	tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(
			exporter,
			sdktrace.WithMaxExportBatchSize(sdktrace.DefaultMaxExportBatchSize),
			sdktrace.WithBatchTimeout(sdktrace.DefaultScheduleDelay*time.Millisecond),
		),
		sdktrace.WithResource(
			resource.NewWithAttributes(
				semconv.SchemaURL,
				semconv.ServiceNameKey.String(svc),
			),
		),
	)

	return tracerProvider, nil
}
