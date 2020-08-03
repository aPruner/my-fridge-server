-- +goose Up
SELECT MAX(id) + 1 FROM users;
CREATE SEQUENCE user_id_sequence START WITH 1;
ALTER TABLE users ALTER COLUMN id SET DEFAULT nextval('user_id_sequence');
ALTER TABLE users ALTER COLUMN id SET NOT NULL;
ALTER SEQUENCE user_id_sequence OWNED BY users.id;


-- +goose Down
ALTER TABLE users ALTER COLUMN id DROP DEFAULT;
DROP SEQUENCE user_id_sequence;
