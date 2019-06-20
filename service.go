package wallet

import (
	"context"
	"fmt"
	"net/http"
	"time"

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
	GetAllPayments(ctx context.Context) ([]Payment, error)
	GetAllAccounts(ctx context.Context) ([]Account, error)
	PostPayment(ctx context.Context, from, to string, amount int) error
}

// Database is a common interface for a database layer
type Database interface {
	GetAllAccounts() ([]Account, error)
	GetAllPayments() ([]Payment, error)
	GetAccount(accountID string) (*Account, error)
	CreatePayment(p Payment) error
}

// Account is a financial account
type Account struct {
	ID         string `json:"id"`
	LastUpdate time.Time
	Balance    int               `json:"balance"`
	Currency   currency.Currency `json:"currency"`
}

// Payment is a financial transaction between accounts
type Payment struct {
	ID        int
	AccFromID string    `json:"account-from"`
	AccToID   string    `json:"account-to"`
	DateTime  time.Time `json:"time,omitempty"`
	Amount    int       `json:"amount"`
}

type walletService struct {
	db Database
}

// NewWalletService creates a new wallet service
func NewWalletService() Service {
	return &walletService{}
}

// GetAllPayments returns a list of all payments in the system
func (s *walletService) GetAllPayments(ctx context.Context) ([]Payment, error) {
	return s.db.GetAllPayments()
}

// GetAllAccounts returns a list of all accounts in the system
func (s *walletService) GetAllAccounts(ctx context.Context) ([]Account, error) {
	return s.db.GetAllAccounts()
}

// PostPayment processes a financial transaction between two accounts
func (s *walletService) PostPayment(ctx context.Context, fromID, toID string, amount int) error {
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

	// check if the payer has enough money on the balance
	if accFrom.Balance < amount {
		return NewHTTPErrorf(http.StatusBadRequest, "account %d has not enough money", accFrom.ID)
	}

	payment := Payment{
		AccFromID: fromID,
		AccToID:   toID,
		Amount:    amount,
	}

	return s.db.CreatePayment(payment)
}
