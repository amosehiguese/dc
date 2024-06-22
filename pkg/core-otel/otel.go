package coreotel

import (
	"context"
	"errors"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.uber.org/zap"
)

func SetUpOtelSDK(ctx context.Context, log *zap.Logger, svc string) (func(context.Context) error, error) {
	var shutdownFuncs []func(context.Context) error

	shutdown := func(context.Context) error {
		var err error
		for _, fn := range shutdownFuncs {
			err = errors.Join(err, fn(ctx))
		}
		shutdownFuncs = nil
		return err
	}

	var err error
	handlerErr := func(inErr error) {
		err = errors.Join(inErr, shutdown(ctx))
	}

	// Set up propagator
	prop := NewPropagator()
	otel.SetTextMapPropagator(prop)

	// Set up trace provider
	traceProvider, err := InitTracing(ctx, log, svc)
	if err != nil {
		handlerErr(err)
		return shutdown, err
	}

	shutdownFuncs = append(shutdownFuncs, traceProvider.Shutdown)
	otel.SetTracerProvider(traceProvider)

	// Set up meter provider
	meterProvider, err := InitMetrics(ctx, log)
	if err != nil {
		handlerErr(err)
		return shutdown, err
	}

	shutdownFuncs = append(shutdownFuncs, meterProvider.Shutdown)
	otel.SetMeterProvider(meterProvider)

	return shutdown, nil
}

func NewPropagator() propagation.TextMapPropagator {
	return propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	)
}
