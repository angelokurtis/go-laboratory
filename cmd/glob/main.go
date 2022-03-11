package main

import (
	"github.com/gobwas/glob"

	"github.com/angelokurtis/go-laboratory/internal/log"
)

const rule = "**/*/github.com/tiagoangelozup/charlescd-operator/internal/tracing/main.go"

//const str = "/home/tiagoangelo/wrkspc/github.com/tiagoangelozup/charlescd-operator/main.go"
const str = "/home/tiagoangelo/wrkspc/github.com/tiagoangelozup/charlescd-operator/internal/tracing/main.go"

func main() {
	var g glob.Glob

	g, err := glob.Compile(rule, '/')
	if err != nil {
		log.Fatal(err)
	}
	if g.Match(str) {
		log.Info("found")
	} else {
		log.Info("not found")
	}
}
