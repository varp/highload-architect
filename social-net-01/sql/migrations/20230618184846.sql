USE `social-net`
;

CREATE TABLE IF NOT EXISTS `auth`
(
    `api_key`          varchar(50) PRIMARY KEY,
    `user_id`          varchar(50),
    `authenticated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    INDEX `auth_user_id_fk` (`user_id`),
    FOREIGN KEY (`user_id`)
        REFERENCES `users` (id)
        ON DELETE CASCADE
)

;
