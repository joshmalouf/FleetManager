-- name: GetEngineByID :one
SELECT * FROM assets.Engines
WHERE id = $1;

-- name: GetEngineBySerial :one
SELECT * FROM assets.Engines
WHERE serial_number = $1;

-- name: GetEngines :many
SELECT * FROM assets.Engines
ORDER BY make;

-- name: GetEnginesByMake :many
SELECT * FROM assets.Engines
WHERE make = $1
ORDER BY model;

-- name: GetEngineByMakeModel :many
SELECT * FROM assets.Engines
WHERE make = $1 AND model = $2
ORDER BY throws;

-- name: GetAvailEngines :many
SELECT * FROM assets.Engines
WHERE opstatus = "active" AND unit_id = NULL;

-- name: DeactivateEngine :exec
UPDATE assets.Engines
SET opstatus = "inactive", unit_id = NULL
WHERE id = $1;
