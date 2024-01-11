package repository

import (
	"context"
	"time"

	"github.com/dhimweray222/test-be-rekadigital/config"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Store interface {
	WithTransaction(ctx context.Context, fn func(pgx.Tx) error) error
	WithoutTransaction(ctx context.Context, fn func(*pgxpool.Pool) error) error
}

type StoreImpl struct {
	Db *pgxpool.Pool
}

func NewStore(db *pgxpool.Pool) Store {
	return &StoreImpl{
		Db: db,
	}
}

func (r *StoreImpl) WithTransaction(ctx context.Context, fn func(pgx.Tx) error) error {

	c, cancel := context.WithTimeout(context.Background(), time.Duration(config.TimeOutDuration)*time.Second)
	defer cancel()

	tx, err := r.Db.Begin(c)
	if err != nil {
		return err
	}

	if err := fn(tx); err != nil {
		_ = tx.Rollback(ctx)
		return err
	}

	return tx.Commit(ctx)
}

func (r *StoreImpl) WithoutTransaction(ctx context.Context, fn func(*pgxpool.Pool) error) error {

	if err := fn(r.Db); err != nil {
		return err
	}

	return nil
}
