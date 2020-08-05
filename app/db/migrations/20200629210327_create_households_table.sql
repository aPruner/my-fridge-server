-- +goose Up
CREATE TABLE households(
    id INT PRIMARY KEY,
    name VARCHAR(50),
    city VARCHAR(50)
);

-- +goose Down
DROP TABLE households;
