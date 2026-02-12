-- +goose Up
-- +goose StatementBegin
ALTER TABLE `users`
ADD COLUMN `receive_email` TINYINT(1) NOT NULL DEFAULT 0 AFTER `is_admin`;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE `users`
DROP COLUMN `receive_email`;
-- +goose StatementEnd
