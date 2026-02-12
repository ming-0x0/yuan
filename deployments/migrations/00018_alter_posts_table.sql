-- +goose Up
-- +goose StatementBegin
ALTER TABLE `posts`
ADD COLUMN `flagship` TINYINT(1) NOT NULL DEFAULT 0 AFTER `type`;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE `posts`
DROP COLUMN `flagship`;
-- +goose StatementEnd
