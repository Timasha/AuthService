create table if not exists users (
    UserId text PRIMARY KEY,
    Login text UNIQUE,
    Password text
);