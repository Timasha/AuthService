create table if not exists users (
    UserId text PRIMARY KEY,
    Login text UNIQUE,
    Password text,
    OtpEnabled boolean,
    OtpKey text,
    RoleId bigserial REFERENCES roles (RoleId)
);