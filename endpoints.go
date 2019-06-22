package wallet

import (
	"context"
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/ilyakaznacheev/tiny-wallet/pkg/currency"
)

// Endpoints is a set of service API endpoints
type Endpoints struct {
	GetAllPaymentsEndpoint endpoint.Endpoint
	// GetOnePaymentEndpoint  endpoint.Endpoint
	GetAllAccountsEndpoint endpoint.Endpoint
	// GetOneAccountEndpoint  endpoint.Endpoint
	PostPayment endpoint.Endpoint
}

// MakeServerEndpoints creates server handlers for each endpoint
func MakeServerEndpoints(s Service) Endpoints {
	return Endpoints{
		GetAllPaymentsEndpoint: MakeGetAllPaymentsEndpoint(s),
		GetAllAccountsEndpoint: MakeGetAllAccountsEndpoint(s),
		PostPayment:            MakePostPaymentEndpoint(s),
	}
}

// MakeGetAllPaymentsEndpoint creates a GetAllPayments endpoint handler
func MakeGetAllPaymentsEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		payments, err := s.GetAllPayments(ctx)
		if err != nil {
			return nil, err
		}
		// convert results into the responce format
		res := GetAllPaymentsResponse{
			Payments: make([]Payment, 0, len(payments)),
		}
		for _, p := range payments {
			res.Payments = append(res.Payments, Payment{
				AccFromID: p.AccFromID,
				AccToID:   p.AccToID,
				DateTime:  p.DateTime,
				Amount:    currency.ConvertToExternal(p.Amount, p.Currency),
				Currency:  p.Currency,
			})
		}
		return res, nil
	}
}

// MakeGetAllAccountsEndpoint creates a GetAllAccounts endpoint handler
func MakeGetAllAccountsEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		accounts, err := s.GetAllAccounts(ctx)
		if err != nil {
			return nil, err
		}
		// convert results into the responce format
		res := GetAllAccountsResponse{
			Accounts: make([]Account, 0, len(accounts)),
		}
		for _, a := range accounts {
			res.Accounts = append(res.Accounts, Account{
				ID:       a.ID,
				Balance:  currency.ConvertToExternal(a.Balance, a.Currency),
				Currency: a.Currency,
			})
		}
		return res, nil
	}
}

// MakePostPaymentEndpoint creates a PostPayment endpoint handler
func MakePostPaymentEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(PostPaymentRequest)
		err = s.PostPayment(ctx, req.AccountFromID, req.AccountToID, req.Amount)
		return nil, err
	}
}

// API data structures
type (
	// PostPaymentRequest is a request structure for the PostPayment endpoint
	PostPaymentRequest struct {
		AccountFromID string  `json:"from-account"`
		AccountToID   string  `json:"to-account"`
		Amount        float64 `json:"amount"`
	}

	// GetAllPaymentsResponse  is a request structure for the GetAllPayments endpoint
	GetAllPaymentsResponse struct {
		Payments []Payment `json:"payments"`
	}

	// GetAllAccountsResponse is a request structure for the GetAllAccounts endpoint
	GetAllAccountsResponse struct {
		Accounts []Account `json:"accounts"`
	}

	// Account is a financial account
	Account struct {
		ID       string            `json:"id"`
		Balance  float64           `json:"balance"`
		Currency currency.Currency `json:"currency"`
	}

	// Payment is a financial transaction between accounts
	Payment struct {
		AccFromID string            `json:"account-from"`
		AccToID   string            `json:"account-to"`
		DateTime  time.Time         `json:"time,omitempty"`
		Amount    float64           `json:"amount"`
		Currency  currency.Currency `json:"currency"`
	}
)
