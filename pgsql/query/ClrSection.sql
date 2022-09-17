-- name: GetClrSectionByID :one
SELECT * FROM assets.clr_sections
WHERE id = $1;

-- name: GetClrSectionBySerial :one
SELECT * FROM assets.clr_sections
WHERE serial_number = $1;

-- name: GetClrSections :many
SELECT * FROM assets.clr_sections
ORDER BY make;

-- name: CreateClrSection :one
INSERT INTO assets.clr_sections
(make, serial_number, mawp, num_tubes, num_rows, passes)
VALUES
($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: AssignClrSection :exec
UPDATE assets.clr_sections
SET opstatus = "active", cooler_id = $2
WHERE id = $1;

-- name: DeactivateClrSection :one
UPDATE assets.clr_sections
SET opstatus = "inactive", cooler_id = NULL
WHERE id = $1
RETURNING *;

-- name: DisposeClrSection :one
UPDATE assets.clr_sections
SET opstatus = "disposed", cooler_id = NULL
WHERE id = $1
RETURNING *;

-- name: DeleteClrSection :exec
DELETE FROM assets.clr_sections
WHERE id = $1;