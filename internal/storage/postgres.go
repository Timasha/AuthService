package storage

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/Timasha/AuthService/internal/entities"
	"github.com/Timasha/AuthService/internal/usecase"
	"github.com/Timasha/AuthService/pkg/errlist"
	"github.com/Timasha/AuthService/utils/consts"
	"github.com/gofiber/fiber/v2/log"

	//nolint:revive // it's ok
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

type PostgresStorage struct {
	conn *sqlx.DB
	cfg  PostgresStorageConfig

	usecase.UserStorage
	usecase.RolesStorage
}

func NewPostgresStorage(cfg PostgresStorageConfig) (p *PostgresStorage) {
	p = &PostgresStorage{
		cfg: cfg,
	}

	return
}

type PostgresStorageConfig struct {
	Login    string `validate:"required"`
	Password string `validate:"required"`
	IP       string `validate:"required"`
	Port     string `validate:"required"`
}

func (p *PostgresStorage) Start(ctx context.Context) error {
	conn, err := sqlx.ConnectContext(
		ctx,
		"pgx",
		fmt.Sprintf(consts.SqlxConnectFmt,
			p.cfg.Login,
			p.cfg.Password,
			p.cfg.IP,
			p.cfg.Port,
		),
	)
	if err != nil {
		log.Errorf("Start postgres error: %s", err.Error())
	}

	p.conn = conn
	for i := 0; err != nil && i <= 10; i++ {
		conn, err = sqlx.ConnectContext(
			ctx,
			"pgx",
			fmt.Sprintf(consts.SqlxConnectFmt,
				p.cfg.Login,
				p.cfg.Password,
				p.cfg.IP,
				p.cfg.Port,
			),
		)
		if err != nil {
			log.Errorf("Start postgres error: %s", err.Error())
		}

		p.conn = conn

		time.Sleep(time.Second)
	}

	return err
}

func (p *PostgresStorage) Stop(_ context.Context) error {
	return p.conn.Close()
}

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
func (p *PostgresStorage) DeleteUserByLogin(ctx context.Context, login string) (err error) {
	_, err = p.conn.ExecContext(ctx, queryDeleteUser, login, "")
	if errors.Is(err, sql.ErrNoRows) {
		return errlist.ErrUserNotExists
	}

	return err
}

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

func (p *PostgresStorage) DeleteRoleByID(
	ctx context.Context,
	roleID int64,
) (err error) {
	_, err = p.conn.ExecContext(ctx, queryDeleteUser, roleID, "")

	return err
}

func (p *PostgresStorage) MigrateUp(ctx context.Context, migrationsPath string) error {
	{
		file, openErr := os.Open(migrationsPath + "/roles_storage_up.sql")
		if openErr != nil {
			return openErr
		}
		fileData, readErr := io.ReadAll(file)

		if readErr != nil {
			return readErr
		}

		err := p.conn.QueryRowContext(ctx, string(fileData)).Scan()
		if err != nil && err != sql.ErrNoRows {
			return err
		}
	}
	{
		file, openErr := os.Open(migrationsPath + "/user_storage_up.sql")
		if openErr != nil {
			return openErr
		}
		fileData, readErr := io.ReadAll(file)

		if readErr != nil {
			return readErr
		}

		err := p.conn.QueryRowContext(ctx, string(fileData)).Scan()
		if err != nil && err != sql.ErrNoRows {
			return err
		}
	}

	return nil
}
