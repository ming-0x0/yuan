-- +goose Up
-- +goose StatementBegin
INSERT INTO `permission_groups` (`id`, `name`, `description`, `full_permission`)
VALUES (1, "Admin", "admin", 1);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
TRUNCATE TABLE `permission_groups`;
-- +goose StatementEnd
