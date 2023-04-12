package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

func newHttpExporter(ctx context.Context) (*otlptrace.Exporter, error) {

	exp, err := otlptrace.New(
		ctx,
		otlptracehttp.NewClient(),
	)

	return exp, err
}

func newStdOutExporter() (trace.SpanExporter, error) {
	return stdouttrace.New(
		stdouttrace.WithPrettyPrint(),
		stdouttrace.WithoutTimestamps(),
	)
}

func newFileExporter(w io.Writer) (trace.SpanExporter, error) {
	return stdouttrace.New(
		stdouttrace.WithPrettyPrint(),
		stdouttrace.WithoutTimestamps(),
		stdouttrace.WithWriter(w),
	)
}

// newResource returns a resource describing this application.
func newResource() *resource.Resource {
	r, _ := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("fib"),
			semconv.ServiceVersionKey.String("v0.1.0"),
			attribute.String("environment", "demo"),
		),
	)
	return r
}

func main() {

	l := log.New(os.Stdout, "", 0)

	// Write telemetry data to a file.
	f, err := os.Create("traces.txt")
	if err != nil {
		l.Fatal(err)
	}

	defer f.Close()

	fileExp, _ := newFileExporter(f)
	stdOutExp, err := newStdOutExporter()
	httpExp, err := newHttpExporter(context.Background())
	if err != nil {
		l.Fatal(err)
	}

	tp := trace.NewTracerProvider(
		trace.WithBatcher(httpExp),
		trace.WithBatcher(stdOutExp),
		trace.WithBatcher(fileExp),
		trace.WithResource(newResource()),
	)
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			l.Fatal(err)
		}
	}()
	otel.SetTracerProvider(tp)

	helloHandler := func(w http.ResponseWriter, req *http.Request) {

		_, _ = io.WriteString(w, "Hello, world!\n")
	}

	// Wrapping http handlerwith otel auto instrumentation
	autoInstrumentedHandler := otelhttp.NewHandler(http.HandlerFunc(helloHandler), "Hello")

	logrus.Infoln("Hello Open Telemetry!")
	http.Handle("/hello", autoInstrumentedHandler)
	errHTTP := http.ListenAndServe(":7777", nil)
	if errHTTP != nil {
		panic(err)
	}
}
