create table if not exists users (
    user_id     integer PRIMARY KEY not null,
    login       text UNIQUE not null,
    password    text not null,
    otp_enabled boolean not null,
    otp_key     text not null,
    role_id     bigserial not null REFERENCES roles (role_id)
);