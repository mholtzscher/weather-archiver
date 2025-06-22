-- +goose Up
-- +goose StatementBegin
CREATE TABLE sensors (
    id serial PRIMARY KEY,
	  type varchar(50),
	  location varchar(50)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE sensors;
-- +goose StatementEnd
