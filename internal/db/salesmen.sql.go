// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: salesmen.sql

package db

import (
	"context"
	"database/sql"
)

const adminGetSalesmaByCode = `-- name: AdminGetSalesmaByCode :one
select id, user_id, codigo, nombre, telefono_1, telefono_2, telefono_movil, status, supervpor, sector, subcodigo, email, created_at, updated_at, deleted_at
from vendedor
where codigo = ?
`

func (q *Queries) AdminGetSalesmaByCode(ctx context.Context, codigo string) (Vendedor, error) {
	row := q.db.QueryRowContext(ctx, adminGetSalesmaByCode, codigo)
	var i Vendedor
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Codigo,
		&i.Nombre,
		&i.Telefono1,
		&i.Telefono2,
		&i.TelefonoMovil,
		&i.Status,
		&i.Supervpor,
		&i.Sector,
		&i.Subcodigo,
		&i.Email,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const adminGetSalesmanById = `-- name: AdminGetSalesmanById :one
select id, user_id, codigo, nombre, telefono_1, telefono_2, telefono_movil, status, supervpor, sector, subcodigo, email, created_at, updated_at, deleted_at
from vendedor
where id = ?
`

func (q *Queries) AdminGetSalesmanById(ctx context.Context, id string) (Vendedor, error) {
	row := q.db.QueryRowContext(ctx, adminGetSalesmanById, id)
	var i Vendedor
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Codigo,
		&i.Nombre,
		&i.Telefono1,
		&i.Telefono2,
		&i.TelefonoMovil,
		&i.Status,
		&i.Supervpor,
		&i.Sector,
		&i.Subcodigo,
		&i.Email,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const adminGetSalesmen = `-- name: AdminGetSalesmen :many
select id, user_id, codigo, nombre, telefono_1, telefono_2, telefono_movil, status, supervpor, sector, subcodigo, email, created_at, updated_at, deleted_at
from vendedor
`

func (q *Queries) AdminGetSalesmen(ctx context.Context) ([]Vendedor, error) {
	rows, err := q.db.QueryContext(ctx, adminGetSalesmen)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Vendedor
	for rows.Next() {
		var i Vendedor
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Codigo,
			&i.Nombre,
			&i.Telefono1,
			&i.Telefono2,
			&i.TelefonoMovil,
			&i.Status,
			&i.Supervpor,
			&i.Sector,
			&i.Subcodigo,
			&i.Email,
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

const createSalesman = `-- name: CreateSalesman :exec
insert into vendedor (
    id,
    user_id,
    codigo,
    nombre,
    telefono_1,
    telefono_2,
    telefono_movil,
    status,
    supervpor,
    sector,
    subcodigo,
    email,
    created_at,
    updated_at
)
values (
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    NOW(),
    NOW()
)
`

type CreateSalesmanParams struct {
	ID            string
	UserID        sql.NullString
	Codigo        string
	Nombre        string
	Telefono1     string
	Telefono2     string
	TelefonoMovil string
	Status        int16
	Supervpor     string
	Sector        int32
	Subcodigo     int32
	Email         string
}

func (q *Queries) CreateSalesman(ctx context.Context, arg CreateSalesmanParams) error {
	_, err := q.db.ExecContext(ctx, createSalesman,
		arg.ID,
		arg.UserID,
		arg.Codigo,
		arg.Nombre,
		arg.Telefono1,
		arg.Telefono2,
		arg.TelefonoMovil,
		arg.Status,
		arg.Supervpor,
		arg.Sector,
		arg.Subcodigo,
		arg.Email,
	)
	return err
}

const getSalesmanByCode = `-- name: GetSalesmanByCode :one
select id, user_id, codigo, nombre, telefono_1, telefono_2, telefono_movil, status, supervpor, sector, subcodigo, email, created_at, updated_at, deleted_at
from vendedor
where codigo = ? and deleted_at is null
`

func (q *Queries) GetSalesmanByCode(ctx context.Context, codigo string) (Vendedor, error) {
	row := q.db.QueryRowContext(ctx, getSalesmanByCode, codigo)
	var i Vendedor
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Codigo,
		&i.Nombre,
		&i.Telefono1,
		&i.Telefono2,
		&i.TelefonoMovil,
		&i.Status,
		&i.Supervpor,
		&i.Sector,
		&i.Subcodigo,
		&i.Email,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getSalesmanById = `-- name: GetSalesmanById :one
select id, user_id, codigo, nombre, telefono_1, telefono_2, telefono_movil, status, supervpor, sector, subcodigo, email, created_at, updated_at, deleted_at
from vendedor
where id = ? and deleted_at is null
`

func (q *Queries) GetSalesmanById(ctx context.Context, id string) (Vendedor, error) {
	row := q.db.QueryRowContext(ctx, getSalesmanById, id)
	var i Vendedor
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Codigo,
		&i.Nombre,
		&i.Telefono1,
		&i.Telefono2,
		&i.TelefonoMovil,
		&i.Status,
		&i.Supervpor,
		&i.Sector,
		&i.Subcodigo,
		&i.Email,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getSalesmenByManager = `-- name: GetSalesmenByManager :many
select id, user_id, codigo, nombre, telefono_1, telefono_2, telefono_movil, status, supervpor, sector, subcodigo, email, created_at, updated_at, deleted_at
from vendedor
`

func (q *Queries) GetSalesmenByManager(ctx context.Context) ([]Vendedor, error) {
	rows, err := q.db.QueryContext(ctx, getSalesmenByManager)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Vendedor
	for rows.Next() {
		var i Vendedor
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Codigo,
			&i.Nombre,
			&i.Telefono1,
			&i.Telefono2,
			&i.TelefonoMovil,
			&i.Status,
			&i.Supervpor,
			&i.Sector,
			&i.Subcodigo,
			&i.Email,
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

const softDeleteSalesman = `-- name: SoftDeleteSalesman :exec
update vendedor
set deleted_at = NOW()
where id = ?
`

func (q *Queries) SoftDeleteSalesman(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, softDeleteSalesman, id)
	return err
}

const updateSalesman = `-- name: UpdateSalesman :exec
update vendedor
set codigo = ?,
    nombre = ?,
    telefono_1 = ?,
    telefono_2 = ?,
    telefono_movil = ?,
    status = ?,
    supervpor = ?,
    sector = ?,
    subcodigo = ?,
    email = ?,
    updated_at = NOW()
where id = ?
`

type UpdateSalesmanParams struct {
	Codigo        string
	Nombre        string
	Telefono1     string
	Telefono2     string
	TelefonoMovil string
	Status        int16
	Supervpor     string
	Sector        int32
	Subcodigo     int32
	Email         string
	ID            string
}

func (q *Queries) UpdateSalesman(ctx context.Context, arg UpdateSalesmanParams) error {
	_, err := q.db.ExecContext(ctx, updateSalesman,
		arg.Codigo,
		arg.Nombre,
		arg.Telefono1,
		arg.Telefono2,
		arg.TelefonoMovil,
		arg.Status,
		arg.Supervpor,
		arg.Sector,
		arg.Subcodigo,
		arg.Email,
		arg.ID,
	)
	return err
}
