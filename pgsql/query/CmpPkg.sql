-- name: GetCmpPkgByID :one
SELECT * FROM assets.CmpPkgs
WHERE id = $1;

-- name: GetCmpPkgByUnitNumber :one
SELECT * FROM assets.CmpPkgs
WHERE unit_number = $1;

-- name: GetCmpPkgs :many
SELECT * FROM assets.CmpPkgs
ORDER BY unit_number;

-- name: GetCmpPkgsByStages :many
SELECT * FROM assets.CmpPkgs
WHERE stages = $1;

-- name: GetCmpPkgsByEngine :many
SELECT pkgs.unit_number, eng.make, eng.model
FROM assets.CmpPkgs pkgs
INNER JOIN assets.drivers dvr ON dvr.id = pkgs.driver_id
INNER JOIN assets.engines eng ON eng.id = dvr.id
WHERE eng.make = $1 AND eng.model = $2;

-- name: CreateCmpPkg :one
INSERT INTO assets.CmpPkgs 
(unit_number,stages, drawing_ref)
VALUES
($1,$2, $3)
RETURNING *;

-- name: CmpPkgChgDriver :one
UPDATE assets.CmpPkgs 
SET driver_id = $2
WHERE id = $1
RETURNING *;

-- name: CmpPkgChgComp :one
UPDATE assets.CmpPkgs
SET compressor_id = $2
WHERE id = $1
RETURNING *;

-- name: DeactivateCmpPkg :one
UPDATE assets.CmpPkgs
SET op_status = "inactive"
WHERE id = $1
RETURNING *;

-- name: DisposeCmpPkg :one
UPDATE assets.CmpPkgs
SET op_status = "disposed"
WHERE id = $1
RETURNING *;

-- name: DeleteCmpPkg :exec
DELETE FROM assets.CmpPkgs
WHERE id = $1;
