-- +goose Up
CREATE TABLE ips (
    id INTEGER PRIMARY KEY,
    address TEXT NOT NULL,
    is_used INTEGER NOT NULL DEFAULT 0,
    subnet_id INTEGER NOT NULL,
    FOREIGN KEY (subnet_id) REFERENCES subnets(id)
);

-- +goose Down
DROP TABLE ips;
