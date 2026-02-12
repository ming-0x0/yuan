-- +goose Up
-- +goose StatementBegin
ALTER TABLE `posts`
ADD COLUMN `avatar` BIGINT NOT NULL AFTER `description_zh`,
ADD CONSTRAINT `fk_posts_avatar`
FOREIGN KEY (`avatar`) REFERENCES `resources` (`id`);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE `posts`
DROP FOREIGN KEY `fk_posts_avatar`,
DROP COLUMN `avatar`;
-- +goose StatementEnd
