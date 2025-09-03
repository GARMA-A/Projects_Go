-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS workouts (
    id  BIGSERIAL  PRIMARY KEY,
    -- user_id BIGINT REFERENCES users(id) ON DELETE CASCADE,
    title VARCHAR(100) NOT NULL,
    description TEXT,
    duration_minutes INT NOT NULL, -- duration in minutes
    colorized_burned INTEGER,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
)
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS workouts;
-- +goose StatementEnd

