package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log"
	wallet "github.com/ilyakaznacheev/tiny-wallet"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)
	s := wallet.NewWalletService()

	h := wallet.MakeHTTPHandler(s, log.With(logger, "component", "HTTP"))
	a := parseArgs()

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	address := fmt.Sprintf("%s:%s", a.Host, a.Port)

	go func() {
		logger.Log("transport", "HTTP", "addr", address)
		errs <- http.ListenAndServe(address, h)
	}()

	logger.Log("exit", <-errs)
}

type args struct {
	Host string
	Port string
}

func parseArgs() args {
	var a args

	flag.StringVar(&a.Host, "h", "localhost", "server host")
	flag.StringVar(&a.Port, "p", "8080", "server port")

	flag.Parse()

	return a
}