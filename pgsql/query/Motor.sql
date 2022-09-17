-- name: GetMotorByID :one
SELECT * FROM assets.Motors
WHERE id = $1;

-- name: GetMotorBySerial :one
SELECT * FROM assets.Motors
WHERE serial_number = $1;

-- name: GetMotors :many
SELECT * FROM assets.Motors
ORDER BY make;

-- name: GetMotorsByMake :many
SELECT * FROM assets.Motors
WHERE make = $1
ORDER BY model;

-- name: GetMotorByMakeModel :many
SELECT * FROM assets.Motors
WHERE make = $1 AND model = $2
ORDER BY throws;

-- name: GetAvailMotors :many
SELECT * FROM assets.Motors
WHERE opstatus = "active" AND unit_id = NULL;

-- name: DeactivateMotor :exec
UPDATE assets.Motors
SET opstatus = "inactive", unit_id = NULL
WHERE id = $1;

-- name: AssignMotor :exec
UPDATE assets.Motors
SET unit_id = $2
WHERE id = $1;