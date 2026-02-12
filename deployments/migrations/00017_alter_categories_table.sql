-- +goose Up
-- +goose StatementBegin
ALTER TABLE `categories`
ADD COLUMN `editable` TINYINT(1) NOT NULL DEFAULT 1 AFTER `type`;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE `categories`
DROP COLUMN `editable`;
-- +goose StatementEnd
