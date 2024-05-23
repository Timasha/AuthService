create table if not exists users (
    user_id integer PRIMARY KEY,
    login text UNIQUE,
    password text,
    otp_enabled boolean,
    otp_key text,
    role_id bigserial REFERENCES roles (role_id)
);