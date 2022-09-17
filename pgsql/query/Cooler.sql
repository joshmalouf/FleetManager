-- name: GetCoolerByID :one
SELECT * FROM assets.Coolers
WHERE id = $1;

-- name: GetCoolerByJob :one
SELECT * FROM assets.Coolers
WHERE job_number = $1;

-- name: GetCoolers :many
SELECT * FROM assets.Coolers
ORDER BY make;

-- name: GetCoolersByMake :many
SELECT * FROM assets.Coolers
WHERE make = $1
ORDER BY model;

-- name: GetCoolerByMakeModel :many
SELECT * FROM assets.Coolers
WHERE make = $1 AND model = $2
ORDER BY size;

-- name: GetAvailCoolers :many
SELECT * FROM assets.Coolers
WHERE opstatus = "active" AND unit_id = NULL;

-- name: CreateCooler :one
INSERT INTO assets.coolers
(make, model, size, job_number)
VALUES
($1, $2, $3, $4)
RETURNING *;

-- name: DeactivateCooler :exec
UPDATE assets.Coolers
SET opstatus = "inactive", unit_id = NULL
WHERE id = $1;

-- name: AssignCooler :exec
UPDATE assets.Coolers
SET unit_id = $2
WHERE id = $1;