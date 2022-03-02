package main

import (
	"context"

	"github.com/angelokurtis/go-laboratory/internal/log"
	"github.com/angelokurtis/go-laboratory/internal/wire"
)

func main() {
	ctx := context.Background()

	objects, err := getObjects()
	dieOnError(err)

	dm, err := wire.DynamicClientAndMapper()
	dieOnError(err)

	c := NewClient(dm.Dynamic, dm.RESTMapper)

	for i := 0; i < 5; i++ {
		for _, obj := range objects {
			err = c.Apply(ctx, obj)
			dieOnError(err)

			unstructured, err := c.Find(ctx, obj)
			dieOnError(err)

			log.Infof(unstructured.GetResourceVersion())
		}
	}
}

func dieOnError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
