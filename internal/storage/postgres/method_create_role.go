package postgres

import (
	"context"
	"github.com/Timasha/AuthService/internal/entities"
	"github.com/Timasha/AuthService/pkg/errlist"
)

const (
	queryCreateRole = `
	insert into roles values($1, $2) 
	returning role_id;
`
)

func (p *PostgresStorage) CreateRole(ctx context.Context, role entities.Role) (err error) {
	row, err := p.conn.NamedExecContext(ctx, queryCreateRole, role)
	if rowsAff, rowsAffErr := row.RowsAffected(); rowsAff != 1 || rowsAffErr != nil {
		return errlist.ErrUserAlreadyExists
	}

	if err != nil {
		return err
	}

	return nil
}
