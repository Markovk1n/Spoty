CREATE TABLE users
(
    id            serial       not null unique,
    username      varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE comments
(
    id            serial       not null unique,
    comment      varchar,
    user_id int
);

CREATE TABLE track_comments
(
    id            serial       not null unique,
    track_id      varchar,
    comment_id int
);

CREATE TABLE album_comments
(
    id            serial       not null unique,
    album_id      varchar,
    comment_id int
);
