CREATE TABLE payments
(
    id bigserial PRIMARY KEY NOT NULL,
    account_from_id character varying(30) NOT NULL,
    account_to_id character varying(30) NOT NULL,
    trx_time timestamp without time zone NOT NULL DEFAULT CURRENT_DATE,
    amount bigint NOT NULL,
);

CREATE TABLE accounts
(
    id character varying(30) PRIMARY KEY NOT NULL,
    last_update timestamp without time zone,
    currency character varying(3) NOT NULL,
    balance bigint NOT NULL,
    balance_date timestamp without time zone NOT NULL
);

CREATE OR REPLACE VIEW v_accounts AS
SELECT
	a.id, 
	last_update, 
	coalesce((a.balance + sum(p.amount)), a.balance) as balance,
	a.currency
FROM accounts AS a
	LEFT OUTER JOIN 
        (SELECT account_to_id as id, trx_time, amount
            FROM payments 
		UNION SELECT account_from_id as id, trx_time, amount * -1 as amount
            FROM payments) AS p ON
			p.id = a.id AND
			p.trx_time > a.balance_date	
GROUP BY
	a.id,
	a.last_update,
	a.currency;
