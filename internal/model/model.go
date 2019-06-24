// Package model contains common application models and types
package model

import (
	"time"

	"github.com/ilyakaznacheev/tiny-wallet/pkg/currency"
)

// Account is a financial account entity representation
type Account struct {
	ID         string
	LastUpdate *time.Time
	Balance    int
	Currency   currency.Currency
}

// Payment is a financial transaction between accounts
type Payment struct {
	ID        int
	AccFromID string
	AccToID   string
	DateTime  time.Time
	Amount    int
	Currency  currency.Currency
}
