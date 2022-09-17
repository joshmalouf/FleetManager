-- name: GetDriverById :one
SELECT * FROM assets.Drivers
WHERE id = $1;

-- name: GetDriverByUnitID :one
SELECT * FROM assets.drivers
WHERE unit_id = $1;

-- name: CreateEngineDriver :exec
INSERT INTO assets.drivers
(engine_id, unit_id)
VALUES
($1, $2);

-- name: CreateMotorDriver :exec
INSERT INTO assets.drivers
(motor_id, unit_id)
VALUES
($1, $2);

-- name: AssignEngine :exec
UPDATE assets.drivers
SET unit_id =$2
WHERE id = $1;

-- name: DeletDriver :exec
DELETE FROM assets.drivers
WHERE id = $1;