package account

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/hako/durafmt"
	"github.com/stretchr/testify/assert"
)

func TestRepository_Deposit(t *testing.T) {
	x, cleanup, derr := initialize()
	assert.NoError(t, derr)
	defer cleanup()

	type fields struct {
		repository Repository
	}

	type args struct {
		ctx      context.Context
		username string
		amount   float64
	}

	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "Using optimistic locking for concurrency control",
			fields: fields{repository: x.OptimisticRepository},
			args:   args{ctx: context.Background(), username: "kurtis", amount: 1000},
		},
		{
			name:   "Using pessimistic locking for concurrency control",
			fields: fields{repository: x.PessimisticRepository},
			args:   args{ctx: context.Background(), username: "kurtis", amount: 1000},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			executions := 100
			errs := make(chan error, executions)
			start := time.Now()

			var wg sync.WaitGroup

			wg.Add(executions)

			for i := 0; i < executions; i++ {
				go func() {
					defer wg.Done()

					if derr = tt.fields.repository.Deposit(tt.args.ctx, tt.args.username, tt.args.amount); derr != nil {
						errs <- derr
					}
				}()
			}
			wg.Wait()
			close(errs)
			fmt.Printf("In %v executions, there were %v errors within %v.\n", executions, len(errs), durafmt.ParseShort(time.Since(start)))
		})
	}
}
