package postgres

import (
	"context"
	"database/sql"
	"errors"
	"github.com/Timasha/AuthService/internal/entities"
	"github.com/Timasha/AuthService/pkg/errlist"
)

const (
	queryUpdateUser = `
	update users
	set id = $3,
	    login = $4,
	    password = $5,
	    otp_enabled = $6,
	    otp_key = $7,
	    role_id = $8
	where login = $1 or id = $2
	returning 1
`
)

func (p *PostgresStorage) UpdateUserByLogin(
	ctx context.Context,
	login string,
	user entities.User,
) (err error) {
	_, err = p.conn.ExecContext(
		ctx,
		queryUpdateUser,
		login,
		"",
		user.ID,
		user.Login,
		user.Password,
		user.OtpEnabled,
		user.OtpKey,
		user.Role.ID,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return errlist.ErrUserNotExists
	}

	return err
}
