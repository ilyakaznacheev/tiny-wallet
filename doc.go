/*
Package wallet is a simplistic and minimalistic distributed wallet microservice

Tiny Wallet is a service that allows you to process
any payments between simple financial accounts.

Tiny Wallet helps you to serve simple payment between accounts.
As a internal microservice (without direct access to outside)
it hasn't authentication capabilities, but they can be easily
implemented as a go-kit middleware.

The service is thread-safe and lock-free scalable application,
so it can run multiple replicas over any load balancer without concurrent problems.

As a cloud-native application it can run on any cloud platform
(if it doesn't support Go, you can just build it before deploy,
as Go supports cross-compilation), but also Docker or K8s.

Overview

The project is based on go-kit (https://github.com/go-kit/kit).
Thus, the application consists of several main levels combining together with interfaces.
There are several levels, implemented in this app:

- transport: HTTP routing and the way the app communicate with the outer world:  transport.go (https://github.com/ilyakaznacheev/tiny-wallet/blob/master/transport.go)

- endpoint: the implementation of each route handling that calls underlying business logic:  endpoints.go (https://github.com/ilyakaznacheev/tiny-wallet/blob/master/endpoints.go)

- service: main business logic implementation:
service.go (https://github.com/ilyakaznacheev/tiny-wallet/blob/master/service.go)

- database: communication with database and transaction control:
internal/database/postgres.go (https://github.com/ilyakaznacheev/tiny-wallet/blob/master/internal/database/postgres.go)

Running the app

To run the service call the entry-point CLI app

	cmd/tiny-wallet/wallet.go

You can just run it with

	go run cmd/tiny-wallet/wallet.go

or compile into an executable with

	go build cmd/tiny-wallet/wallet.go

Usage and help

To get help run the app with -h flag. You will get a list of command-line arguments and a list of used environment variables.

	go run cmd/tiny-wallet/wallet.go -h
*/
package wallet
