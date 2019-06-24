// Package database contains Database interface implementation for certain database.
//
// Here is a PostgreSQL implementation.
//
// To connect to the Postgres database call method `NewPostgresClient()` that will initiate database connection
package database

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/ilyakaznacheev/tiny-wallet/internal/model"
	"github.com/lib/pq"
	"golang.org/x/xerrors"
)

// PostgresClient is a database communication manager
type PostgresClient struct {
	db  *sql.DB
	ctx context.Context
}

// NewPostgresClient create a new database communication manager
//
// - ctx: context of the database client. Can be used to interrupt connection wait loop or long transactions;
// - options: database connection options. Please provide a string of options in format `host=localhost port=5432 ...`. For more information about possible options see [Database Connection Control Functions](https://www.postgresql.org/docs/current/libpq-connect.html);
// - wait: describes will the app wait until the database will up or fails after first unsuccessful ping. Useful for orchestration environments like K8s or Docker Compose or Swarm to wait when the database container of proxy will up.
func NewPostgresClient(ctx context.Context, options string, wait bool) (*PostgresClient, error) {
	db, err := sql.Open("postgres", options)
	if err != nil {
		return nil, err
	}

	// try to ping the database
	err = db.Ping()
	if err != nil {
		if !wait {
			return nil, err
		}
		log.Println(err)
		// wait until the database will up
		itr := 0
	db_wait:
		for {
			itr++
			log.Printf("waiting for a database connection... [%d]\n", itr)
			select {
			case <-ctx.Done():
				return nil, ctx.Err()
			case <-time.After(1 * time.Second):
				err = db.Ping()
				if err == nil {
					break db_wait
				}
			}
		}
	}

	return &PostgresClient{
		db:  db,
		ctx: ctx,
	}, nil
}

// GetAllAccounts returns a list of existing accounts.
//
// The view v_accounts calculates a sum of account balance and following payments affecting this account.
//
// To improve database performance you can periodically calculate a sum op payments related to each account and update its fields `balance` and `balance_date`. Thus, the payments older than balance_date will not be affected in aggregations anymore. All dates should be in UTC+0.
func (pg *PostgresClient) GetAllAccounts() ([]model.Account, error) {
	// fetch the data
	rows, err := pg.db.Query(
		`SELECT *
			FROM v_accounts`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	// process the result
	res := make([]model.Account, 0)

	for rows.Next() {
		rec := model.Account{}
		if err := rows.Scan(&rec.ID, &rec.LastUpdate, &rec.Balance, &rec.Currency); err != nil {
			return nil, err
		}
		res = append(res, rec)
	}

	return res, nil
}

// GetAllPayments returns a list of existing payments in historical order
//
// Since the payment doesn't contain currency code, it will be received from the corresponding payer account
func (pg *PostgresClient) GetAllPayments() ([]model.Payment, error) {
	// fetch the data
	rows, err := pg.db.Query(
		`SELECT p.*, a.currency
			FROM payments AS p
				INNER JOIN accounts AS a ON
					a.id = p.account_from_id
			ORDER BY account_from_id, trx_time`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	// process the result
	res := make([]model.Payment, 0)

	for rows.Next() {
		rec := model.Payment{}
		if err := rows.Scan(&rec.ID, &rec.AccFromID, &rec.AccToID, &rec.DateTime, &rec.Amount, &rec.Currency); err != nil {
			return nil, err
		}
		res = append(res, rec)
	}

	return res, nil
}

// GetAccount returns an existing account.
//
// The view v_accounts calculates a sum of account balance and following payments affecting this account.
//
// To improve database performance you can periodically calculate a sum op payments related to each account and update its fields `balance` and `balance_date`. Thus, the payments older than balance_date will not be affected in aggregations anymore. All dates should be in UTC+0.
func (pg *PostgresClient) GetAccount(accountID string) (*model.Account, error) {
	// fetch the data
	row := pg.db.QueryRow(`
		SELECT *
			FROM v_accounts
			WHERE
				id = $1`, accountID)

	// process the result
	rec := model.Account{}
	if err := row.Scan(&rec.ID, &rec.LastUpdate, &rec.Balance, &rec.Currency); err != nil {
		return nil, err
	}

	return &rec, nil
}

// CreatePayment tries to create a financial transaction
// Concurrent data access is managed by means of MVCC (Multiversion Concurrency Control)
// In case of any inconsistency, race condition or any other concurrency problem it raises an error
func (pg *PostgresClient) CreatePayment(p model.Payment, lastChangedFrom, lastChangedTo *time.Time) (*model.Payment, error) {
	now := time.Now()
	// get pg transaction
	tx, err := pg.db.BeginTx(pg.ctx, &sql.TxOptions{
		Isolation: sql.LevelSerializable,
		ReadOnly:  false,
	})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// try to update the payer account if it wasn't updated from any concurrent process
	_, err = tx.Exec(`
		UPDATE accounts SET
			last_update = $1
		WHERE
			id = $2 AND
			last_update = $3`, now, p.AccFromID, lastChangedFrom)
	if err != nil {
		return nil, err
	}

	// try to update the receiver account if it wasn't updated from any concurrent process
	_, err = tx.Exec(`
		UPDATE accounts SET
			last_update = $1
		WHERE
			id = $2 AND
			last_update = $3`, now, p.AccToID, lastChangedTo)
	if err != nil {
		return nil, err
	}

	// create a new payment
	row := tx.QueryRow(`
		INSERT INTO payments (account_from_id, account_to_id, amount, trx_time)
			VALUES($1, $2, $3, $4)
			RETURNING *`,
		p.AccFromID, p.AccToID, p.Amount, now)

	rec := model.Payment{}

	if err := row.Scan(&rec.ID, &rec.AccFromID, &rec.AccToID, &rec.DateTime, &rec.Amount); err != nil {
		return nil, err
	}

	// commit changes
	return &rec, tx.Commit()
}

// CreateAccount creates a new account.
//
// If the account already exists, the method will return `model.ErrRowExists` error
func (pg *PostgresClient) CreateAccount(a model.Account) (*model.Account, error) {
	now := time.Now()
	row := pg.db.QueryRow(`
		INSERT INTO accounts (id, last_update, currency, balance, balance_date)
			VALUES($1, $2, $3, $4, $5)
			RETURNING *`,
		a.ID, now, a.Currency, a.Balance, now)

	var (
		rec       model.Account
		dateDummy time.Time
	)

	if err := row.Scan(&rec.ID, &rec.LastUpdate, &rec.Currency, &rec.Balance, &dateDummy); err != nil {
		var pqErr *pq.Error
		if xerrors.As(err, &pqErr) {
			// check Postgres errors class
			switch pqErr.Code.Class() {
			case "23": //integrity_constraint_violation
				return nil, model.ErrRowExists
			}
		}
		return nil, err
	}

	return &rec, nil
}
