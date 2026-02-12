-- +goose Up
-- +goose StatementBegin
INSERT INTO `users` (`id`, `email`, `username`, `full_name`, `password`, `is_admin`, `permission_group_id`)
VALUES (1, "admin@gmail.com", "admin", "admin", "$2a$10$GF9i2mdKQf2wVMRVSYfRbugg609XvtjrdTZAWt7yIRzX1PDGXXylS", 1, 1);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
TRUNCATE TABLE `users`;
-- +goose StatementEnd
