-- +goose Up
-- +goose StatementBegin
CREATE TABLE driver (
    id SERIAL PRIMARY KEY,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    place_of_birth TEXT NOT NULL,
    date_of_birth DATE NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE driver;
-- +goose StatementEnd
