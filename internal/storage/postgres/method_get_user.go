package postgres

import (
	"context"
	"database/sql"
	"errors"
	"github.com/Timasha/AuthService/internal/entities"
	"github.com/Timasha/AuthService/pkg/errlist"
)

const (
	queryGetUser = `
	select users.user_id as user_id,
	       users.login as login,
	       users.password as password,
	       users.otp_enabled as otp_enabled,
	       users.otp_key as otp_key,
	       roles.id as role_id,
	       roles.access as role_access,
	       roles.name as role_name
	from users join roles on roles.role_id = users.role_id
	where login = $1 or user_id = $2
`
)

func (p *PostgresStorage) GetUserByUserID(
	ctx context.Context,
	userID int64,
) (user entities.User, err error) {
	err = p.conn.GetContext(ctx, &user, queryGetUser, "", userID)

	switch {
	case errors.Is(err, sql.ErrNoRows):
		return user, errlist.ErrUserNotExists
	case errors.Is(err, nil):
		return user, nil
	default:
		return user, err
	}
}

func (p *PostgresStorage) GetUserByLogin(
	ctx context.Context,
	login string,
) (user entities.User, err error) {
	err = p.conn.GetContext(ctx, &user, queryGetUser, login, "")

	switch {
	case errors.Is(err, sql.ErrNoRows):
		return user, errlist.ErrUserNotExists
	case errors.Is(err, nil):
		return user, nil
	default:
		return user, err
	}
}
