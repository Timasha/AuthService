package postgres

import (
	"context"
	"database/sql"
	"errors"
	"github.com/Timasha/AuthService/internal/entities"
	"github.com/Timasha/AuthService/pkg/errlist"
)

const (
	queryUpdateRole = `
	update roles
	set role_id = $3,
	    role_access = $4,
	    role_name = $5
	where role_id = $1 or role_name = $2
	returning 1
`
)

func (p *PostgresStorage) UpdateRoleByID(
	ctx context.Context,
	roleID int64,
	role entities.Role,
) (err error) {
	_, err = p.conn.ExecContext(
		ctx,
		queryUpdateRole,
		roleID,
		"",
		role.ID,
		role.Access,
		role.Name,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return errlist.ErrUserNotExists
	}

	return err
}
