CREATE TABLE users
(
    id            serial       not null unique,
    username      varchar(255) not null unique,
    password_hash varchar(255) not null
);