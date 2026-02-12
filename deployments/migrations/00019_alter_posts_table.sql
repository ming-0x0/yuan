-- +goose Up
-- +goose StatementBegin
ALTER TABLE `posts`
ADD COLUMN `color_palette` TEXT AFTER `flagship`;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE `posts`
DROP COLUMN `color_palette`;
-- +goose StatementEnd
