package main

import (
	"fmt"

	"golang.org/x/net/context"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/angelokurtis/go-laboratory/internal/log"
	"github.com/angelokurtis/go-laboratory/internal/wire"
)

func main() {
	ctx := context.TODO()

	objects, err := getObjects()
	dieOnError(err)

	dm, err := wire.DynamicClientAndMapper()
	dieOnError(err)

	c := &Client{client: dm.Dynamic, mapper: dm.RESTMapper}

	for _, obj := range objects {
		current, err := c.Get(ctx, obj)
		dieOnError(err)

		if current == nil { // create
			name := fmt.Sprintf("%s/%s", obj.GetNamespace(), obj.GetName())
			log.Infof("Creating %q of type %q", name, obj.GroupVersionKind())
			c.Create(ctx, obj)
		} else { // update
			patch := client.MergeFrom(obj)
			diff, err := patch.Data(current)
			dieOnError(err)

			name := fmt.Sprintf("%s/%s", obj.GetNamespace(), obj.GetName())
			log.Infof("Updating %q of type %q", name, obj.GroupVersionKind())
			_ = diff
		}
	}
}

func dieOnError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
