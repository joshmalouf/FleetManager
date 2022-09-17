-- name: GetCylinderByID :one
SELECT * FROM assets.Cylinders
WHERE id = $1;

-- name: GetCylinderBySerial :one
SELECT * FROM assets.Cylinders
WHERE serial_number = $1;

-- name: GetCylinders :many
SELECT * FROM assets.Cylinders
ORDER BY make;

-- name: GetCylindersByMakeModel :many
SELECT * FROM assets.Cylinders
WHERE make = $1 AND model = $2
ORDER BY bore;

-- name: CreateCylinder :one
INSERT INTO assets.Cylinders
(make, model, mawp, serial_number)
VALUES
($1, $2, $3, $4)
RETURNING *;

-- name: AssignCylinder :exec
UPDATE assets.Cylinders
SET opstatus = "active", comp_id = $2
WHERE id = $1;

-- name: DeactivateCylinder :one
UPDATE assets.Cylinders
SET opstatus = "inactive", comp_id = NULL
WHERE id = $1
RETURNING *;

-- name: DisposeCylinder :one
UPDATE assets.Cylinders
SET opstatus = "disposed", comp_id = NULL
WHERE id = $1
RETURNING *;

-- name: DeleteCylinder :exec
DELETE FROM assets.Cylinders
WHERE id = $1;