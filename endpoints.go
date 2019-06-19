package wallet

import (
	"context"

	"github.com/go-kit/kit/endpoint"
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
		return GetAllPaymentsResponse{
			Payments: payments,
		}, err
	}
}

// MakeGetAllAccountsEndpoint creates a GetAllAccounts endpoint handler
func MakeGetAllAccountsEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		accounts, err := s.GetAllAccounts(ctx)
		return GetAllAccountsResponse{
			Accounts: accounts,
		}, err
	}
}

// MakePostPaymentEndpoint creates a PostPayment endpoint handler
func MakePostPaymentEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(PostPaymentRequest)
		err = s.PostPayment(ctx, req.Payment)
		return nil, err
	}
}

// API data structures
type (
	// PostPaymentRequest is a request structure for the PostPayment endpoint
	PostPaymentRequest struct {
		Payment Payment `json:"payment"`
	}

	// GetAllPaymentsResponse  is a request structure for the GetAllPayments endpoint
	GetAllPaymentsResponse struct {
		Payments []Payment `json:"payments"`
	}

	// GetAllAccountsResponse is a request structure for the GetAllAccounts endpoint
	GetAllAccountsResponse struct {
		Accounts []Account `json:"accounts"`
	}
)
