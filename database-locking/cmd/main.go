package main

import (
	"log"

	"github.com/pkg/errors"
	_ "go.uber.org/automaxprocs"
)

func main() {
	type stackTracer interface {
		StackTrace() errors.StackTrace
	}

	if err := run(); err != nil {
		if sterr, ok := err.(stackTracer); ok {
			st := sterr.StackTrace()
			log.Fatalf("%s%+v\n", err.Error(), st[:len(st)-2])
		} else {
			log.Fatalln(err)
		}
	}
}

func run() error {
	return nil
}
