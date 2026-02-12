-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `user_tokens` (
    `id` BIGINT AUTO_INCREMENT,
    `user_id` BIGINT NOT NULL,
    `token` TEXT NOT NULL,
    `token_id` TEXT NOT NULL,
    `expired_at` TIMESTAMP NOT NULL,
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    CONSTRAINT `fk_user` FOREIGN KEY (`user_id`) REFERENCES `users`(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "user_tokens";
-- +goose StatementEnd
