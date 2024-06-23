package pgx

import (
	"errors"
	"github.com/jackc/pgx/v5/pgconn"
)

func CheckAlreadyExists(err error) bool {
	pgErr := new(pgconn.PgError)
	return err != nil && errors.As(err, &pgErr) && pgErr.Code == "23505"
}
