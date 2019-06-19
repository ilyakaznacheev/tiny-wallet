package wallet

import (
	"context"
	"time"

	"github.com/ilyakaznacheev/tiny-wallet/pkg/currency"
)

// Service is a set of CRUD operations that the backend can process
type Service interface {
	GetAllPayments(ctx context.Context) ([]Payment, error)
	// GetOnePayment(ctx context.Context, paymentID int) (Payment, error)
	GetAllAccounts(ctx context.Context) ([]Account, error)
	// GetOneAccount(ctx context.Context, accountID int) (Account, error)
	PostPayment(ctx context.Context, p Payment) error
}

// Account is a financial account
type Account struct {
	ID       int
	Name     string            `json:"id"`
	Balance  int               `json:"balance"`
	Currency currency.Currency `json:"currency"`
}

// Payment is a financial transaction between accounts
type Payment struct {
	ID        int
	AccFromID int               `json:"account-from"`
	AccToID   int               `json:"account-to"`
	DateTime  time.Time         `json:"time"`
	Amount    int               `json:"amount"`
	Currency  currency.Currency `json:"currency"`
}

type walletService struct {
}

// NewWalletService creates a new wallet service
func NewWalletService() Service {
	return &walletService{}
}

// GetAllPayments returns a list of all payments in the system
func (s *walletService) GetAllPayments(ctx context.Context) ([]Payment, error) {
	return nil, nil
}

// GetAllAccounts returns a list of all accounts in the system
func (s *walletService) GetAllAccounts(ctx context.Context) ([]Account, error) {
	return nil, nil
}

// PostPayment processes a financial transaction between two accounts
func (s *walletService) PostPayment(ctx context.Context, p Payment) error {
	return nil
}
