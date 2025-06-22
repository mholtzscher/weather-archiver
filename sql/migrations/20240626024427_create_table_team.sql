-- +goose Up
-- +goose StatementBegin
CREATE TABLE team (
    id SERIAL PRIMARY KEY,
    team_name TEXT NOT NULL,
    base TEXT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE team;
-- +goose StatementEnd
