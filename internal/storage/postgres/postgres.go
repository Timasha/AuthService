package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/Timasha/AuthService/internal/usecase"
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
	Login          string `validate:"required"`
	Password       string `validate:"required"`
	IP             string `validate:"required"`
	Port           string `validate:"required"`
	MigrationsPath string `validate:"required"`
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
	for i := 0; err != nil; i++ {
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

	err = p.MigrateUp(ctx, p.cfg.MigrationsPath)
	if err != nil {
		return err
	}

	return nil
}

func (p *PostgresStorage) Stop(_ context.Context) error {
	return p.conn.Close()
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
