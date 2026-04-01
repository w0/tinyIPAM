CREATE TABLE subnet (
    id INTEGER PRIMARY KEY,
    name TEXT,
    network_prefix TEXT NOT NULL,
    cidr TEXT NOT NULL
);
