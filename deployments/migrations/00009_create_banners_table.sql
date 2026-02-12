-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `banners` (
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `name_vi` TEXT NOT NULL,
    `name_en` TEXT NOT NULL,
    `name_zh` TEXT NOT NULL,
    `description_vi` TEXT,
    `description_en` TEXT,
    `description_zh` TEXT,
    `position` INT UNIQUE,
    `status` INT NOT NULL DEFAULT 1,
    `resource_id` BIGINT NOT NULL,
    `link` TEXT,
    `button_name_vi` TEXT,
    `button_name_en` TEXT,
    `button_name_zh` TEXT,
    `has_content` TINYINT(1) NOT NULL DEFAULT 0,
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`resource_id`) REFERENCES `resources` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `banners`;
-- +goose StatementEnd
