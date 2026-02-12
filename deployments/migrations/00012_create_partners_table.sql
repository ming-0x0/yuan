-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `partners` (
  `id` BIGINT PRIMARY KEY AUTO_INCREMENT NOT NULL,
  `name` TEXT NOT NULL,
  `description_vi` TEXT,
  `description_en` TEXT,
  `description_zh` TEXT,
  `resource_id` BIGINT NOT NULL,
  `status` INT NOT NULL,
  `position` INT UNIQUE,
  `link` TEXT,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  FOREIGN KEY (`resource_id`) REFERENCES `resources` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `partners`;
-- +goose StatementEnd
