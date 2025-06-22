-- name: GetDriverById :one
SELECT * FROM driver
WHERE id = $1 LIMIT 1;

-- name: CreateDriver :one
INSERT INTO driver
(first_name, last_name, place_of_birth, date_of_birth)
VALUES (
$1, $2, $3, $4
)
RETURNING id;

