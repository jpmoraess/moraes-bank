package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Store store provides all functions to execute db queries and transactions
type Store struct {
	*Queries
	pool *pgxpool.Pool
}

func NewStore(pool *pgxpool.Pool) *Store {
	return &Store{
		pool:    pool,
		Queries: New(pool),
	}
}

// execTx executes a function within a database transaction
func (s *Store) execTx(ctx context.Context, fn func(queries *Queries) error) error {
	tx, err := s.pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(context.Background()); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit(context.Background())
}

// TransferTxParams contains the input parameters of the transfer transaction
type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

// TransferTxResult is the result of the transfer transaction
type TransferTxResult struct {
	Transfer    Transfer `json:"transfer"`
	FromAccount Account  `json:"from_account"`
	ToAccount   Account  `json:"to_account"`
	FromEntry   Entry    `json:"from_entry"`
	ToEntry     Entry    `json:"to_entry"`
}

var txKey = struct{}{}

// TransferTx performs a money transfer from one account to the other.
// It creates a new transfer record, add account entries and update accounts balance within a single database transaction
func (s *Store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult

	err := s.execTx(ctx, func(q *Queries) error {
		var err error

		result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: arg.FromAccountID,
			ToAccountID:   arg.ToAccountID,
			Amount:        arg.Amount,
		})
		if err != nil {
			return err
		}

		result.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.FromAccountID,
			Amount:    -arg.Amount,
		})
		if err != nil {
			return err
		}

		result.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.ToAccountID,
			Amount:    arg.Amount,
		})
		if err != nil {
			return err
		}

		if arg.FromAccountID < arg.ToAccountID {
			result.FromAccount, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
				ID:     arg.FromAccountID,
				Amount: -arg.Amount,
			})
			if err != nil {
				return err
			}

			result.ToAccount, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
				ID:     arg.ToAccountID,
				Amount: arg.Amount,
			})
			if err != nil {
				return err
			}
		} else {
			result.ToAccount, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
				ID:     arg.ToAccountID,
				Amount: arg.Amount,
			})
			if err != nil {
				return err
			}

			result.FromAccount, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
				ID:     arg.FromAccountID,
				Amount: -arg.Amount,
			})
			if err != nil {
				return err
			}
		}

		return nil
	})

	return result, err
}

func addMoney(
	ctx context.Context,
	q *Queries,
	accountID1 int64,
	amount1 int64,
	accountID2 int64,
	amount2 int64,
) (account1 Account, account2 Account, err error) {
	account1, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
		ID:     accountID1,
		Amount: amount1,
	})
	if err != nil {
		return
	}

	account2, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
		ID:     accountID2,
		Amount: amount2,
	})
	if err != nil {
		return
	}

	return
}
