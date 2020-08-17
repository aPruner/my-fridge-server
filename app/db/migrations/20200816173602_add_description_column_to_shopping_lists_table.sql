-- +goose Up
ALTER TABLE shopping_lists ADD COLUMN description VARCHAR(100);

-- +goose Down
ALTER TABLE shopping_lists DROP COLUMN description;
