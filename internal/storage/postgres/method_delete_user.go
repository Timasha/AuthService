package postgres

import (
	"context"
	"database/sql"
	"errors"
	"github.com/Timasha/AuthService/pkg/errlist"
)

const (
	queryDeleteUser = `
	delete from users
	where login = $1 or user_id = $2
	returning 1
`
)

func (p *PostgresStorage) DeleteUserByLogin(ctx context.Context, login string) (err error) {
	_, err = p.conn.ExecContext(ctx, queryDeleteUser, login, "")
	if errors.Is(err, sql.ErrNoRows) {
		return errlist.ErrUserNotExists
	}

	return err
}
