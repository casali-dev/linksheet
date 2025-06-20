-- +goose Up
CREATE TABLE IF NOT EXISTS links (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT,
    url TEXT NOT NULL,
    is_public INTEGER NOT NULL DEFAULT 1,
    author_id TEXT NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL,
    UNIQUE (url, author_id)
);

-- +goose Down
DROP TABLE IF EXISTS links;
