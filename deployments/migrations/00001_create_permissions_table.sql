-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `permissions` (
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `permission_name` VARCHAR(255) NOT NULL UNIQUE,
    `function_code` VARCHAR(255) NOT NULL,
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "permissions";
-- +goose StatementEnd
