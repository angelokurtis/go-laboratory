package main

import (
	defaultLog "log"

	"go.uber.org/zap"

	"github.com/angelokurtis/go-laboratory/internal/kustomize"
	"github.com/angelokurtis/go-laboratory/internal/wire"
)

var log *zap.SugaredLogger

func init() {
	logger, err := wire.ZapLogger()
	if err != nil {
		defaultLog.Fatal(err)
	}
	log = logger.Sugar()
}

func main() {
	generator := &kustomize.ConfigMapGenerator{
		Name:     "a-configmap",
		Files:    []string{".gitignore"},
		Literals: []string{"altGreeting=Good Morning!"},
	}

	k, err := wire.KustomizeClient()
	dieOnError(err)

	err = k.GenerateConfigMap("/home/kurtis/wrkspc/github.com/angelokurtis/go-laboratory", generator)
	dieOnError(err)
}

func dieOnError(err error) {
	if err != nil {
		log.Fatalf("%+v", err)
	}
}
