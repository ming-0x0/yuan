-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `categories` (
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `name_vi` TEXT NOT NULL,
    `name_en` TEXT NOT NULL,
    `name_zh` TEXT NOT NULL,
    `description_vi` TEXT,
    `description_en` TEXT,
    `description_zh` TEXT,
    `type` INT NOT NULL,
    `router_vi` TEXT,
    `router_en` TEXT,
    `router_zh` TEXT,
    `position` INT,
    `parent_id` BIGINT,
    `status` INT NOT NULL DEFAULT 1,
    `resource_id` BIGINT,
    `level` INT NOT NULL,
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_parent_type_position` (`type`,`parent_id`,`position`),
    FOREIGN KEY (`parent_id`) REFERENCES `categories` (`id`),
    FOREIGN KEY (`resource_id`) REFERENCES `resources` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "categories";
-- +goose StatementEnd
