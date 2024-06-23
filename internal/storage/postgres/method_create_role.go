package postgres

import (
	"context"
	"github.com/Timasha/AuthService/internal/entities"
	"github.com/Timasha/AuthService/pkg/errlist"
	"github.com/Timasha/AuthService/utils/pgx"
)

const (
	queryCreateRole = `
	insert into roles values($1, $2) 
	returning role_id;
`
)

func (p *PostgresStorage) CreateRole(ctx context.Context, role entities.Role) (err error) {
	_, err = p.conn.NamedExecContext(ctx, queryCreateRole, role)
	if pgx.CheckAlreadyExists(err) {
		return errlist.ErrRoleAlreadyExists
	}

	if err != nil {
		return err
	}

	return nil
}
