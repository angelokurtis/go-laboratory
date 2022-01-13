package main

import (
	defaultLog "log"

	"github.com/angelokurtis/go-laboratory/internal/wire"
)

func main() {
	log, err := wire.ZapLogger()
	dieOnError(err)

	log.Info("ha!")
	log.Warn("ha!")
	log.Debug("ha!")
}

func dieOnError(err error) {
	if err != nil {
		defaultLog.Fatal(err)
	}
}
