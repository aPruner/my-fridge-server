-- +goose Up
ALTER TABLE food_items ALTER COLUMN household_id SET NOT NULL;

-- +goose Down
ALTER TABLE food_items ALTER COLUMN household_id DROP NOT NULL;
