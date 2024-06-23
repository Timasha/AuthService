package postgres

import (
	"context"
	"github.com/Timasha/AuthService/internal/entities"
	"github.com/Timasha/AuthService/pkg/errlist"
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
	row, err := p.conn.NamedExecContext(ctx, queryCreateUser, user)
	if rowsAff, rowsAffErr := row.RowsAffected(); rowsAff != 1 || rowsAffErr != nil {
		return errlist.ErrUserAlreadyExists
	}

	if err != nil {
		return err
	}

	return nil
}
