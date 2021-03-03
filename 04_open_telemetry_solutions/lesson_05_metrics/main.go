package main

import (
	"context"
	"fib/fibonacci"
	"fib/kafka"
	"fmt"
	"net/http"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

func main() {
	if err := initGlobalTracer(nil); err != nil {
		panic(err)
	}
	initMetricsProvider()
	http.HandleFunc("/fib", FibServer)
	http.HandleFunc("/favicon.ico", faviconHandler)
	http.ListenAndServe(":28080", nil)
}

// FibServer handles HTTP requests for fibonacci calculation
func FibServer(w http.ResponseWriter, r *http.Request) {
	tracer := otel.Tracer("http")
	ctx := context.Background()

	var span trace.Span
	ctx, span = tracer.Start(ctx, "http-request")
	defer span.End()

	if n, err := getIntParam(r); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {
		result, numIterations := fibonacci.Wrap(ctx, fibonacci.New()).Calc(n)
		reportMetric(n, numIterations)
		kafka.Send(result)
		w.Write([]byte(fmt.Sprintf("%d", result)))
	}
}

/*
DO NOT REMOVE ANY TEXT BELOW THIS LINE






















*/
func hide(v interface{}) {
	kafka.PreserveImport()
	fibonacci.PreserveImport()
	context.Background()
	var tracer trace.Tracer
	tracer = otel.Tracer("")
	hide(tracer)
	hide(fmt.Sprintf("%d", 3))
}
