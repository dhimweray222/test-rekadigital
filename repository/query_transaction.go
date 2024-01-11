package repository

import (
	"context"
	"fmt"
	"sync"

	"github.com/dhimweray222/test-be-rekadigital/model/domain"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TransactionQuery interface {
	CreateTrasaction(c context.Context, tx pgx.Tx, ticket domain.Transaction) error
	FindByID(c context.Context, db *pgxpool.Pool, id string) (domain.Transaction, error)
	FindAll(c context.Context, db *pgxpool.Pool, menu, customerId string) ([]domain.Transaction, int, error)
}

type TransactionQueryImpl struct {
	mu sync.Mutex
}

func NewTransaction() TransactionQuery {
	return &TransactionQueryImpl{}
}

func (repository *TransactionQueryImpl) CreateTrasaction(c context.Context, tx pgx.Tx, trasaction domain.Transaction) error {
	repository.mu.Lock()
	defer repository.mu.Unlock()
	query := `INSERT INTO transactions 
	(
		"id", 
		"customer_id",
		"menu",
		"price",
		"qty",
		"payment",
		"total",
		"created_at"
	) 
	VALUES ($1,$2,$3,$4,$5,$6,$7,$8)`

	_, err := tx.Exec(c,
		query,
		trasaction.Id,
		trasaction.CustomerId,
		trasaction.Menu,
		trasaction.Price,
		trasaction.Qty,
		trasaction.Payment,
		trasaction.Total,
		trasaction.CreatedAt,
	)

	return err
}

func (repository *TransactionQueryImpl) FindByID(c context.Context, db *pgxpool.Pool, id string) (domain.Transaction, error) {
	query := `
		SELECT
		t.id, 
		t.customer_id,
		t.menu,
		t.price,
		t.qty,
		t.payment,
		t.total,
		t.created_at
		FROM transactions AS t
		WHERE t.id=$1`
	row := db.QueryRow(c, query, id)

	var data domain.Transaction
	err := row.Scan(&data.Id, &data.CustomerId, &data.Menu, &data.Price, &data.Qty, &data.Payment, &data.Total, &data.CreatedAt)
	return data, err
}

func (repository *TransactionQueryImpl) FindAll(c context.Context, db *pgxpool.Pool, menu, customerId string) ([]domain.Transaction, int, error) {

	// Initial query without filters
	query := fmt.Sprintf(
		`SELECT 		
            t.id, 
            t.customer_id,
            t.menu,
            t.price,
            t.qty,
            t.payment,
            t.total,
            t.created_at
        FROM transactions AS t
        WHERE 1=1
        `,
	)

	if customerId != "" {
		query += fmt.Sprintf("AND t.customer_id = '%s' ", customerId)
	}

	if menu != "" {
		query += fmt.Sprintf("AND t.menu = '%s' ", menu)
	}

	rows, err := db.Query(c, query)
	if err != nil {
		return []domain.Transaction{}, 0, err
	}
	defer rows.Close()

	var datas []domain.Transaction
	for rows.Next() {
		var data domain.Transaction
		err := rows.Scan(&data.Id, &data.CustomerId, &data.Menu, &data.Price, &data.Qty, &data.Payment, &data.Total, &data.CreatedAt)
		if err != nil {
			return []domain.Transaction{}, 0, err
		}
		datas = append(datas, data)
	}

	return datas, 0, nil
}
