package main

import (
	"os"

	"github.com/enescakir/emoji"
	"github.com/gookit/color"
	"github.com/pkg/errors"
)

type stackTracer interface {
	Cause() error
	StackTrace() errors.StackTrace
}

func dieIfErr(err error) {
	if err != nil {
		if serr, ok := err.(stackTracer); ok {
			color.Redf("%s %+v", emoji.CrossMark, serr.Cause())
			stack := serr.StackTrace()
			color.Redf("%+v", stack[:len(stack)-2])
			os.Exit(1)
		}

		color.Redf("%s %+v", emoji.CrossMark, err)
		os.Exit(1)
	}
}
