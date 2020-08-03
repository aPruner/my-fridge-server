-- +goose Up
CREATE TABLE foodItems(
    id INT PRIMARY KEY,
    name VARCHAR(50),
    category VARCHAR(50),
    amount INT
);

-- +goose Down
DROP TABLE foodItems;
