package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/pkg/errors"
	cli "github.com/urfave/cli/v2"
)

func main() {
	flags := []cli.Flag{
		&cli.StringFlag{
			Name:    "user",
			Aliases: []string{"u"},
			Usage:   "the username or account identifier",
			Value:   "kurtis",
		},
		&cli.Float64Flag{
			Name:    "amount",
			Aliases: []string{"a"},
			Usage:   "specify how much money you want to deposit into a specific account",
			Value:   0.1,
		},
		&cli.IntFlag{
			Name:    "executions",
			Aliases: []string{"e"},
			Usage:   "the number of executions or operations to be performed",
			Value:   10,
		},
	}
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:   "pix",
				Action: defaultAction,
				Subcommands: []*cli.Command{
					{Name: "optimistic", Action: optimisticAction, Flags: flags},
					{Name: "pessimistic", Action: pessimisticAction, Flags: flags},
					{Name: "distributed", Action: distributedAction, Flags: flags},
				},
				Flags: flags,
			},
		},
		Flags:  flags,
		Action: defaultAction,
	}

	if err := app.Run(os.Args); err != nil {
		if sterr, ok := err.(stackTracer); ok {
			st := sterr.StackTrace()
			slog.Error(fmt.Sprintf("%s%+v\n", err.Error(), st[:len(st)-2]))
		} else {
			slog.Error(err.Error())
		}
	}
}

type stackTracer interface {
	StackTrace() errors.StackTrace
}
