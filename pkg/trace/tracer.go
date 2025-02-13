package trace

import (
	"context"
	"runtime"
	"strings"

	"go.opentelemetry.io/otel"
	semconv "go.opentelemetry.io/otel/semconv/v1.16.0"
	"go.opentelemetry.io/otel/trace"
)

func GetTracer() trace.Tracer {
	return otel.GetTracerProvider().Tracer(domain)
}

func StartTrace(ctx context.Context) (context.Context, trace.Span) {
	return GetTracer().Start(ctx, getFuncName())
}

func getFuncName() string {
	pc, _, _, ok := runtime.Caller(2)
	details := runtime.FuncForPC(pc)
	if ok && details != nil {
		name := details.Name()
		index := strings.Index(name, serviceName)
		if index >= 0 {
			return name[index+len(serviceName):]
		}
		return name
	}
	return ""
}

func WithHTTPMethodAttributes(method string) trace.SpanStartEventOption {
	return trace.WithAttributes(semconv.HTTPMethodKey.String(method))
}

func GetTraceIDFromContext(ctx context.Context) string {
	return trace.SpanFromContext(ctx).SpanContext().TraceID().String()
}
