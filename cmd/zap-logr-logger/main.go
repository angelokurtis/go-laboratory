package main

import defaultLog "log"

func main() {
	log, err := NewLogger()
	dieOnError(err)

	// DebugLevel is -1
	log.V(-1).Info("ha!", "level", -1)
	// InfoLevel is 0
	log.V(0).Info("ha!", "level", 0)
	// WarnLevel is 1
	log.V(1).Info("ha!", "level", 1)
	log.Error(nil, "ha!")
	// for i := -1; i <= 3; i++ {
	// 	log.V(i).Info("ha!", "level", i)
	// }
}

func dieOnError(err error) {
	if err != nil {
		defaultLog.Fatal(err)
	}
}
