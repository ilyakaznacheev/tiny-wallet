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
	"github.com/ilyakaznacheev/cleanenv"
	wallet "github.com/ilyakaznacheev/tiny-wallet"
	"github.com/ilyakaznacheev/tiny-wallet/internal/config"
	"github.com/ilyakaznacheev/tiny-wallet/internal/database"
)

type args struct {
	Config string
}

// run application
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	errs := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		exit := <-c
		cancel()
		errs <- fmt.Errorf("%s", exit)
	}()

	logger := log.NewLogfmtLogger(os.Stderr)

	var conf config.MainConfig

	a := parseArgs(&conf)
	err := cleanenv.ReadConfig(a.Config, &conf)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	var dbConfigURL string
	if conf.Database.DatabaseURL != nil {
		dbConfigURL = *conf.Database.DatabaseURL
	} else {
		dbConfigURL = fmt.Sprintf("host=%s port=%s user=%s "+
			"password=%s dbname=%s sslmode=%s",
			conf.Database.Host, conf.Database.Port, conf.Database.Username, conf.Database.Password, conf.Database.Database, conf.Database.SSL)
	}
	db, err := database.NewPostgresClient(ctx, dbConfigURL, conf.Database.ConnectionWait)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	s := wallet.NewWalletService(db)

	h := wallet.MakeHTTPHandler(s, log.With(logger, "component", "HTTP"))

	address := fmt.Sprintf("%s:%s", conf.Server.Host, conf.Server.Port)

	go func() {
		logger.Log("transport", "HTTP", "addr", address)
		errs <- http.ListenAndServe(address, h)
	}()

	logger.Log("exit", <-errs)
}

func parseArgs(conf interface{}) args {
	var a args

	f := flag.NewFlagSet("Tiny Wallet", 1)
	f.StringVar(&a.Config, "c", "configs/config.yml", "path to configuration file")

	fu := f.Usage
	f.Usage = func() {
		fu()
		envHelp, _ := cleanenv.GetDescription(conf, nil)
		fmt.Println()
		fmt.Println(envHelp)
	}

	f.Parse(os.Args[1:])

	return a
}
