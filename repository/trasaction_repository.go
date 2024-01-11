package repository

import (
	"context"

	"github.com/dhimweray222/test-be-rekadigital/model/domain"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type transactionRepository struct {
	db               Store
	TransactionQuery TransactionQuery
}

type TransactionRepository interface {
	CreateTrasaction(c context.Context, transaction domain.Transaction) error
	FindByID(c context.Context, id string) (domain.Transaction, error)
	FindAll(c context.Context, menu, customerId string) ([]domain.Transaction, int, error)
}

func NewTransactionRepository(db Store, q TransactionQuery) TransactionRepository {
	return &transactionRepository{
		db:               db,
		TransactionQuery: q,
	}
}

func (r *transactionRepository) CreateTrasaction(c context.Context, transaction domain.Transaction) error {
	var err error

	err = r.db.WithTransaction(c, func(tx pgx.Tx) error {
		if err = r.TransactionQuery.CreateTrasaction(c, tx, transaction); err != nil {

			return err
		}
		return nil
	})

	return err
}

func (r *transactionRepository) FindByID(c context.Context, id string) (domain.Transaction, error) {
	var transaction domain.Transaction
	var err error
	err = r.db.WithoutTransaction(c, func(db *pgxpool.Pool) error {
		if transaction, err = r.TransactionQuery.FindByID(c, db, id); err != nil {
			return err
		}
		return nil
	})

	return transaction, err
}

func (r *transactionRepository) FindAll(c context.Context, menu, customerId string) ([]domain.Transaction, int, error) {
	var transactions []domain.Transaction
	var err error
	var total int
	err = r.db.WithoutTransaction(c, func(db *pgxpool.Pool) error {
		if transactions, total, err = r.TransactionQuery.FindAll(c, db, menu, customerId); err != nil {
			return err
		}
		return nil
	})

	return transactions, total, err
}
