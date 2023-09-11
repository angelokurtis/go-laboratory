package main

import (
	"fmt"
	"log/slog"
	"sync"
	"sync/atomic"
	"time"

	retry "github.com/avast/retry-go/v4"
	"github.com/hako/durafmt"
	cli "github.com/urfave/cli/v2"
)

func optimisticAction(cCtx *cli.Context) error {
	user := cCtx.String("user")
	amount := cCtx.Float64("amount")
	executions := cCtx.Int("executions")

	repo, cleanup, err := initializeOptimistic()
	if err != nil {
		return err
	}

	defer cleanup()

	var retries, errs uint64

	start := time.Now()

	var wg sync.WaitGroup

	wg.Add(executions)

	for i := 0; i < executions; i++ {
		go func() {
			defer wg.Done()

			rerr := retry.Do(
				func() error {
					return repo.Deposit(cCtx.Context, user, amount)
				},
				retry.OnRetry(func(u uint, err error) {
					atomic.AddUint64(&retries, 1)
				}),
				retry.Delay(1*time.Second),
				retry.DelayType(retry.RandomDelay),
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
}
