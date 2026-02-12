-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `users` (
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `email` VARCHAR(512) NOT NULL UNIQUE,
    `username` VARCHAR(50) NOT NULL UNIQUE,
    `full_name` TEXT NOT NULL,
    `password` TEXT NOT NULL,
    `status` INTEGER NOT NULL DEFAULT 1,
    `is_admin` TINYINT(1) NOT NULL DEFAULT 0,
    `permission_group_id` BIGINT NOT NULL,
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`permission_group_id`) REFERENCES `permission_groups` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "users";
-- +goose StatementEnd
