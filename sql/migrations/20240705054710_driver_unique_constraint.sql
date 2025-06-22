-- +goose Up
-- +goose StatementBegin
ALTER TABLE driver
ADD constraint unique_driver UNIQUE (first_name, last_name);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE driver
DROP CONSTRAINT unique_driver;
-- +goose StatementEnd
