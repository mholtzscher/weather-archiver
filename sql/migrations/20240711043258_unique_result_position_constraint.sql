-- +goose Up
-- +goose StatementBegin
ALTER TABLE result
ADD constraint unique_position_result UNIQUE (race_id, position);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE result
DROP CONSTRAINT unique_position_result;
-- +goose StatementEnd
