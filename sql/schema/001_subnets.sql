-- +goose Up
CREATE TABLE subnets (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    network_prefix TEXT UNIQUE NOT NULL,
    cidr INTEGER NOT NULL
);

-- +goose Down
DROP TABLE subnets;
