-- +goose Up
-- +goose StatementBegin
ALTER TABLE team
ADD constraint unique_team UNIQUE (name);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE team
DROP CONSTRAINT unique_team;
-- +goose StatementEnd
