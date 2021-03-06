-- +goose Up
SELECT MAX(id) + 1 FROM shopping_lists;
CREATE SEQUENCE shopping_lists_id_sequence START WITH 1;
ALTER TABLE shopping_lists ALTER COLUMN id SET DEFAULT nextval('shopping_lists_id_sequence');
ALTER TABLE shopping_lists ALTER COLUMN id SET NOT NULL;
ALTER SEQUENCE shopping_lists_id_sequence OWNED BY shopping_lists.id;


-- +goose Down
ALTER TABLE shopping_lists ALTER COLUMN id DROP DEFAULT;
DROP SEQUENCE shopping_lists_id_sequence;
