create table if not exists roles (
    RoleId bigserial PRIMARY KEY,
    RoleName text UNIQUE
);