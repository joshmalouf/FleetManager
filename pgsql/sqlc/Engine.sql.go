// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: Engine.sql

package pgsql

import (
	"context"
)

const deactivateEngine = `-- name: DeactivateEngine :exec
UPDATE assets.Engines
SET opstatus = "inactive", unit_id = NULL
WHERE id = $1
`

func (q *Queries) DeactivateEngine(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deactivateEngine, id)
	return err
}

const getAvailEngines = `-- name: GetAvailEngines :many
SELECT id, make, model, serial_number, unit_id, op_status, created_at, modified_at FROM assets.Engines
WHERE opstatus = "active" AND unit_id = NULL
`

func (q *Queries) GetAvailEngines(ctx context.Context) ([]AssetsEngine, error) {
	rows, err := q.db.QueryContext(ctx, getAvailEngines)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []AssetsEngine{}
	for rows.Next() {
		var i AssetsEngine
		if err := rows.Scan(
			&i.ID,
			&i.Make,
			&i.Model,
			&i.SerialNumber,
			&i.UnitID,
			&i.OpStatus,
			&i.CreatedAt,
			&i.ModifiedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getEngineByID = `-- name: GetEngineByID :one
SELECT id, make, model, serial_number, unit_id, op_status, created_at, modified_at FROM assets.Engines
WHERE id = $1
`

func (q *Queries) GetEngineByID(ctx context.Context, id int64) (AssetsEngine, error) {
	row := q.db.QueryRowContext(ctx, getEngineByID, id)
	var i AssetsEngine
	err := row.Scan(
		&i.ID,
		&i.Make,
		&i.Model,
		&i.SerialNumber,
		&i.UnitID,
		&i.OpStatus,
		&i.CreatedAt,
		&i.ModifiedAt,
	)
	return i, err
}

const getEngineByMakeModel = `-- name: GetEngineByMakeModel :many
SELECT id, make, model, serial_number, unit_id, op_status, created_at, modified_at FROM assets.Engines
WHERE make = $1 AND model = $2
ORDER BY throws
`

type GetEngineByMakeModelParams struct {
	Make  string `json:"make"`
	Model string `json:"model"`
}

func (q *Queries) GetEngineByMakeModel(ctx context.Context, arg GetEngineByMakeModelParams) ([]AssetsEngine, error) {
	rows, err := q.db.QueryContext(ctx, getEngineByMakeModel, arg.Make, arg.Model)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []AssetsEngine{}
	for rows.Next() {
		var i AssetsEngine
		if err := rows.Scan(
			&i.ID,
			&i.Make,
			&i.Model,
			&i.SerialNumber,
			&i.UnitID,
			&i.OpStatus,
			&i.CreatedAt,
			&i.ModifiedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getEngineBySerial = `-- name: GetEngineBySerial :one
SELECT id, make, model, serial_number, unit_id, op_status, created_at, modified_at FROM assets.Engines
WHERE serial_number = $1
`

func (q *Queries) GetEngineBySerial(ctx context.Context, serialNumber string) (AssetsEngine, error) {
	row := q.db.QueryRowContext(ctx, getEngineBySerial, serialNumber)
	var i AssetsEngine
	err := row.Scan(
		&i.ID,
		&i.Make,
		&i.Model,
		&i.SerialNumber,
		&i.UnitID,
		&i.OpStatus,
		&i.CreatedAt,
		&i.ModifiedAt,
	)
	return i, err
}

const getEngines = `-- name: GetEngines :many
SELECT id, make, model, serial_number, unit_id, op_status, created_at, modified_at FROM assets.Engines
ORDER BY make
`

func (q *Queries) GetEngines(ctx context.Context) ([]AssetsEngine, error) {
	rows, err := q.db.QueryContext(ctx, getEngines)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []AssetsEngine{}
	for rows.Next() {
		var i AssetsEngine
		if err := rows.Scan(
			&i.ID,
			&i.Make,
			&i.Model,
			&i.SerialNumber,
			&i.UnitID,
			&i.OpStatus,
			&i.CreatedAt,
			&i.ModifiedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getEnginesByMake = `-- name: GetEnginesByMake :many
SELECT id, make, model, serial_number, unit_id, op_status, created_at, modified_at FROM assets.Engines
WHERE make = $1
ORDER BY model
`

func (q *Queries) GetEnginesByMake(ctx context.Context, make string) ([]AssetsEngine, error) {
	rows, err := q.db.QueryContext(ctx, getEnginesByMake, make)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []AssetsEngine{}
	for rows.Next() {
		var i AssetsEngine
		if err := rows.Scan(
			&i.ID,
			&i.Make,
			&i.Model,
			&i.SerialNumber,
			&i.UnitID,
			&i.OpStatus,
			&i.CreatedAt,
			&i.ModifiedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}