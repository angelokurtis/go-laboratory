package account

import (
	"context"

	"github.com/pkg/errors"

	"github.com/angelokurtis/go-laboratory/database-locking/internal/persistence"
)

type OptimisticRepository struct {
	queries *persistence.Queries
}

func NewOptimisticRepository(queries *persistence.Queries) *OptimisticRepository {
	return &OptimisticRepository{queries: queries}
}

func (o *OptimisticRepository) Deposit(ctx context.Context, username string, amount float64) error {
	account, err := o.queries.GetAccount(ctx, username)
	if err != nil {
		return errors.WithStack(err)
	}

	account.Balance += amount

	res, err := o.queries.UpdateAccountBalanceVersion(ctx, persistence.UpdateAccountBalanceVersionParams{
		ID:      account.ID,
		Balance: account.Balance,
		Version: account.Version,
	})
	if err != nil {
		return errors.WithStack(err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return errors.WithStack(err)
	}

	if rowsAffected == 0 {
		return errors.WithStack(new(ZeroRowsAffectedError))
	}

	return nil
}

type ZeroRowsAffectedError struct{}

func (z *ZeroRowsAffectedError) Error() string {
	return "no changes made: zero rows were affected"
}
