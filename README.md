# Tiny Wallet

<img src="/logo.png" alt="Tiny Wallet" title="Tiny Wallet" width="35%" />

A simplistic and minimalistic distributed wallet microservice

[![Go Report Card](https://goreportcard.com/badge/github.com/ilyakaznacheev/tiny-wallet)](https://goreportcard.com/report/github.com/ilyakaznacheev/tiny-wallet) 
[![GoDoc](https://godoc.org/github.com/ilyakaznacheev/tiny-wallet?status.svg)](https://godoc.org/github.com/ilyakaznacheev/tiny-wallet)
[![Build Status](https://travis-ci.org/ilyakaznacheev/tiny-wallet.svg?branch=master)](https://travis-ci.org/ilyakaznacheev/tiny-wallet)
[![Heroku](https://heroku-badge.herokuapp.com/?app=tiny-wallet&root=api/accounts&style=flat&svg=1)](https://tiny-wallet.herokuapp.com/api)
[![Coverage Status](https://coveralls.io/repos/github/ilyakaznacheev/tiny-wallet/badge.svg?branch=master)](https://coveralls.io/github/ilyakaznacheev/tiny-wallet?branch=master)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/ilyakaznacheev/tiny-wallet/blob/master/LICENSE)


## Overview

Tiny Wallet is a service that allows you to note any payments between simple financial accounts.

## Contents

- [About Tiny Wallet](#about-tiny-wallet)
- [Requirements](#requirements)
- [Usage](#usage)
    - [Run Local](#run-local)
    - [Docker Compose](#docker-compose)
    - [Deployment](#deployment)
        - [Heroku](#heroku)
    - [Online](#online)
- [API documentation](#api-documentation)
- [Testing](#testing)

## About Tiny Wallet

## Requirements

## Usage

### Run Local

*Prerequirements*

To run the service locally, you need to have a running PostgreSQL database.

It can be local, dockerized or even cloud database, you only need to provide its connection information to the app.

To set up the database itself, please apply `.sql` files from `migrations` directory in alphanumerical order. This action will create the schema required for the application to run.

*Starting the service*



### Docker Compose

### Deployment

The service in containerized with Docker, so you can use the [deployments/docker/wallet/Dockerfile](/deployments/docker/wallet/Dockerfile) to deploy it to any service that supports Docker containers, e.g. GKE, AWS, Heroku, etc.

#### Heroku

This service is already configured to be deployed on Heroku. You only need to push the repo to you own Heroku app or deploy it other possible way.

### Online

You can try the service online at [tiny-wallet.herokuapp.com/api](https://tiny-wallet.herokuapp.com/api)

## API documentation

Service public API is documented in [plain text](/api/api.md) and [swagger](/api/swagger.yml).

## Testing

