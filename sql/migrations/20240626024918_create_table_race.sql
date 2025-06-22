-- +goose Up
-- +goose StatementBegin
CREATE TABLE race (
    id SERIAL PRIMARY KEY,
    season_id INT NOT NULL,
    race_name TEXT NOT NULL,
    location TEXT NOT NULL,
    race_date DATE NOT NULL,
    FOREIGN KEY (season_id) REFERENCES season (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE race;
-- +goose StatementEnd
