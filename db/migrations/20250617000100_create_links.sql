-- +goose Up
CREATE TABLE IF NOT EXISTS links (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT,
    url TEXT NOT NULL UNIQUE,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS links;