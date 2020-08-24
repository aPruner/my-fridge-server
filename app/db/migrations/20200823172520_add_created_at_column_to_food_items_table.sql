-- +goose Up
ALTER TABLE food_items ADD COLUMN created_at DATE;

-- +goose Down
ALTER TABLE food_items DROP COLUMN created_at;
