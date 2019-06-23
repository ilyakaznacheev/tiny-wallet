package wallet

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/ilyakaznacheev/tiny-wallet/internal/model"
	"github.com/ilyakaznacheev/tiny-wallet/pkg/currency"
)

// HTTPError is an error with an HTTP status code
type HTTPError interface {
	error
	Code() int
}

// ErrHTTPStatus is an error with HTTP status and error description
type ErrHTTPStatus struct {
	code int
	text string
	err  error
}

// NewErrHTTPStatusf creates a new HTTP error based on HTTP code and formatted string
// code: HTTP status code
// err: underlying error to wrap in
// format: string formatting pattern
// a: formatting attributes
func NewErrHTTPStatusf(code int, err error, format string, a ...interface{}) *ErrHTTPStatus {
	return &ErrHTTPStatus{
		code: code,
		text: fmt.Sprintf(format, a...),
		err:  err,
	}
}

// Error returns an text description of the error
func (e ErrHTTPStatus) Error() string {
	// return fmt.Sprintf("%s: %s", http.StatusText(e.code), e.text)
	return e.text
}

// Code returns a HTTP status code of the error
func (e ErrHTTPStatus) Code() int {
	return e.code
}

// Unwrap returns wrapped error
func (e ErrHTTPStatus) Unwrap() error {
	return e.err
}

// Service is a set of CRUD operations that the backend can process
type Service interface {
	GetAllPayments(ctx context.Context) ([]model.Payment, error)
	GetAllAccounts(ctx context.Context) ([]model.Account, error)
	PostPayment(ctx context.Context, from, to string, amount float64) error
	PostAccount(ctx context.Context, id string, balance float64, curr string) error
}

// Database is a common interface for a database layer
type Database interface {
	GetAllAccounts() ([]model.Account, error)
	GetAllPayments() ([]model.Payment, error)
	GetAccount(accountID string) (*model.Account, error)
	CreatePayment(p model.Payment, lastChangedFrom, lastChangedTo *time.Time) error
	CreateAccount(a model.Account) error
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
	if err == sql.ErrNoRows {
		return NewErrHTTPStatusf(http.StatusNotFound, err, "account %s not found", fromID)
	} else if err != nil {
		return NewErrHTTPStatusf(http.StatusInternalServerError, err, "unexpected error")
	}

	accTo, err := s.db.GetAccount(toID)
	if err == sql.ErrNoRows {
		return NewErrHTTPStatusf(http.StatusNotFound, err, "account %s not found", toID)
	} else if err != nil {
		return NewErrHTTPStatusf(http.StatusInternalServerError, err, "unexpected error")
	}

	// check if the payer and the receiver have the same balance currency
	if accFrom.Currency != accTo.Currency {
		return NewErrHTTPStatusf(http.StatusBadRequest, nil, "accounts %s and %s have different balance currencies, payment can't be processed", accFrom.ID, accTo.ID)
	}

	intAmount := currency.ConvertToInternal(amount, accFrom.Currency)

	// check if the payer has enough money on the balance
	if accFrom.Balance < intAmount {
		return NewErrHTTPStatusf(http.StatusBadRequest, nil, "account %s has not enough money", accFrom.ID)
	}

	payment := model.Payment{
		AccFromID: fromID,
		AccToID:   toID,
		Amount:    intAmount,
	}

	err = s.db.CreatePayment(payment, accFrom.LastUpdate, accTo.LastUpdate)
	if err != nil {
		return NewErrHTTPStatusf(http.StatusInternalServerError, err, "payment processing failed")
	}
	return nil
}

func (s *walletService) PostAccount(ctx context.Context, id string, balance float64, curr string) error {
	currKey, err := currency.AtoCurrency(curr)
	if err != nil {
		return NewErrHTTPStatusf(http.StatusBadRequest, err, "can't process account creation with currency %s", curr)
	}

	if balance < 0 {
		return NewErrHTTPStatusf(http.StatusBadRequest, err, "can't process account creation with negative balance %f", balance)
	}
	a := model.Account{
		ID:       id,
		Balance:  currency.ConvertToInternal(balance, *currKey),
		Currency: *currKey,
	}

	err = s.db.CreateAccount(a)
	if err != nil {
		return NewErrHTTPStatusf(http.StatusInternalServerError, err, "account creation failed")
	}
	return nil
}
