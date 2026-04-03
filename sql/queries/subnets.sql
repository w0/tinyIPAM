-- name: GetSubnets :many
SELECT * FROM subnets;

-- name: NewSubnet :exec
INSERT INTO subnets (name, network_prefix, cidr) VALUES (?, ?, ?);

-- name: GetSubnet :one
SELECT * FROM subnets WHERE id = ?;

-- name: GetSubnetByName :one
SELECT * FROM subnets WHERE name = ?;
