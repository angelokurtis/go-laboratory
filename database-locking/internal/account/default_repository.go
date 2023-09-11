package account

import (
	"context"

	"github.com/pkg/errors"
	"github.com/shopspring/decimal"

	"github.com/angelokurtis/go-laboratory/database-locking/internal/persistence"
)

type DefaultRepository struct {
	queries *persistence.Queries
}

func NewDefaultRepository(queries *persistence.Queries) *DefaultRepository {
	return &DefaultRepository{queries: queries}
}

func (o *DefaultRepository) Deposit(ctx context.Context, username string, amount float64) error {
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
