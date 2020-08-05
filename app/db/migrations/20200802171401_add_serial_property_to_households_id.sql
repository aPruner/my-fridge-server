-- +goose Up
SELECT MAX(id) + 1 FROM households;
CREATE SEQUENCE household_id_sequence START WITH 1;
ALTER TABLE households ALTER COLUMN id SET DEFAULT nextval('household_id_sequence');
ALTER TABLE households ALTER COLUMN id SET NOT NULL;
ALTER SEQUENCE household_id_sequence OWNED BY households.id;


-- +goose Down
ALTER TABLE households ALTER COLUMN id DROP DEFAULT;
DROP SEQUENCE household_id_sequence;
