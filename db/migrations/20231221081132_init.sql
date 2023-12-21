-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

CREATE TABLE todo (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    content TEXT NOT NULL,
    is_finished BOOLEAN DEFAULT false,
    created_at DATE DEFAULT CURRENT_TIMESTAMP
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

DROP TABLE todo

-- +goose StatementEnd
