USE `social-net`
;

ALTER TABLE `users`
    CHANGE COLUMN password_hash password_hash varchar(255)
;
