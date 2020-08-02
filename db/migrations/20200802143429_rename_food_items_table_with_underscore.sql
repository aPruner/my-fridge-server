-- +goose Up
ALTER TABLE fooditems RENAME to food_items;

-- +goose Down
ALTER TABLE food_items RENAME to fooditems;
