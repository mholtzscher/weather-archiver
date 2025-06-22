-- +goose Up
-- +goose StatementBegin
ALTER TABLE race
ADD constraint unique_race UNIQUE (season_id, name);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE race
DROP CONSTRAINT unique_race;
-- +goose StatementEnd
