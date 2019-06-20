package model

import (
	"time"

	"github.com/ilyakaznacheev/tiny-wallet/pkg/currency"
)

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
