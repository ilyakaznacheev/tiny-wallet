package wallet

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ilyakaznacheev/tiny-wallet/internal/model"
	"github.com/ilyakaznacheev/tiny-wallet/pkg/currency"
)

// HTTPError is an error with HTTP status and error description
type HTTPError struct {
	code int
	text string
}

// NewHTTPErrorf creates a new HTTP error based on HTTP code and formatted string
// code: HTTP status code
// format: string formatting pattern
// a: formatting attributes
func NewHTTPErrorf(code int, format string, a ...interface{}) *HTTPError {
	return &HTTPError{
		code: code,
		text: fmt.Sprintf(format, a...),
	}
}

// Error returns an text description of the error
func (e HTTPError) Error() string {
	return fmt.Sprintf("%s: %s", http.StatusText(e.code), e.text)
}

// Code returns a HTTP status code of the error
func (e HTTPError) Code() int {
	return e.code
}

// Service is a set of CRUD operations that the backend can process
type Service interface {
	GetAllPayments(ctx context.Context) ([]model.Payment, error)
	GetAllAccounts(ctx context.Context) ([]model.Account, error)
	PostPayment(ctx context.Context, from, to string, amount float64) error
}

// Database is a common interface for a database layer
type Database interface {
	GetAllAccounts() ([]model.Account, error)
	GetAllPayments() ([]model.Payment, error)
	GetAccount(accountID string) (*model.Account, error)
	CreatePayment(p model.Payment) error
}

type walletService struct {
	db Database
}

// NewWalletService creates a new wallet service
func NewWalletService(db Database) Service {
	return &walletService{db}
}

// GetAllPayments returns a list of all payments in the system
func (s *walletService) GetAllPayments(ctx context.Context) ([]model.Payment, error) {
	return s.db.GetAllPayments()
}

// GetAllAccounts returns a list of all accounts in the system
func (s *walletService) GetAllAccounts(ctx context.Context) ([]model.Account, error) {
	return s.db.GetAllAccounts()
}

// PostPayment processes a financial transaction between two accounts
func (s *walletService) PostPayment(ctx context.Context, fromID, toID string, amount float64) error {
	accFrom, err := s.db.GetAccount(fromID)
	if err != nil {
		return NewHTTPErrorf(http.StatusBadRequest, "account %s not found", fromID)
	}

	accTo, err := s.db.GetAccount(toID)
	if err != nil {
		return NewHTTPErrorf(http.StatusBadRequest, "account %s not found", toID)
	}

	// check if the payer and the receiver have the same balance currency
	if accFrom.Currency != accTo.Currency {
		return NewHTTPErrorf(http.StatusBadRequest, "account %s and %s have different balance currencies, payment can't be processed", accFrom.ID, accTo.ID)
	}

	intAmount := currency.ConvertToInternal(amount, accFrom.Currency)

	// check if the payer has enough money on the balance
	if accFrom.Balance < intAmount {
		return NewHTTPErrorf(http.StatusBadRequest, "account %d has not enough money", accFrom.ID)
	}

	payment := model.Payment{
		AccFromID: fromID,
		AccToID:   toID,
		Amount:    intAmount,
	}

	return s.db.CreatePayment(payment)
}
