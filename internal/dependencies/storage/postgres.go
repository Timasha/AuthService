package storage

import (
	"auth/internal/logic/errs"
	"auth/internal/logic/models"
	"context"
	"errors"
	"io"
	"os"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresUserStorage struct {
	Pool *pgxpool.Pool
}

type PostgresUserStorageConfig struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Ip       string `json:"ip"`
	Port     string `json:"port"`
}

func Connect(ctx context.Context, conf PostgresUserStorageConfig) (*PostgresUserStorage, error) {
	conn, err := pgxpool.New(ctx, "postgres://"+conf.Login+":"+conf.Password+"@"+conf.Ip+":"+conf.Port+"/users")
	if err != nil {
		return nil, err
	}
	go func(context.Context, *pgxpool.Pool) {
		select {
		case <-ctx.Done():
			conn.Close()
		}
	}(ctx, conn)
	return &PostgresUserStorage{
		Pool: conn,
	}, nil
}

func (p *PostgresUserStorage) CreateUser(ctx context.Context, user models.User) error {
	err := p.Pool.QueryRow(ctx, "insert into users values($1, $2, $3);", user.UserID, user.Login, user.Password).Scan()

	if err != nil {
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) {
			if pgErr.Code == pgerrcode.UniqueViolation {
				return errs.ErrUserAlreadyExists{}
			}
		}
	}
	return err
}
func (p *PostgresUserStorage) GetUserByLogin(ctx context.Context, login string) (models.User, error) {
	var user models.User
	err := p.Pool.QueryRow(ctx, "select * from users where Login = $1;", login).Scan(&(user.UserID), &(user.Login), &(user.Password))
	if err == pgx.ErrNoRows {
		return models.User{}, errs.ErrUserNotExists{}
	}
	return user, err
}

func (p *PostgresUserStorage) UpdateUserByLogin(ctx context.Context, login string, user models.User) error {
	var retLogin string
	err := p.Pool.QueryRow(ctx, "update users set UserID = $1, Login = $2, Password = $3 where Login = $4", user.UserID, user.Login, user.Password, login).Scan(&retLogin)
	if err == pgx.ErrNoRows {
		return errs.ErrUserNotExists{}
	}
	return err
}
func (p *PostgresUserStorage) DeleteUserByLogin(ctx context.Context, login string) error {
	var retLogin string
	err := p.Pool.QueryRow(ctx, "delete from users where Login = $1", login).Scan(&retLogin)
	if err == pgx.ErrNoRows {
		return errs.ErrUserNotExists{}
	}
	return err
}

func (p *PostgresUserStorage) MigrateUp(ctx context.Context, migrationsPath string) error {
	file, openErr := os.Open(migrationsPath + "/user_storage_up.sql")
	if openErr != nil {
		return openErr
	}
	fileData, readErr := io.ReadAll(file)

	if readErr != nil {
		return readErr
	}

	err := p.Pool.QueryRow(ctx, string(fileData)).Scan()
	return err
}
