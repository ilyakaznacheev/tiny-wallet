package wallet

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

// MakeHTTPHandler mounts all of the service endpoints into an http.Handler
func MakeHTTPHandler(s Service, logger log.Logger) http.Handler {

	r := mux.NewRouter()
	e := MakeServerEndpoints(s)
	options := []httptransport.ServerOption{
		// httptransport.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		httptransport.ServerErrorEncoder(encodeError),
	}

	r.Methods("GET").Path("/payments").Handler(httptransport.NewServer(
		e.GetAllPaymentsEndpoint,
		decodeDummy,
		encodeResponse,
		options...,
	))

	r.Methods("GET").Path("/accounts").Handler(httptransport.NewServer(
		e.GetAllAccountsEndpoint,
		decodeDummy,
		encodeResponse,
		options...,
	))

	r.Methods("POST").Path("/payment").Handler(httptransport.NewServer(
		e.PostPayment,
		decodePostPaymentRequest,
		encodeResponse,
		options...,
	))

	return r
}

func decodeDummy(_ context.Context, r *http.Request) (request interface{}, err error) {
	return nil, nil
}

func decodePostPaymentRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req PostPaymentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

type errorer interface {
	error() error
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if err, ok := response.(errorer); ok && err.error() != nil {
		encodeError(ctx, err.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func encodeRequest(_ context.Context, req *http.Request, request interface{}) error {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(request)
	if err != nil {
		return err
	}
	req.Body = ioutil.NopCloser(&buf)
	return nil
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
