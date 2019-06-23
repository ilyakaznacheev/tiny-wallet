# Service API Documentation

This document describes API of the Tiny Wallet service

## Contents
- [Main information](#main-information)
- [Endpoints](#endpoints)
    - [Accounts](#accounts)
        - [Get Account List](#get-account-list)
        - [Create A New Account](#create-a-new-account)
    - [Payments](#payments)
        - [Get Payment List](#get-payment-list)
        - [Create A New Payment](#create-a-new-payment)
- [Entities](#entities)
    - [PostAccountRequest](#postaccountrequest)
    - [PostPaymentRequest](#postpaymentrequest)
    - [GetAllAccountsResponse](#getallaccountsresponse)
    - [GetAllPaymentsResponse](#getallpaymentsresponse)
    - [Payment](#payment)
    - [Account](#account)
    - [Error](#error)

## Main information

All communication with the service is organized by means of RESTful HTTP requests.

Each HTTP endpoint receives data in JSON format (whereas needed) and returns the responses also in JSON formats .

## Endpoints

Here is the list of endpoints with description of possible operations.

### Accounts

Accounts represent financial account of the payer and/or receiver of financial transaction. Account have balance in one currency.

#### Get Account List

Returns a full list of accounts on the server.

##### Request

Fetching a list of accounts existing on the service.
```
GET /api/accounts
```

No body or query parameters required.

Possible responses:

- `200`: successful operation: [GetAllAccountsResponse](#getallaccountsresponse).
- `404`: not found: [Error](#error).
- `500`: internal server error: [Error](#error).

#### Create A New Account

Adds a new account with some balance if no account with the same id exists.

Creating a new account.
```
POST /api/account
```

Body should contain a JSON structure of type [PostAccountRequest](#postaccountrequest).

Possible responses:

- `200`: successful operation: [Account](#account).
- `400`: bad request: [Error](#error).
- `409`: conflict: [Error](#error).
- `500`: internal server error: [Error](#error).

### Payments

Payment represents financial transaction of money movement between two accounts.

This tiny service supports only transactions in the same currency. Co currency exchange allowed.

#### Get Payment List

Returns a full list of payments on the server.

```
GET: /api/payments
```

No body or query parameters required.

Possible responses:

- `200`: successful operation: [GetAllPaymentsResponse](#getallpaymentsresponse).
- `404`: not found: [Error](#error).
- `500`: internal server error: [Error](#error).

#### Create A New Payment

Creates a new financial transaction of money movement between two accounts.

Balance of the payer's account should not be smaller than a payment amount.

Currency of the payer and the receiver should be the same.

```
POST: /api/payment
```

Body should contain a JSON structure of type [PostPaymentRequest](#postpaymentrequest).

Possible responses:

- `200`: successful operation: [Payment](#payment).
- `400`: bad request: [Error](#error).
- `404`: not found: [Error](#error).
- `500`: internal server error: [Error](#error).

## Entities

This is a description of JSON types used in request and response body as a data structure.

### PostAccountRequest

| Attribute                | Description                                                  | Type     | Optional |
| ------------------------ | ------------------------------------------------------------ | -------- | -------- |
| `id`                     | Account identification number                                | string   | no       |
| `balance`                | Amount of money on the account balance                       | number   | no       |
| `currency`               | Balance currency  (ISO 4216)                                 | string   | no       |

#### Example

```json
{
	"id": "bob123",
	"balance": 100,
	"currency": "USD"
}
```

### PostPaymentRequest

Payment creation request structure.

> **Note:** `amount` should not be greater than a current payer balance.

| Attribute                | Description                                                  | Type     | Optional |
| ------------------------ | ------------------------------------------------------------ | -------- | -------- |
| `account-from`           | Payer's account id                                           | string   | no       |
| `account-to`             | Receivers account id                                         | string   | no       |
| `amount`                 | Payment amount                                               | number   | no       |

#### Example

```json
{
	"account-from": "bob123",
	"account-to": "alice456",
	"amount": 12.25
}
```

### GetAllAccountsResponse

A list of accounts in the system.

| Attribute                | Description         | Type                        | Optional |
| ------------------------ | ------------------- | --------------------------- | -------- |
| `accounts`               | List of accounts    | list of [Payment](#payment) | no       |

#### Example

```json
{
	"accounts": [
		{
			"id": "bob123",
			"balance": 107.02,
			"currency": "USD"
		},
		{
			"id": "alice456",
			"balance": 92.98,
			"currency": "USD"
        }
    ]
}
```

### GetAllPaymentsResponse

A list of payments in the system.

| Attribute                | Description                                   | Type                        | Optional |
| ------------------------ | --------------------------------------------- | --------------------------- | -------- |
| `payments`               | List of payments in chronological order       | list of [Payment](#payment) | no       |

#### Example

```json
{
	"payments": [
		{
			"account-from": "alice456",
			"account-to": "bob123",
			"time": "2019-06-23T00:37:47.998996Z",
			"amount": 12.34,
			"currency": "USD"
		},
		{
			"account-from": "bob123",
			"account-to": "alice456",
			"time": "2019-06-23T01:41:46.944434Z",
			"amount": 2.34,
			"currency": "USD"
        }
    ]
}
```

### Account

Account entity structure.

| Attribute                | Description                                                  | Type     | Optional |
| ------------------------ | ------------------------------------------------------------ | -------- | -------- |
| `id`                     | Account identification number                                | string   | no       |
| `balance`                | Amount of money on the account balance                       | number   | no       |
| `currency`               | Balance currency  (ISO 4216)                                 | string   | no       |

#### Example

```json
{
    "id": "alice456",
    "balance": 92.98,
    "currency": "USD"
}
```

### Payment

Payment entity structure.

| Attribute                | Description                                                  | Type      | Optional |
| ------------------------ | ------------------------------------------------------------ | --------- | -------- |
| `account-from`           | Payer's account id                                           | string    | no       |
| `account-to`             | Receivers account id                                         | string    | no       |
| `time`                   | Transaction time                                             | timestamp | yes      |
| `amount`                 | Payment amount                                               | number    | no       |
| `currency`               | Balance currency  (ISO 4216)                                 | string    | no       |

#### Example

```json
{
    "account-from": "alice456",
    "account-to": "bob123",
    "time": "2019-06-23T00:37:47.998996Z",
    "amount": 12.34,
    "currency": "USD"
}
```

### Error

Error status code and description.

| Attribute            | Description                    | Type                          | Optional |
| -------------------- | ------------------------------ | ----------------------------- | -------- |
| `code`               | HTTP Status Code               | integer                       | no       |
| `error`              | Error message                  | [ErrorMessage](#errormessage) | no       |

#### Example

```json
{
    "code": 400,
    "error": {
        "text": "some error",
        "details": [
            "database error",
            "wrong key"
        ]
    }
}
```

### ErrorMessage

Error text and some additional information.

| Attribute            | Description                                      | Type           | Optional |
| -------------------- | ------------------------------------------------ | -------------- | -------- |
| `text`               | Error text                                       | string         | no       |
| `details`            | Error details - some specific error information  | list of string | yes      |

#### Example

```json
{
    "text": "some error",
    "details": [
        "database error",
        "wrong key"
    ]
}
```
