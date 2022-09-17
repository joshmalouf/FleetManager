// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: ClrSection.sql

package pgsql

import (
	"context"
	"database/sql"
)

const assignClrSection = `-- name: AssignClrSection :exec
UPDATE assets.clr_sections
SET opstatus = "active", cooler_id = $2
WHERE id = $1
`

type AssignClrSectionParams struct {
	ID       int64         `json:"id"`
	CoolerID sql.NullInt32 `json:"cooler_id"`
}

func (q *Queries) AssignClrSection(ctx context.Context, arg AssignClrSectionParams) error {
	_, err := q.db.ExecContext(ctx, assignClrSection, arg.ID, arg.CoolerID)
	return err
}

const createClrSection = `-- name: CreateClrSection :one
INSERT INTO assets.clr_sections
(make, serial_number, mawp, num_tubes, num_rows, passes)
VALUES
($1, $2, $3, $4, $5, $6)
RETURNING id, make, serial_number, mawp, num_tubes, num_rows, passes, cooler_id, op_status, created_at, modified_at
`

type CreateClrSectionParams struct {
	Make         string        `json:"make"`
	SerialNumber string        `json:"serial_number"`
	Mawp         int32         `json:"mawp"`
	NumTubes     sql.NullInt32 `json:"num_tubes"`
	NumRows      sql.NullInt32 `json:"num_rows"`
	Passes       sql.NullInt32 `json:"passes"`
}

func (q *Queries) CreateClrSection(ctx context.Context, arg CreateClrSectionParams) (AssetsClrSection, error) {
	row := q.db.QueryRowContext(ctx, createClrSection,
		arg.Make,
		arg.SerialNumber,
		arg.Mawp,
		arg.NumTubes,
		arg.NumRows,
		arg.Passes,
	)
	var i AssetsClrSection
	err := row.Scan(
		&i.ID,
		&i.Make,
		&i.SerialNumber,
		&i.Mawp,
		&i.NumTubes,
		&i.NumRows,
		&i.Passes,
		&i.CoolerID,
		&i.OpStatus,
		&i.CreatedAt,
		&i.ModifiedAt,
	)
	return i, err
}

const deactivateClrSection = `-- name: DeactivateClrSection :one
UPDATE assets.clr_sections
SET opstatus = "inactive", cooler_id = NULL
WHERE id = $1
RETURNING id, make, serial_number, mawp, num_tubes, num_rows, passes, cooler_id, op_status, created_at, modified_at
`

func (q *Queries) DeactivateClrSection(ctx context.Context, id int64) (AssetsClrSection, error) {
	row := q.db.QueryRowContext(ctx, deactivateClrSection, id)
	var i AssetsClrSection
	err := row.Scan(
		&i.ID,
		&i.Make,
		&i.SerialNumber,
		&i.Mawp,
		&i.NumTubes,
		&i.NumRows,
		&i.Passes,
		&i.CoolerID,
		&i.OpStatus,
		&i.CreatedAt,
		&i.ModifiedAt,
	)
	return i, err
}

const deleteClrSection = `-- name: DeleteClrSection :exec
DELETE FROM assets.clr_sections
WHERE id = $1
`

func (q *Queries) DeleteClrSection(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteClrSection, id)
	return err
}

const disposeClrSection = `-- name: DisposeClrSection :one
UPDATE assets.clr_sections
SET opstatus = "disposed", cooler_id = NULL
WHERE id = $1
RETURNING id, make, serial_number, mawp, num_tubes, num_rows, passes, cooler_id, op_status, created_at, modified_at
`

func (q *Queries) DisposeClrSection(ctx context.Context, id int64) (AssetsClrSection, error) {
	row := q.db.QueryRowContext(ctx, disposeClrSection, id)
	var i AssetsClrSection
	err := row.Scan(
		&i.ID,
		&i.Make,
		&i.SerialNumber,
		&i.Mawp,
		&i.NumTubes,
		&i.NumRows,
		&i.Passes,
		&i.CoolerID,
		&i.OpStatus,
		&i.CreatedAt,
		&i.ModifiedAt,
	)
	return i, err
}

const getClrSectionByID = `-- name: GetClrSectionByID :one
SELECT id, make, serial_number, mawp, num_tubes, num_rows, passes, cooler_id, op_status, created_at, modified_at FROM assets.clr_sections
WHERE id = $1
`

func (q *Queries) GetClrSectionByID(ctx context.Context, id int64) (AssetsClrSection, error) {
	row := q.db.QueryRowContext(ctx, getClrSectionByID, id)
	var i AssetsClrSection
	err := row.Scan(
		&i.ID,
		&i.Make,
		&i.SerialNumber,
		&i.Mawp,
		&i.NumTubes,
		&i.NumRows,
		&i.Passes,
		&i.CoolerID,
		&i.OpStatus,
		&i.CreatedAt,
		&i.ModifiedAt,
	)
	return i, err
}

const getClrSectionBySerial = `-- name: GetClrSectionBySerial :one
SELECT id, make, serial_number, mawp, num_tubes, num_rows, passes, cooler_id, op_status, created_at, modified_at FROM assets.clr_sections
WHERE serial_number = $1
`

func (q *Queries) GetClrSectionBySerial(ctx context.Context, serialNumber string) (AssetsClrSection, error) {
	row := q.db.QueryRowContext(ctx, getClrSectionBySerial, serialNumber)
	var i AssetsClrSection
	err := row.Scan(
		&i.ID,
		&i.Make,
		&i.SerialNumber,
		&i.Mawp,
		&i.NumTubes,
		&i.NumRows,
		&i.Passes,
		&i.CoolerID,
		&i.OpStatus,
		&i.CreatedAt,
		&i.ModifiedAt,
	)
	return i, err
}

const getClrSections = `-- name: GetClrSections :many
SELECT id, make, serial_number, mawp, num_tubes, num_rows, passes, cooler_id, op_status, created_at, modified_at FROM assets.clr_sections
ORDER BY make
`

func (q *Queries) GetClrSections(ctx context.Context) ([]AssetsClrSection, error) {
	rows, err := q.db.QueryContext(ctx, getClrSections)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []AssetsClrSection{}
	for rows.Next() {
		var i AssetsClrSection
		if err := rows.Scan(
			&i.ID,
			&i.Make,
			&i.SerialNumber,
			&i.Mawp,
			&i.NumTubes,
			&i.NumRows,
			&i.Passes,
			&i.CoolerID,
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
