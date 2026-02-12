-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `customers` (
  `id` BIGINT NOT NULL AUTO_INCREMENT,
  `customer_name` TEXT NOT NULL,
  `email` TEXT NOT NULL,
  `phone_number` TEXT NOT NULL,
  `message` TEXT,
  `note` TEXT,
  `service_type` INT NOT NULL,
  `status` INT NOT NULL DEFAULT 2,
  `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "customers";
-- +goose StatementEnd
