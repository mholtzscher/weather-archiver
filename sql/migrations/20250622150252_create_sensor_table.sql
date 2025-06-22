-- +goose Up
-- +goose StatementBegin
CREATE TABLE sensors (
    id serial PRIMARY KEY,
    source text NOT NULL,
    created_at timestamptz DEFAULT now() NOT NULL
);

CREATE TABLE sensor_data (
    time TIMESTAMPTZ NOT NULL,
    sensor_id INT NOT NULL,
    -- sensor_type_id INT NOT NULL,
    numeric_value DOUBLE PRECISION,
    metadata JSONB,
    FOREIGN KEY (sensor_id) REFERENCES sensors(id)
    -- FOREIGN KEY (sensor_type_id) REFERENCES sensor_types(id)
);

SELECT create_hypertable('sensor_data', 'time');

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE sensor_data;
DROP TABLE sensors;
-- +goose StatementEnd
