-- +goose Up
CREATE TABLE users(
    id INT PRIMARY KEY,
    username VARCHAR(50)
);

-- +goose Down
DROP TABLE users;