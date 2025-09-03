-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS workout_entries (
    id  BIGSERIAL  PRIMARY KEY,
    workout_id BIGINT NOT NULL,
    exercise_name VARCHAR(255) NOT NULL,
    sets INT NOT NULL,
    reps INT NOT NULL,
    duration_seconds INT,
    weight_kg DECIMAL(5,2),
    notes TEXT,
    order_index INT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    CONSTRAINT valid_workout_entry CHECK (
        (sets > 0 AND reps > 0 AND duration_seconds IS NULL AND weight_kg IS NULL) OR
        (sets > 0 AND reps > 0 AND duration_seconds IS NULL AND weight_kg > 0) OR
        (sets IS NULL AND reps IS NULL AND duration_seconds > 0 AND weight_kg IS NULL)
    )

)
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS workout_entries;
-- +goose StatementEnd

