package postgres

import "context"

const (
	queryDeleteRole = `
	delete from roles
	where role_id = $2
`
)

func (p *PostgresStorage) DeleteRoleByID(
	ctx context.Context,
	roleID int64,
) (err error) {
	_, err = p.conn.ExecContext(ctx, queryDeleteRole, roleID)

	return err
}
