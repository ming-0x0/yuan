-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `posts` (
  `id` BIGINT NOT NULL AUTO_INCREMENT,
  `title_vi` TEXT NOT NULL,
  `title_en` TEXT NOT NULL,
  `title_zh` TEXT NOT NULL,
  `slug_vi` TEXT NOT NULL,
  `slug_en` TEXT NOT NULL,
  `slug_zh` TEXT NOT NULL,
  `alt_vi` TEXT NOT NULL,
  `alt_en` TEXT NOT NULL,
  `alt_zh` TEXT NOT NULL,
  `description_vi` TEXT,
  `description_en` TEXT,
  `description_zh` TEXT,
  `info_vi` TEXT,
  `info_en` TEXT,
  `info_zh` TEXT,
  `resource_ids` TEXT,
  `content_vi` TEXT NOT NULL,
  `content_en` TEXT NOT NULL,
  `content_zh` TEXT NOT NULL,
  `status` INT NOT NULL,
  `type` INT NOT NULL DEFAULT 1,
  `category_id` BIGINT NOT NULL,
  `public_date` DATETIME,
  `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`category_id`) REFERENCES `categories`(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "posts";
-- +goose StatementEnd
