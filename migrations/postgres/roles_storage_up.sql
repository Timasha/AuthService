create table if not exists roles (
    id      integer PRIMARY KEY not null,
    access  text,
    name    text UNIQUE not null
);