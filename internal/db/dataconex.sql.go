// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: dataconex.sql

package db

import (
	"context"
)

<<<<<<< HEAD
const createCompany = `-- name: CreateCompany :one
=======
const createCompany = `-- name: CreateCompany :exec
>>>>>>> 14869e
INSERT INTO ke_dataconex (
        ked_codigo,
        ked_nombre,
        ked_status,
        ked_enlace,
        ked_agen,
        created_at,
        updated_at
    )
VALUES ($1, $2, $3, $4, $5, NOW(), NOW())
<<<<<<< HEAD
RETURNING ked_codigo, ked_nombre, ked_status, ked_enlace, ked_agen, created_at, updated_at, deleted_at
`

type CreateCompanyParams struct {
	KedCodigo string
	KedNombre string
	KedStatus int32
	KedEnlace string
	KedAgen   string
}

func (q *Queries) CreateCompany(ctx context.Context, arg CreateCompanyParams) (KeDataconex, error) {
	row := q.db.QueryRowContext(ctx, createCompany,
		arg.KedCodigo,
		arg.KedNombre,
		arg.KedStatus,
		arg.KedEnlace,
		arg.KedAgen,
	)
	var i KeDataconex
	err := row.Scan(
		&i.KedCodigo,
		&i.KedNombre,
		&i.KedStatus,
		&i.KedEnlace,
		&i.KedAgen,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
=======
`

func (q *Queries) CreateCompany(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, createCompany)
	return err
>>>>>>> 14869e
}

const deleteCompany = `-- name: DeleteCompany :exec
UPDATE ke_dataconex
SET ked_status = 0,
    deleted_at = NOW()
WHERE ked_codigo = $1
`

<<<<<<< HEAD
func (q *Queries) DeleteCompany(ctx context.Context, kedCodigo string) error {
	_, err := q.db.ExecContext(ctx, deleteCompany, kedCodigo)
=======
func (q *Queries) DeleteCompany(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, deleteCompany)
>>>>>>> 14869e
	return err
}

const getAllCompanies = `-- name: GetAllCompanies :many
SELECT ked_codigo, ked_nombre, ked_status, ked_enlace, ked_agen, created_at, updated_at, deleted_at FROM ke_dataconex
`

func (q *Queries) GetAllCompanies(ctx context.Context) ([]KeDataconex, error) {
	rows, err := q.db.QueryContext(ctx, getAllCompanies)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []KeDataconex
	for rows.Next() {
		var i KeDataconex
		if err := rows.Scan(
			&i.KedCodigo,
			&i.KedNombre,
			&i.KedStatus,
			&i.KedEnlace,
			&i.KedAgen,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
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

const getCompanies = `-- name: GetCompanies :many
SELECT ked_codigo, ked_nombre, ked_status, ked_enlace, ked_agen, created_at, updated_at, deleted_at FROM ke_dataconex
WHERE deleted_at IS NULL
`

func (q *Queries) GetCompanies(ctx context.Context) ([]KeDataconex, error) {
	rows, err := q.db.QueryContext(ctx, getCompanies)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []KeDataconex
	for rows.Next() {
		var i KeDataconex
		if err := rows.Scan(
			&i.KedCodigo,
			&i.KedNombre,
			&i.KedStatus,
			&i.KedEnlace,
			&i.KedAgen,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
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

const getCompanyById = `-- name: GetCompanyById :one
SELECT ked_codigo, ked_nombre, ked_status, ked_enlace, ked_agen, created_at, updated_at, deleted_at FROM ke_dataconex
WHERE 
ked_codigo = $1 AND
ked_status = 1 AND
deleted_at IS NULL
`

<<<<<<< HEAD
func (q *Queries) GetCompanyById(ctx context.Context, kedCodigo string) (KeDataconex, error) {
	row := q.db.QueryRowContext(ctx, getCompanyById, kedCodigo)
=======
func (q *Queries) GetCompanyById(ctx context.Context) (KeDataconex, error) {
	row := q.db.QueryRowContext(ctx, getCompanyById)
>>>>>>> 14869e
	var i KeDataconex
	err := row.Scan(
		&i.KedCodigo,
		&i.KedNombre,
		&i.KedStatus,
		&i.KedEnlace,
		&i.KedAgen,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getOneCompanyById = `-- name: GetOneCompanyById :one
SELECT ked_codigo, ked_nombre, ked_status, ked_enlace, ked_agen, created_at, updated_at, deleted_at FROM ke_dataconex
WHERE ked_codigo = $1
`

<<<<<<< HEAD
func (q *Queries) GetOneCompanyById(ctx context.Context, kedCodigo string) (KeDataconex, error) {
	row := q.db.QueryRowContext(ctx, getOneCompanyById, kedCodigo)
=======
func (q *Queries) GetOneCompanyById(ctx context.Context) (KeDataconex, error) {
	row := q.db.QueryRowContext(ctx, getOneCompanyById)
>>>>>>> 14869e
	var i KeDataconex
	err := row.Scan(
		&i.KedCodigo,
		&i.KedNombre,
		&i.KedStatus,
		&i.KedEnlace,
		&i.KedAgen,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

<<<<<<< HEAD
const updateCompany = `-- name: UpdateCompany :one
=======
const updateCompany = `-- name: UpdateCompany :exec
>>>>>>> 14869e
UPDATE ke_dataconex
SET ked_codigo = $1,
    ked_nombre = $2,
    ked_status = $3,
    ked_enlace = $4,
<<<<<<< HEAD
    ked_agen = $5,
    updated_at = NOW()
WHERE ked_codigo = $6
RETURNING ked_codigo, ked_nombre, ked_status, ked_enlace, ked_agen, created_at, updated_at, deleted_at
`

type UpdateCompanyParams struct {
	KedCodigo   string
	KedNombre   string
	KedStatus   int32
	KedEnlace   string
	KedAgen     string
	KedCodigo_2 string
}

func (q *Queries) UpdateCompany(ctx context.Context, arg UpdateCompanyParams) (KeDataconex, error) {
	row := q.db.QueryRowContext(ctx, updateCompany,
		arg.KedCodigo,
		arg.KedNombre,
		arg.KedStatus,
		arg.KedEnlace,
		arg.KedAgen,
		arg.KedCodigo_2,
	)
	var i KeDataconex
	err := row.Scan(
		&i.KedCodigo,
		&i.KedNombre,
		&i.KedStatus,
		&i.KedEnlace,
		&i.KedAgen,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
=======
    ked_agen = 5$,
    updated_at = NOW()
WHERE ked_codigo = $6
`

func (q *Queries) UpdateCompany(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, updateCompany)
	return err
>>>>>>> 14869e
}
