package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"text/tabwriter"

	"github.com/go-kit/kit/log"
	wallet "github.com/ilyakaznacheev/tiny-wallet"
	"github.com/ilyakaznacheev/tiny-wallet/internal/config"
	"github.com/ilyakaznacheev/tiny-wallet/internal/database"
	"github.com/kelseyhightower/envconfig"
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

	a := parseArgs()
	conf, err := config.ReadConfig(a.Config)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	var dbConfigURL string
	if conf.Database.DatabaseURL != nil {
		dbConfigURL = *conf.Database.DatabaseURL
	} else {
		dbConfigURL = fmt.Sprintf("host=%s port=%s user=%s "+
			"password=%s dbname=%s",
			conf.Database.Host, conf.Database.Port, conf.Database.Username, conf.Database.Password, conf.Database.Database)
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

func parseArgs() args {
	var a args

	f := flag.NewFlagSet("Tiny Wallet", 1)
	f.StringVar(&a.Config, "c", "config/config.yml", "path to configuration file")
	fu := f.Usage
	f.Usage = func() {
		fu()

		tabs := tabwriter.NewWriter(os.Stdout, 1, 0, 4, ' ', 0)
		envconfig.Usagef("", &config.MainConfig{}, tabs, `
This application is configured via the environment. 
The following environment variables can be used:
{{range .}}
  {{usage_key .}} [{{usage_type .}}]
	{{usage_description .}}
{{end}}`)
		tabs.Flush()
	}

	f.Parse(os.Args[1:])

	return a
}
