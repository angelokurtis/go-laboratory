package main

import defaultLog "log"

func main() {
	log, err := NewLogger()
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
