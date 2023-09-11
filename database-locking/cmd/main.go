package main

import (
	"fmt"
	"log/slog"
	"os"
	"sync"
	"sync/atomic"
	"time"

	retry "github.com/avast/retry-go/v4"
	"github.com/hako/durafmt"
	"github.com/pkg/errors"
	cli "github.com/urfave/cli/v2"

	"github.com/angelokurtis/go-laboratory/database-locking/internal/account"
)

func main() {
	var (
		optimistic  bool
		pessimistic bool
		username    string
		amount      float64
	)

	app := &cli.App{
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "optimistic",
				Aliases:     []string{"o"},
				Usage:       "database locking allows simultaneous access by checking data versions for conflicts",
				Value:       false,
				Destination: &optimistic,
			},
			&cli.BoolFlag{
				Name:        "pessimistic",
				Aliases:     []string{"p"},
				Usage:       "database locking locks data to prevent simultaneous access",
				Value:       false,
				Destination: &pessimistic,
			},
			&cli.StringFlag{
				Name:        "username",
				Usage:       "the username or account identifier",
				Value:       "kurtis",
				Destination: &username,
			},
			&cli.Float64Flag{
				Name:        "amount",
				Usage:       "specify how much money you want to deposit into a specific account",
				Value:       0.1,
				Destination: &amount,
			},
		},
		Before: func(cCtx *cli.Context) error {
			if optimistic && pessimistic {
				return errors.New("cannot use both optimistic and pessimistic flags simultaneously")
			}
			return nil
		},
		Action: func(cCtx *cli.Context) error {
			var (
				repo    account.Repository
				cleanup = func() {}
				err     error
			)
			if optimistic {
				if repo, cleanup, err = initializeOptimistic(); err != nil {
					return err
				}
			} else if pessimistic {
				if repo, cleanup, err = initializePessimistic(); err != nil {
					return err
				}
			} else {
				if repo, cleanup, err = initialize(); err != nil {
					return err
				}
			}
			defer cleanup()

			executions := 100

			var retries, errs uint64

			start := time.Now()

			var wg sync.WaitGroup

			wg.Add(executions)

			for i := 0; i < executions; i++ {
				go func() {
					defer wg.Done()

					rerr := retry.Do(
						func() error {
							return repo.Deposit(cCtx.Context, username, amount)
						},
						retry.OnRetry(func(u uint, err error) {
							atomic.AddUint64(&retries, 1)
						}),
						retry.Delay(1*time.Second),
						retry.DelayType(retry.BackOffDelay),
						retry.Attempts(5),
					)
					if rerr != nil {
						atomic.AddUint64(&errs, 1)
					}
				}()
			}
			wg.Wait()
			slog.Info(fmt.Sprintf("During %v executions, we had %v errors, even after retrying %v times in just %v.", executions, errs, retries, durafmt.ParseShort(time.Since(start))))
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		if sterr, ok := err.(stackTracer); ok {
			st := sterr.StackTrace()
			slog.Error("%s%+v\n", err.Error(), st[:len(st)-2])
		} else {
			slog.Error(err.Error())
		}
	}
}

type stackTracer interface {
	StackTrace() errors.StackTrace
}
