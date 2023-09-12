package account

import (
	"context"
	"fmt"
	"log/slog"

	redsync "github.com/go-redsync/redsync/v4"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"

	"github.com/angelokurtis/go-laboratory/database-locking/internal/persistence"
)

type DistributedRepository struct {
	queries *persistence.Queries
	rs      *redsync.Redsync
}

func NewDistributedRepository(queries *persistence.Queries, rs *redsync.Redsync) *DistributedRepository {
	return &DistributedRepository{queries: queries, rs: rs}
}

func (o *DistributedRepository) Deposit(ctx context.Context, username string, amount float64) error {
	key := fmt.Sprintf("account:username:%s", username)

	mutex := o.rs.NewMutex(key)
	if err := mutex.Lock(); err != nil {
		return errors.WithStack(err)
	}

	defer func() {
		if ok, merr := mutex.Unlock(); !ok || merr != nil {
			slog.Warn("an error happened while unlocking",
				slog.String("component", fmt.Sprintf("%T", mutex)),
				slog.String("err", merr.Error()),
			)
		}
	}()

	account, err := o.queries.GetAccount(ctx, username)
	if err != nil {
		return errors.WithStack(err)
	}

	current, err := decimal.NewFromString(account.Balance)
	if err != nil {
		return errors.WithStack(err)
	}

	updated := current.Add(decimal.NewFromFloat(amount))

	if err = o.queries.UpdateAccountBalance(ctx, persistence.UpdateAccountBalanceParams{
		ID:      account.ID,
		Balance: updated.String(),
	}); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
