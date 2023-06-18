USE `social-net`
;

DELIMITER $$

DROP PROCEDURE IF EXISTS migration_20230613212639;
CREATE PROCEDURE migration_20230613212639()
BEGIN
    DECLARE EXIT HANDLER FOR SQLEXCEPTION
        BEGIN
            SHOW ERRORS;
            ROLLBACK;
        END;

    IF NOT EXISTS (SELECT *
                   FROM information_schema.columns
                   WHERE table_schema = 'social-net'
                     AND table_name = 'users'
                     AND column_name = 'second_name')
    THEN
        START TRANSACTION;

        ALTER TABLE `users`
            CHANGE COLUMN `last_name` `second_name` VARCHAR(50);

        SELECT 'migrating';
        COMMIT;
    END IF;

END$$

DELIMITER ;

CALL migration_20230613212639();
DROP PROCEDURE IF EXISTS migration_20230613212639;
