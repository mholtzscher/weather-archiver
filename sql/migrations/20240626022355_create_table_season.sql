-- +goose Up
-- +goose StatementBegin
CREATE TABLE season (
    id SERIAL PRIMARY KEY,
    season_year INT NOT NULL,
    series TEXT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE season;
-- +goose StatementEnd
