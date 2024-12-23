package telemetry

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

func NewExporter(ctx context.Context, serviceName, serviceVersion string) (func(context.Context) error, error) {
	resource, err := resource.New(ctx, resource.WithAttributes(
		semconv.ServiceNameKey.String(serviceName),
		semconv.ServiceVersionKey.String(serviceVersion),
	))
	if err != nil {
		return nil, err
	}

	exporter, err := otlptracegrpc.New(ctx, otlptracegrpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	bsp := trace.NewBatchSpanProcessor(exporter)

	traceProvider := trace.NewTracerProvider(trace.WithSampler(trace.AlwaysSample()), trace.WithSpanProcessor(bsp), trace.WithResource(resource))

	otel.SetTracerProvider(traceProvider)

	return exporter.Shutdown, nil
}
