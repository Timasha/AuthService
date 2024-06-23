package postgres

import (
	"context"
	"github.com/Timasha/AuthService/internal/entities"
	"github.com/Timasha/AuthService/pkg/errlist"
	"github.com/Timasha/AuthService/utils/pgx"
)

const (
	queryCreateUser = `
	insert into users( login, password, otp_enabled, otp_key, role_id)
	values(:login, :password, :otp_enabled, :otp_key, :role_id)
`
)

func (p *PostgresStorage) CreateUser(
	ctx context.Context,
	user entities.User,
) (err error) {
	_, err = p.conn.NamedExecContext(ctx, queryCreateUser, user)
	if pgx.CheckAlreadyExists(err) {
		return errlist.ErrUserAlreadyExists
	}

	if err != nil {
		return err
	}

	return nil
}
