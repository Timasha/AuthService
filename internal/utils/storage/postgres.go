package storage

import (
	"AuthService/internal/logic"
	"AuthService/internal/logic/models"
	"context"
	"errors"
	"io"
	"os"
	"time"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresStorage struct {
	Pool *pgxpool.Pool
}

func NewPostgresStorage() (p *PostgresStorage) {
	p = &PostgresStorage{}
	return
}

type PostgresStorageConfig struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Ip       string `json:"ip"`
	Port     string `json:"port"`
}

func (p *PostgresStorage) Connect(ctx context.Context, conf PostgresStorageConfig) error {
	var (
		err  error
		conn *pgxpool.Pool
	)

	conn, err = pgxpool.New(ctx, "postgres://"+conf.Login+":"+conf.Password+"@"+conf.Ip+":"+conf.Port+"/auth")
	p.Pool = conn
	for i := 0; err != nil && i <= 10; i++ {
		conn, err = pgxpool.New(ctx, "postgres://"+conf.Login+":"+conf.Password+"@"+conf.Ip+":"+conf.Port+"/auth")
		p.Pool = conn
		time.Sleep(time.Second)
	}
	return err
}
func (p *PostgresStorage) Close() {
	p.Pool.Close()
}
func (p *PostgresStorage) CreateUser(ctx context.Context, user models.User) error {
	var retLogin string
	err := p.Pool.QueryRow(ctx, "insert into users values($1, $2, $3, $4,$5,$6) returning Login;", user.UserID, user.Login, user.Password, user.OtpEnabled, user.OtpKey, user.Role.RoleId).Scan(&retLogin)

	if err == pgx.ErrNoRows && err != nil {
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) {
			if pgErr.Code == pgerrcode.UniqueViolation {
				return logic.ErrUserAlreadyExists
			}
		}
	}
	return err
}
func (p *PostgresStorage) GetUserByLogin(ctx context.Context, login string) (models.User, error) {
	var user models.User
	err := p.Pool.QueryRow(ctx, "select users.UserId, users.Login, users.Password, users.OtpEnabled, users.OtpKey, roles.RoleId, roles.RoleName from users join roles on roles.RoleId = users.RoleId where Login = $1;", login).Scan(&(user.UserID), &(user.Login), &(user.Password), &(user.OtpEnabled), &(user.OtpKey), &(user.Role.RoleId), &(user.Role.RoleName))
	if err == pgx.ErrNoRows {
		return models.User{}, logic.ErrUserNotExists
	}
	return user, err
}

func (p *PostgresStorage) GetUserByUserId(ctx context.Context, userId string)(models.User,error){
	var user models.User
	err := p.Pool.QueryRow(ctx, "select users.UserId, users.Login, users.Password, users.OtpEnabled, users.OtpKey, roles.RoleId, roles.RoleName from users join roles on roles.RoleId = users.RoleId where UserId = $1;", userId).Scan(&(user.UserID), &(user.Login), &(user.Password), &(user.OtpEnabled), &(user.OtpKey), &(user.Role.RoleId), &(user.Role.RoleName))
	if err == pgx.ErrNoRows {
		return models.User{}, logic.ErrUserNotExists
	}
	return user, err
}

func (p *PostgresStorage) UpdateUserByLogin(ctx context.Context, login string, user models.User) error {
	var retLogin string
	err := p.Pool.QueryRow(ctx, "update users set UserId = $1, Login = $2, Password = $3, OtpEnabled = $4, OtpKey = $5 where Login = $6 returning Login;", user.UserID, user.Login, user.Password, user.OtpEnabled, user.OtpKey, login).Scan(&retLogin)
	if err == pgx.ErrNoRows {
		return logic.ErrUserNotExists
	}
	return err
}
func (p *PostgresStorage) DeleteUserByLogin(ctx context.Context, login string) error {
	var retLogin string
	err := p.Pool.QueryRow(ctx, "delete from users where Login = $1 returning Login;", login).Scan(&retLogin)
	if err == pgx.ErrNoRows {
		return logic.ErrUserNotExists
	}
	return err
}

func (p *PostgresStorage) CreateRole(ctx context.Context, role models.Role) error {
	var roleId models.RoleId
	err := p.Pool.QueryRow(ctx, "insert into roles values($1, $2) returning RoleId;", role.RoleId, role.RoleName).Scan(&roleId)

	if err == pgx.ErrNoRows && err != nil {
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) {
			if pgErr.Code == pgerrcode.UniqueViolation {
				return logic.ErrRoleAlreadyExists
			}
		}
	}
	return err
}

func (p *PostgresStorage) GetRoleById(ctx context.Context, roleId models.RoleId) (models.Role, error) {
	var role models.Role
	err := p.Pool.QueryRow(ctx, "select * from roles where RoleId = $1;", roleId).Scan(&(role.RoleId), &(role.RoleName))
	if err == pgx.ErrNoRows {
		return models.Role{}, logic.ErrRoleNotExists
	}
	return role, err
}

func (p *PostgresStorage) UpdateRoleById(ctx context.Context, roleId models.RoleId, role models.Role) error {
	var retId models.RoleId
	err := p.Pool.QueryRow(ctx, "update roles set RoleId = $1, RoleName = $2 where RoleId = $3 returning RoleId;", role.RoleId, role.RoleName, roleId).Scan(&retId)
	if err == pgx.ErrNoRows {
		return logic.ErrRoleNotExists
	}
	return err
}

func (p *PostgresStorage) DeleteRoleById(ctx context.Context, roleId models.RoleId) error {
	var retId models.RoleId
	err := p.Pool.QueryRow(ctx, "delete from roles where RoleId = $1 returning RoleId", roleId).Scan(&retId)
	if err == pgx.ErrNoRows {
		return logic.ErrRoleNotExists
	}
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

		err := p.Pool.QueryRow(ctx, string(fileData)).Scan()
		if err != nil && err != pgx.ErrNoRows {
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

		err := p.Pool.QueryRow(ctx, string(fileData)).Scan()
		if err != nil && err != pgx.ErrNoRows {
			return err
		}
	}

	return nil
}
