-- name: GetCompressorByID :one
SELECT * FROM assets.Compressors
WHERE id = $1;

-- name: GetCompressorBySerial :one
SELECT * FROM assets.Compressors
WHERE serial_number = $1;

-- name: GetCompressors :many
SELECT * FROM assets.Compressors
ORDER BY make;

-- name: GetCompressorsByMake :many
SELECT * FROM assets.Compressors
WHERE make = $1
ORDER BY model;

-- name: GetCompressorByMakeModel :many
SELECT * FROM assets.Compressors
WHERE make = $1 AND model = $2
ORDER BY throws;

-- name: GetAvailCompressors :many
SELECT * FROM assets.Compressors
WHERE opstatus = "active" AND unit_id = NULL;

-- name: CreateCompressor :one
INSERT INTO assets.Compressors
(make, model, serial_number, throws)
VALUES
($1, $2, $3, $4)
RETURNING *;

-- name: DeactivateCompressor :exec
UPDATE assets.Compressors
SET opstatus = "inactive", unit_id = NULL
WHERE id = $1;

-- name: DisposeCompressor :exec
UPDATE assets.Compressors
Set op_status = "disposed", unit_id = null;

-- name: AssignCompressor :exec
UPDATE assets.Compressors
SET unit_id = $2
WHERE id = $1;

