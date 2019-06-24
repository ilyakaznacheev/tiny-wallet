package wallet

import (
	"context"
	"os"
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/ilyakaznacheev/tiny-wallet/pkg/currency"
)

const (
	urlAPIDoc  = "https://github.com/ilyakaznacheev/tiny-wallet/blob/master/api/api.md"
	urlAPIMain = "https://github.com/ilyakaznacheev/tiny-wallet"
)

// Endpoints is a set of service API endpoints
type Endpoints struct {
	// GetAllPaymentsEndpoint returns all payments in the system
	GetAllPaymentsEndpoint endpoint.Endpoint
	// GetAllAccountsEndpoint returns all accounts in the system
	GetAllAccountsEndpoint endpoint.Endpoint
	// PostPayment processes a new payment
	PostPayment endpoint.Endpoint
	// PostAccount creates a new account
	PostAccount endpoint.Endpoint
	// RedirectMain redirects the user from the main page
	RedirectMain endpoint.Endpoint
	// RedirectAPI redirects the user from the API page
	RedirectAPI endpoint.Endpoint
}

// MakeServerEndpoints creates server handlers for each endpoint
func MakeServerEndpoints(s Service) Endpoints {
	return Endpoints{
		GetAllPaymentsEndpoint: MakeGetAllPaymentsEndpoint(s),
		GetAllAccountsEndpoint: MakeGetAllAccountsEndpoint(s),
		PostPayment:            MakePostPaymentEndpoint(s),
		PostAccount:            MakePostAccountEndpoint(s),
		RedirectAPI:            MakeRedirectAPIEndpoint(s),
		RedirectMain:           MakeRedirectMainEndpoint(s),
	}
}

// MakeGetAllPaymentsEndpoint creates a GetAllPayments endpoint handler
func MakeGetAllPaymentsEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		// call service logic
		payments, err := s.GetAllPayments(ctx)
		if err != nil {
			return nil, err
		}
		// convert results into the response format
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
		// call service logic
		accounts, err := s.GetAllAccounts(ctx)
		if err != nil {
			return nil, err
		}

		// convert results into the response format
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
		// call service logic
		res, err := s.PostPayment(ctx, req.AccountFromID, req.AccountToID, req.Amount)
		if err != nil {
			return nil, err
		}

		// convert results into the response format
		payment := Payment{
			AccFromID: res.AccFromID,
			AccToID:   res.AccToID,
			DateTime:  res.DateTime,
			Amount:    currency.ConvertToExternal(res.Amount, res.Currency),
			Currency:  res.Currency,
		}
		return &payment, nil
	}
}

// MakePostAccountEndpoint creates a PostAccount endpoint handler
func MakePostAccountEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(PostAccountRequest)
		// call service logic
		res, err := s.PostAccount(ctx, req.ID, req.Balance, req.Currency)
		if err != nil {
			return nil, err
		}

		// convert results into the response format
		account := Account{
			ID:       res.ID,
			Balance:  currency.ConvertToExternal(res.Balance, res.Currency),
			Currency: res.Currency,
		}
		return &account, nil
	}
}

// MakeRedirectAPIEndpoint redirects to api documentation page
func MakeRedirectAPIEndpoint(s Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (response interface{}, err error) {
		// redirect to API documentation
		redirectURL := urlAPIDoc
		return &redirectURL, nil
	}
}

// MakeRedirectMainEndpoint redirects to main project page or preconfigured redirect link
func MakeRedirectMainEndpoint(s Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (response interface{}, err error) {
		// redirect to main project page or specified page
		redirectURL := urlAPIMain
		if redirectEnv := os.Getenv("REDIRECT_MAIN"); redirectEnv != "" {
			redirectURL = redirectEnv
		}
		return &redirectURL, nil
	}
}

// API data structures

type (
	// PostPaymentRequest is a request structure for the PostPayment endpoint
	PostPaymentRequest struct {
		AccountFromID string  `json:"account-from"`
		AccountToID   string  `json:"account-to"`
		Amount        float64 `json:"amount"`
	}

	// PostAccountRequest is a request structure for the PostAccount endpoint
	PostAccountRequest struct {
		ID       string  `json:"id"`
		Balance  float64 `json:"balance"`
		Currency string  `json:"currency"`
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
