package account

import "context"

type Repository interface {
	Deposit(ctx context.Context, username string, amount float64) error
}
