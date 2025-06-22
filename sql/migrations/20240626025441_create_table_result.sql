-- +goose Up
-- +goose StatementBegin
CREATE TABLE result (
    id SERIAL PRIMARY KEY,
    race_id INT NOT NULL,
    driver_id INT NOT NULL,
    team_id INT NOT NULL,
    position INT NOT NULL,
    points FLOAT NOT NULL,
    FOREIGN KEY (race_id) REFERENCES race (id),
    FOREIGN KEY (driver_id) REFERENCES driver (id),
    FOREIGN KEY (team_id) REFERENCES team (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE result;
-- +goose StatementEnd
