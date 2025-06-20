-- +goose Up
CREATE TABLE IF NOT EXISTS api_keys (
    key TEXT PRIMARY KEY,
    author_id TEXT NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE IF EXISTS api_keys;
