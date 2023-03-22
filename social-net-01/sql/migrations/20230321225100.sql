use `social-net`
;

alter table `users`
    change column password_hash password_hash varchar(255)
;