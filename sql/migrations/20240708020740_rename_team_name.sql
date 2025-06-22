-- +goose Up
-- +goose StatementBegin
ALTER TABLE team
RENAME COLUMN team_name TO name;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE team
RENAME COLUMN name TO team_name;
-- +goose StatementEnd
