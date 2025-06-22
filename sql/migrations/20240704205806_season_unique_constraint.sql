-- +goose Up
-- +goose StatementBegin
ALTER TABLE season
ADD constraint unique_season UNIQUE (season_year, series);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE season
DROP CONSTRAINT unique_season;
-- +goose StatementEnd
