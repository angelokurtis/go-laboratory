package main

import (
	"context"
	"log"
	"sync"

	"github.com/angelokurtis/go-tracing"
)

var properties = struct {
	sync.RWMutex
	data map[int]int
}{data: make(map[int]int)}

func main() {
	// tracing
	closer, err := tracing.Initialize(
		tracing.WithServiceName("mutual-exclusion"),
		tracing.WithSamplerType(tracing.ConstantSampler),
		tracing.WithSamplerParam(1.0),
		tracing.WithEndpoint("http://jaeger-collector.lvh.me/api/traces"),
	)

	dieIfNotNil(err)
	defer closer.Close()
	span, ctx := tracing.StartSpanFromContext(context.TODO())
	defer span.Finish()

	log.Printf("%s starting", span)
	// concurrency test
	go writeLoop(ctx)
	go readLoop(ctx)

	// stop program from exiting, must be killed
	block := make(chan struct{})
	<-block
}

func writeLoop(ctx context.Context) {
	span, ctx := tracing.StartSpanFromContext(ctx)
	defer span.Finish()
	log.Printf("%s writing", span)
	for {
		for i := 0; i < 100; i++ {
			properties.Lock()
			properties.data[i] = i
			log.Printf("%s write %d:%d", span, i, i)
			properties.Unlock()
		}
	}
}

func readLoop(ctx context.Context) {
	span, ctx := tracing.StartSpanFromContext(ctx)
	defer span.Finish()
	log.Printf("%s reading", span)
	for {
		properties.RLock()
		for k, v := range properties.data {
			log.Printf("%s read %d:%d", span, k, v)
		}
		properties.RUnlock()
	}
}

func dieIfNotNil(err error) {
	if err != nil {
		panic(err)
	}
}
