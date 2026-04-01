-- name: GetFreeIP :one
SELECT * FROM ip
WHERE subnet_id = :subnet_id
AND is_used = FALSE
ORDER BY address ASC
LIMIT 1;

-- name: ReserveIP :exec
UPDATE ip
SET is_used = 1
WHERE id = :id;

-- name: ReleaseIP :exec
UPDATE ip
SET is_used = 0
WHERE id = :id;

-- name: GetAllAssigned :many
SELECT * FROM ip
WHERE subnet_id = :subnet_id
AND is_used = TRUE
ORDER BY address ASC;
