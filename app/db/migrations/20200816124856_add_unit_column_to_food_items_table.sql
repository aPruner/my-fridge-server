-- +goose Up
ALTER TABLE food_items ADD COLUMN unit VARCHAR(50);

-- +goose Down
ALTER TABLE food_items DROP COLUMN unit;
