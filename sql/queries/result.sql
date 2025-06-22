-- name: GetResultById :one
SELECT * FROM result
WHERE id = $1 LIMIT 1;

-- name: CreateResult :one
INSERT INTO result 
(race_id, driver_id, team_id, position, points)
VALUES (
$1, $2, $3, $4, $5
)
RETURNING id;

-- name: GetResultsByRaceId :many
SELECT * FROM result
WHERE race_id = $1
ORDER BY position ASC;
