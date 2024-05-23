create table if not exists roles (
    id integer PRIMARY KEY,
    access text,
    name text UNIQUE
);