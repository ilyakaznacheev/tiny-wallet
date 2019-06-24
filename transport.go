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
	"golang.org/x/xerrors"
)

// MakeHTTPHandler mounts all of the service endpoints into an http.Handler
func MakeHTTPHandler(s Service, logger log.Logger) http.Handler {

	r := mux.NewRouter()
	e := MakeServerEndpoints(s)
	options := []httptransport.ServerOption{
		httptransport.ServerErrorLogger(logger),
		httptransport.ServerErrorEncoder(encodeError),
	}

	r.Methods("GET").Path("/api/payments").Handler(httptransport.NewServer(
		e.GetAllPaymentsEndpoint,
		decodeDummy,
		encodeResponse,
		options...,
	))

	r.Methods("GET").Path("/api/accounts").Handler(httptransport.NewServer(
		e.GetAllAccountsEndpoint,
		decodeDummy,
		encodeResponse,
		options...,
	))

	r.Methods("POST").Path("/api/payment").Handler(httptransport.NewServer(
		e.PostPayment,
		decodePostPaymentRequest,
		encodeResponse,
		options...,
	))

	r.Methods("POST").Path("/api/account").Handler(httptransport.NewServer(
		e.PostAccount,
		decodePostAccountRequest,
		encodeResponse,
		options...,
	))

	r.Path("/api").Handler(httptransport.NewServer(
		e.RedirectAPI,
		decodeDummy,
		encodeRedirect,
		options...,
	))

	r.Path("/").Handler(httptransport.NewServer(
		e.RedirectMain,
		decodeDummy,
		encodeRedirect,
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

func decodePostAccountRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req PostAccountRequest
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

	var code int

	// process error
	switch e := err.(type) {
	case HTTPError:
		code = e.Code()
	default:
		code = http.StatusInternalServerError
	}
	errResp := map[string]interface{}{
		"error": err.Error(),
	}

	// get wrapped errors
	errDescr := make([]string, 0)
	e := err
	for {
		if e = xerrors.Unwrap(e); e == nil {
			break
		}
		errDescr = append(errDescr, e.Error())
	}

	if len(errDescr) > 0 {
		errResp["details"] = errDescr
	}

	// process response data
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(&ErrorResponse{
		Code:  code,
		Error: ErrorResponseMessage{err.Error(), errDescr},
	})
}

func encodeRedirect(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if err, ok := response.(errorer); ok && err.error() != nil {
		encodeError(ctx, err.error(), w)
		return nil
	}
	if url, ok := response.(*string); ok {
		w.Header().Set("Location", *url)
		w.WriteHeader(http.StatusMovedPermanently)
	}
	return nil
}

type (
	// ErrorResponse is a JSON error response structure
	ErrorResponse struct {
		Code  int                  `json:"code"`
		Error ErrorResponseMessage `json:"error"`
	}
	// ErrorResponseMessage is an error message and details
	ErrorResponseMessage struct {
		Text    string   `json:"text"`
		Details []string `json:"details,omitempty"`
	}
)
