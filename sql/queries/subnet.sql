-- name: GetSubnets :many
SELECT * FROM subnet;

-- name: NewSubnet :exec
INSERT INTO subnet (name, network_prefix, cidr) VALUES (?, ?, ?);

-- name: GetSubnet :one
SELECT * FROM subnet WHERE id = ?;

-- name: GetSubnetByName :one
SELECT * FROM subnet WHERE name = ?;
