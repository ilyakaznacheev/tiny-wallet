package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log"
	wallet "github.com/ilyakaznacheev/tiny-wallet"
	"github.com/ilyakaznacheev/tiny-wallet/internal/database"
)

type args struct {
	Host   string
	Port   string
	WaitDB bool
}

// run application
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	errs := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
		cancel()
	}()

	a := parseArgs()

	logger := log.NewLogfmtLogger(os.Stderr)
	db, err := database.NewPostgresClient(ctx, "", a.WaitDB)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	s := wallet.NewWalletService(db)

	h := wallet.MakeHTTPHandler(s, log.With(logger, "component", "HTTP"))

	address := fmt.Sprintf("%s:%s", a.Host, a.Port)

	go func() {
		logger.Log("transport", "HTTP", "addr", address)
		errs <- http.ListenAndServe(address, h)
	}()

	logger.Log("exit", <-errs)
}

func parseArgs() args {
	var a args

	flag.StringVar(&a.Host, "h", "localhost", "server host")
	flag.StringVar(&a.Port, "p", "8080", "server port")
	flag.BoolVar(&a.WaitDB, "db-wait", true, "wait until database up")

	flag.Parse()

	return a
}
