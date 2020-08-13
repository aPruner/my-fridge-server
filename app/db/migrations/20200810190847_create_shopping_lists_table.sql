-- +goose Up
CREATE TABLE shopping_lists(
    id INT PRIMARY KEY,
    user_id INT REFERENCES users(id),
    household_id INT REFERENCES households(id),
    name VARCHAR(50),
    created_at DATE
);

-- +goose Down
DROP TABLE shopping_lists;
