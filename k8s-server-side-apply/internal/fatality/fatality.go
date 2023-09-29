package fatality

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/pkg/errors"
)

type tracer interface {
	StackTrace() errors.StackTrace
}

func With(err error) {
	if sterr, ok := err.(tracer); ok {
		st := sterr.StackTrace()
		slog.Error(fmt.Sprintf("%s%+v\n", err.Error(), st[:len(st)-2]))
		os.Exit(1)
	}

	slog.Error(err.Error())
	os.Exit(1)
}
