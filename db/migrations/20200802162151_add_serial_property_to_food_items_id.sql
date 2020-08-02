-- +goose Up
SELECT MAX(id) + 1 FROM food_items;
CREATE SEQUENCE id_sequence START WITH 1;
ALTER TABLE food_items ALTER COLUMN id SET DEFAULT nextval('id_sequence');
ALTER TABLE food_items ALTER COLUMN id SET NOT NULL;
ALTER SEQUENCE id_sequence OWNED BY food_items.id;


-- +goose Down
ALTER TABLE food_items ALTER COLUMN id DROP DEFAULT;
ALTER TABLE food_items ALTER COLUMN id DROP NOT NULL;
DROP SEQUENCE id_sequence;
