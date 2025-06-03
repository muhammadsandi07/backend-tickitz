package models

import "github.com/jackc/pgx/v5/pgconn"

type ResultCommand struct {
	ResultFirst  pgconn.CommandTag
	ResultSecond pgconn.CommandTag
}
