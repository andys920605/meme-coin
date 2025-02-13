package trace

import (
	"context"
	"fmt"

	texporter "github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/trace"
	"go.opentelemetry.io/contrib/detectors/gcp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.16.0"
)

var (
	domain      = "localhost/trace"
	serviceName = "localhost"
)

func NewTracerProvider(projectId, serviceName, serviceVersion, serviceDomain string) error {
	domain = fmt.Sprintf("%s/trace", serviceDomain)

	ctx := context.Background()
	// Create exporter.
	exporter, err := texporter.New(texporter.WithProjectID(projectId))
	if err != nil {
		return err
	}

	// Identify your application using resource detection
	res, err := resource.New(ctx,
		// Use the GCP resource detector to detect information about the GCP platform
		resource.WithDetectors(gcp.NewDetector()),
		// Keep the default detectors
		resource.WithTelemetrySDK(),
		// Add your own custom attributes to identify your application
		resource.WithAttributes(
			semconv.ServiceNameKey.String(serviceName),
			semconv.ServiceVersionKey.String(serviceVersion),
		),
	)
	if err != nil {
		return err
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(res),
	)
	defer tp.ForceFlush(ctx) // flushes any pending spans

	otel.SetTracerProvider(tp) // set global trace provider
	return nil
}
