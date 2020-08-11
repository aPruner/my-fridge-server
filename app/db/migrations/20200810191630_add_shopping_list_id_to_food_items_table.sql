-- +goose Up
ALTER TABLE food_items ADD COLUMN shopping_list_id INT REFERENCES shopping_lists(id);

-- +goose Down
ALTER TABLE food_items DROP COLUMN shopping_list_id;
