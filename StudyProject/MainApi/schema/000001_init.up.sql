CREATE TABLE users (
    id serial not null unique,
    name varchar(100) not null,
    username varchar(255) not null unique,
    password varchar(255) not null
)