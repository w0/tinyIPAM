-- name: GetFreeIP :one
SELECT * FROM ips
WHERE subnet_id = :subnet_id
AND is_used = FALSE
ORDER BY address ASC
LIMIT 1;

-- name: ReserveIP :exec
UPDATE ips
SET is_used = 1
WHERE id = :id;

-- name: ReleaseIP :exec
UPDATE ips
SET is_used = 0
WHERE id = :id;

-- name: GetAllAssigned :many
SELECT * FROM ips
WHERE subnet_id = :subnet_id
AND is_used = TRUE
ORDER BY address ASC;
