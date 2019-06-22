package database

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/ilyakaznacheev/tiny-wallet/internal/model"
	_ "github.com/lib/pq"
)

// PostgresClient is a database communication manager
type PostgresClient struct {
	db *sql.DB
}

// NewPostgresClient create a new database communication manager
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

	return &PostgresClient{db}, nil
}

// GetAllAccounts returns a list of existing accounts
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

// GetAllPayments returns a list of existing payments
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
		if err := rows.Scan(&rec.ID, &rec.AccFromID, &rec.AccToID, &rec.DateTime, &rec.Amount, &res.Currency); err != nil {
			return nil, err
		}
		res = append(res, rec)
	}

	return res, nil
}

// GetAccount returns an existing account
func (pg *PostgresClient) GetAccount(accountID string) (*model.Account, error) {
	// fetch the data
	row := pg.db.QueryRow(`
		SELECT *
			FROM v_accounts
			WHERE
				a.id = ?`, accountID)

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
func (pg *PostgresClient) CreatePayment(p model.Payment) error {
	// get pg transaction
	tx, err := pg.db.BeginTx(context.Background(), &sql.TxOptions{
		Isolation: sql.LevelSerializable,
		ReadOnly:  false,
	})
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// try to update the payer account if it wasn't updated from any concurrent process
	_, err = tx.Exec(`
		UPDATE accounts SET
			last_update = ?
		WHERE
			id = ?`, time.Now(), p.AccFromID)
	if err != nil {
		return err
	}

	// try to update the receiver account if it wasn't updated from any concurrent process
	_, err = tx.Exec(`
		UPDATE accounts SET
			last_update = ?
		WHERE
			id = ?`, time.Now(), p.AccToID)
	if err != nil {
		return err
	}

	// create a new payment
	_, err = tx.Exec(`
		INSERT accounts(id, account_from_id, account_to_id, trx_time)
			VALUES(?, ?, ?, ?)`,
		p.AccFromID, p.AccToID, p.Amount, time.Now())
	if err != nil {
		return err
	}

	// commit changes
	return tx.Commit()
}
