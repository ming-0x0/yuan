-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
	id BIGINT PRIMARY KEY,
	email VARCHAR(265) NOT NULL UNIQUE,
	username VARCHAR(265) NOT NULL UNIQUE,
	password VARCHAR(100) NOT NULL,
	created_by BIGINT NOT NULL,
	created_at TIMESTAMP(0) DEFAULT CURRENT_TIMESTAMP,
	updated_by BIGINT NOT NULL,
	updated_at TIMESTAMP(0) DEFAULT CURRENT_TIMESTAMP,
	deleted_by BIGINT DEFAULT NULL,
	deleted_at TIMESTAMP(0) DEFAULT NULL
);
CREATE UNIQUE INDEX idx_users_email ON users(email) WHERE deleted_at IS NULL;
CREATE UNIQUE INDEX idx_users_username ON users(username) WHERE deleted_at IS NULL;
CREATE INDEX idx_users_deleted_at ON users(deleted_at);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
