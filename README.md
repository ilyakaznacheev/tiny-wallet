# Tiny Wallet

<img src="/logo.png" alt="Tiny Wallet" title="Tiny Wallet" width="35%" />

A simplistic and minimalistic distributed wallet microservice

[![Go Report Card](https://goreportcard.com/badge/github.com/ilyakaznacheev/tiny-wallet)](https://goreportcard.com/report/github.com/ilyakaznacheev/tiny-wallet) 
[![GoDoc](https://godoc.org/github.com/ilyakaznacheev/tiny-wallet?status.svg)](https://godoc.org/github.com/ilyakaznacheev/tiny-wallet)
[![Build Status](https://travis-ci.org/ilyakaznacheev/tiny-wallet.svg?branch=master)](https://travis-ci.org/ilyakaznacheev/tiny-wallet)
[![Heroku](https://pyheroku-badge.herokuapp.com/?app=tiny-wallet&root=api&style=flat)](https://tiny-wallet.herokuapp.com/api)
[![Coverage Status](https://codecov.io/github/ilyakaznacheev/tiny-wallet/coverage.svg?branch=master)](https://codecov.io/gh/ilyakaznacheev/tiny-wallet)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](/LICENSE)


## Overview

Tiny Wallet is a service that allows you to process any payments between simple financial accounts.

## Contents

- [About Tiny Wallet](#about-tiny-wallet)
- [Requirements](#requirements)
- [Usage](#usage)
    - [Download](#download)
    - [Run Local](#run-local)
    - [Docker Compose](#docker-compose)
    - [Deployment](#deployment)
        - [Heroku](#heroku)
    - [Online](#online)
- [API documentation](#api-documentation)
- [Testing](#testing)
- [Contributing](#contributing)

## About Tiny Wallet

Tiny Wallet helps you to serve simple payment between accounts. As a internal microservice (without direct access to outside) it hasn't authentication capabilities, but they can be easily implemented as a go-kit middleware.

The service is thread-safe and lock-free scalable application, so it can run multiple replicas over any load balancer without concurrent problems.

As a cloud-native application it can run on any cloud platform (if it doesn't support Go, you can just build it before deploy, as Go supports cross-compilation), but also Docker or K8s. 

## Requirements

There is a list or application requirements for different deployment scenario:

- local run:
    - Go  1.11.x or greater. Not compatible with earlier versions of Go from 1.10 because of `mod` usage;
    - Go mod should be enabled;
    - PosgreSQL (no specific version, it has to support serializable transaction isolation level). You don't have to install the database on your PC, you can also use dockerized or cloud PostgreSQL;
- Docker Compose:
    - Docker;
    - Docker Compose;
- cloud deployment.
    - Go 1.12.x on the platform;
    - Cloud PostgreSQL;
- cloud deployment with Docker:
    - Docker support;
    - Cloud PostgreSQL;

## Usage

Here is a list of possible ways to run the service.

To get a list of CLI flags and environment variables run from the project root directory:

```bash
go run cmd/tiny-wallet/wallet.go -h
```

### Download

You can get the app with `go get`:

```bash
go get github.com/ilyakaznacheev/tiny-wallet
```

or clone the repo into any directory outside `$GOPATH`:

```bash
git clone https://github.com/ilyakaznacheev/tiny-wallet.git
```

### Run Local

*Prerequirements*

To run the service locally, you need to have a running PostgreSQL database.

It can be local, dockerized or even cloud database, you only need to provide its connection information to the app.

To set up the database itself, please apply `.sql` files from `migrations` directory in alphanumerical order. This action will create the schema required for the application to run.

*Starting the service*

Go to the project root directory and run the app by executing the following command:

```bash
make run
```

or by means of `go` if your operating system doesn't support `make`:

```bash
go run cmd/tiny-wallet/wallet.go
```

*Build*

You can also build an executable file to run it. Call

```bash
make build
```

or 

```bash
go build cmd/tiny-wallet/wallet.go
```

### Docker Compose

You can run the whole app infrastructure in the Docker Compose. See [Requirements](#requirements) for this case.

To run the app call from the project root directory:

```bash
docker-compose up
```

This will start the service and the database and bind the app to port `8080`.

To stop the app run

```bash
docker-compose down
```

### Deployment

The service in containerized with Docker, so you can use the [deployments/docker/wallet/Dockerfile](/deployments/docker/wallet/Dockerfile) to deploy it to any service that supports Docker containers, e.g. GKE, AWS, Heroku, etc.

#### Heroku

This service is already configured to be deployed on Heroku. You only need to push the repo to your own Heroku app or deploy it another possible way.

> **Note:** while deploying on Heroku specify environment variable `HEROKU=X` in the Heroku Dashboard. This will allow the app to run some Heroku-specific start-up logic.

### Online

You can try the service online at [tiny-wallet.herokuapp.com/api](https://tiny-wallet.herokuapp.com/api)

## API documentation

Service public API is documented in [plain text](/api/api.md) and [swagger](/api/swagger.yml). Try it in [Swagger Editor](https://editor.swagger.io/)!

## Testing

The business logic of the app and internal libraries are covered with unit-tests. The generated code or simple technical code (like one-liner that call another function) are not covered with tests now.

To run tests on your PC run

```bash
make test
```

## Contributing

The application is open-sourced under the [MIT](/LICENSE) license.

If you will find some error, want to add something or ask a question - feel free to [create an issue](https://github.com/ilyakaznacheev/tiny-wallet/issues) and/or make a [pull request](https://github.com/ilyakaznacheev/tiny-wallet/pulls).

Any contribution is welcome.
