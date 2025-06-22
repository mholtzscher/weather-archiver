-- name: GetRaceById :one
SELECT * FROM race
WHERE id = $1 LIMIT 1;

-- name: CreateRace :one
INSERT INTO race
(season_id, name, location, date)
VALUES (
$1, $2, $3, $4
)
RETURNING id;


