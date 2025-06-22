-- +goose Up
-- +goose StatementBegin
ALTER TABLE result
ADD constraint unique_driver_result UNIQUE (race_id, driver_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE result
DROP CONSTRAINT unique_driver_result;
-- +goose StatementEnd
