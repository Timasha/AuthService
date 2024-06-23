package postgres

import (
	"context"
	"database/sql"
	"errors"
	"github.com/Timasha/AuthService/internal/entities"
	"github.com/Timasha/AuthService/pkg/errlist"
)

const (
	queryGetRole = `
	select id as role_id,
	       access as role_access,
	       name as role_name
	from roles
	where id = $1 or role_name = $2
`
)

func (p *PostgresStorage) GetRoleByID(
	ctx context.Context,
	roleID int64,
) (role entities.Role, err error) {
	err = p.conn.GetContext(ctx, &role, queryGetRole, roleID, "")
	if errors.Is(err, sql.ErrNoRows) {
		return role, errlist.ErrRoleNotExists
	}

	return role, err
}
