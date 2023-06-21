USE `social-net`
;

CREATE TABLE IF NOT EXISTS `users`
(
    id            varchar(36) PRIMARY KEY,
    age           int8,
    biography     text,
    city          varchar(50),
    first_name    varchar(50),
    last_name     varchar(50),
    password_hash varchar(50)
)
;
