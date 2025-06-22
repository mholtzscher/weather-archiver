-- +goose Up
-- +goose StatementBegin
ALTER TABLE race
RENAME COLUMN race_date TO date;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE race
RENAME COLUMN date TO race_date;
-- +goose StatementEnd
