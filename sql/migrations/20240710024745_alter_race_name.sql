-- +goose Up
-- +goose StatementBegin
ALTER TABLE race
RENAME COLUMN race_name TO name;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE race
RENAME COLUMN name TO race_name;
-- +goose StatementEnd
