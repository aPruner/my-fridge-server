-- +goose Up
SELECT MAX(id) + 1 FROM food_items;
CREATE SEQUENCE food_item_id_sequence START WITH 1;
ALTER TABLE food_items ALTER COLUMN id SET DEFAULT nextval('food_item_id_sequence');
ALTER TABLE food_items ALTER COLUMN id SET NOT NULL;
ALTER SEQUENCE food_item_id_sequence OWNED BY food_items.id;


-- +goose Down
ALTER TABLE food_items ALTER COLUMN id DROP DEFAULT;
DROP SEQUENCE food_item_id_sequence;
