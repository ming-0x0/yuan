-- +goose Up
-- +goose StatementBegin
CREATE TABLE `footers` (
  `id` BIGINT PRIMARY KEY AUTO_INCREMENT NOT NULL,
  `name_vi` TEXT NOT NULL,
  `name_en` TEXT NOT NULL,
  `name_zh` TEXT NOT NULL,
  `content_vi` TEXT,
  `content_en` TEXT,
  `content_zh` TEXT,
  `link` TEXT,
  `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `footers`;
-- +goose StatementEnd
