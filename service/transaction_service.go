package service

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/dhimweray222/test-be-rekadigital/exception"
	"github.com/dhimweray222/test-be-rekadigital/model/domain"
	"github.com/dhimweray222/test-be-rekadigital/model/web"
	"github.com/dhimweray222/test-be-rekadigital/repository"
)

type TransactionService interface {
	CreateTransaction(ctx context.Context, request web.CreateTransactionRequest) (domain.Transaction, error)
	FindAll(ctx context.Context, menu, customerId string) (result []domain.Transaction, totalData int, err error)
}

type transactionService struct {
	transactionRepository repository.TransactionRepository
}

func NewTransactionService(ticketRepository repository.TransactionRepository) TransactionService {
	return &transactionService{
		transactionRepository: ticketRepository,
	}
}

func (s *transactionService) CreateTransaction(c context.Context, request web.CreateTransactionRequest) (domain.Transaction, error) {

	transaction := domain.Transaction{
		CustomerId: request.CustomerId,
		Menu:       request.Menu,
		Price:      request.Price,
		Qty:        request.Qty,
		Payment:    request.Payment,
		Total:      request.Price * request.Qty,
		CreatedAt:  time.Now(),
	}
	transaction.GenerateID()

	if err := s.transactionRepository.CreateTrasaction(c, transaction); err != nil {
		return domain.Transaction{}, err
	}

	NewTransaction, err := s.transactionRepository.FindByID(c, transaction.Id)
	if err != nil {
		return domain.Transaction{}, exception.ErrorInternalServer(fmt.Sprintf("Successfully created transaction, but failed to get the created transaction. Error: %s", err.Error()))
	}

	return NewTransaction, err
}

func (s *transactionService) FindAll(ctx context.Context, menu, customerId string) (result []domain.Transaction, totalData int, err error) {
	repositoryResponse, total, err := s.transactionRepository.FindAll(ctx, menu, customerId)
	if err != nil {
		if strings.Contains(err.Error(), "no rows") {
			return nil, 0, exception.ErrorNotFound("Job not found")
		} else {
			return nil, 0, err
		}
	}

	return repositoryResponse, total, err
}
