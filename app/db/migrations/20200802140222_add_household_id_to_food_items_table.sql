-- +goose Up
ALTER TABLE foodItems ADD COLUMN household_id INT REFERENCES households(id);

-- +goose Down
ALTER TABLE foodItems DROP COLUMN household_id;
