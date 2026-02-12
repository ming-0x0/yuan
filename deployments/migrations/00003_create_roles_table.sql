-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `roles` (
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `permission_id` BIGINT NOT NULL,
    `permission_group_id` BIGINT NOT NULL,
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`permission_id`) REFERENCES `permissions` (`id`),
    FOREIGN KEY (`permission_group_id`) REFERENCES `permission_groups` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "roles";
-- +goose StatementEnd
