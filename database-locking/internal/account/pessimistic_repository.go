package account

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"

	"github.com/angelokurtis/go-laboratory/database-locking/internal/persistence"
)

type PessimisticRepository struct {
	db      *sql.DB
	queries *persistence.Queries
}

func NewPessimisticRepository(db *sql.DB, queries *persistence.Queries) *PessimisticRepository {
	return &PessimisticRepository{db: db, queries: queries}
}

func (p *PessimisticRepository) Deposit(ctx context.Context, username string, amount float64) error {
	tx, err := p.db.Begin()
	if err != nil {
		return errors.WithStack(err)
	}

	defer func() { _ = tx.Rollback() }()

	qtx := p.queries.WithTx(tx)

	account, err := qtx.GetAccountAndLockForUpdates(ctx, username)
	if err != nil {
		return errors.WithStack(err)
	}

	account.Balance += amount

	if err = qtx.UpdateAccountBalance(ctx, persistence.UpdateAccountBalanceParams{
		ID:      account.ID,
		Balance: account.Balance,
	}); err != nil {
		return errors.WithStack(err)
	}

	if err = tx.Commit(); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
