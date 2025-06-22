-- name: GetTeamById :one
SELECT * FROM team
WHERE id = $1 LIMIT 1;

-- name: CreateTeam :one
INSERT INTO team 
(name, base)
VALUES (
$1, $2
)
RETURNING id;

