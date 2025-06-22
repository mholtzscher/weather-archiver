-- name: GetSeasonById :one
SELECT * FROM season
WHERE id = $1 LIMIT 1;

-- name: CreateSeason :one
INSERT INTO season 
(season_year, series)
VALUES (
$1, $2
)
RETURNING id;

-- name: GetAllSeasons :many
SELECT * FROM season
ORDER BY season_year DESC;

