-- +migrate Up
create table users
(
    id           int primary key AUTO_INCREMENT,
    name         varchar(191) not null,
    phone_number varchar(191) not null unique,
    created_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +migrate Down
DROP TABLE users;