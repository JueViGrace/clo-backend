// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: configuration.sql

package db

import (
	"context"
	"time"
)

const adminGetConfigurationById = `-- name: AdminGetConfigurationById :one
select id, cnfg_idconfig, cnfg_clase, cnfg_tipo, cnfg_valnum, cnfg_valsino, cnfg_valtxt, cnfg_lentxt, cnfg_valfch, cnfg_activa, cnfg_etiq, cnfg_ttip, username, created_at, updated_at, deleted_at
from ke_wcnf_conf
where id = ?
`

func (q *Queries) AdminGetConfigurationById(ctx context.Context, id string) (KeWcnfConf, error) {
	row := q.db.QueryRowContext(ctx, adminGetConfigurationById, id)
	var i KeWcnfConf
	err := row.Scan(
		&i.ID,
		&i.CnfgIdconfig,
		&i.CnfgClase,
		&i.CnfgTipo,
		&i.CnfgValnum,
		&i.CnfgValsino,
		&i.CnfgValtxt,
		&i.CnfgLentxt,
		&i.CnfgValfch,
		&i.CnfgActiva,
		&i.CnfgEtiq,
		&i.CnfgTtip,
		&i.Username,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const adminGetConfigurations = `-- name: AdminGetConfigurations :many
select id, cnfg_idconfig, cnfg_clase, cnfg_tipo, cnfg_valnum, cnfg_valsino, cnfg_valtxt, cnfg_lentxt, cnfg_valfch, cnfg_activa, cnfg_etiq, cnfg_ttip, username, created_at, updated_at, deleted_at
from ke_wcnf_conf
`

func (q *Queries) AdminGetConfigurations(ctx context.Context) ([]KeWcnfConf, error) {
	rows, err := q.db.QueryContext(ctx, adminGetConfigurations)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []KeWcnfConf
	for rows.Next() {
		var i KeWcnfConf
		if err := rows.Scan(
			&i.ID,
			&i.CnfgIdconfig,
			&i.CnfgClase,
			&i.CnfgTipo,
			&i.CnfgValnum,
			&i.CnfgValsino,
			&i.CnfgValtxt,
			&i.CnfgLentxt,
			&i.CnfgValfch,
			&i.CnfgActiva,
			&i.CnfgEtiq,
			&i.CnfgTtip,
			&i.Username,
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

const getConfigurationsByUser = `-- name: GetConfigurationsByUser :many
select id, cnfg_idconfig, cnfg_clase, cnfg_tipo, cnfg_valnum, cnfg_valsino, cnfg_valtxt, cnfg_lentxt, cnfg_valfch, cnfg_activa, cnfg_etiq, cnfg_ttip, username, created_at, updated_at, deleted_at
from ke_wcnf_conf
where username = ? and deleted_at is null
`

func (q *Queries) GetConfigurationsByUser(ctx context.Context, username string) ([]KeWcnfConf, error) {
	rows, err := q.db.QueryContext(ctx, getConfigurationsByUser, username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []KeWcnfConf
	for rows.Next() {
		var i KeWcnfConf
		if err := rows.Scan(
			&i.ID,
			&i.CnfgIdconfig,
			&i.CnfgClase,
			&i.CnfgTipo,
			&i.CnfgValnum,
			&i.CnfgValsino,
			&i.CnfgValtxt,
			&i.CnfgLentxt,
			&i.CnfgValfch,
			&i.CnfgActiva,
			&i.CnfgEtiq,
			&i.CnfgTtip,
			&i.Username,
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

const getUserConfigurationById = `-- name: GetUserConfigurationById :one
select id, cnfg_idconfig, cnfg_clase, cnfg_tipo, cnfg_valnum, cnfg_valsino, cnfg_valtxt, cnfg_lentxt, cnfg_valfch, cnfg_activa, cnfg_etiq, cnfg_ttip, username, created_at, updated_at, deleted_at
from ke_wcnf_conf
where id = ? and username = ? and deleted_at is null
`

type GetUserConfigurationByIdParams struct {
	ID       string
	Username string
}

func (q *Queries) GetUserConfigurationById(ctx context.Context, arg GetUserConfigurationByIdParams) (KeWcnfConf, error) {
	row := q.db.QueryRowContext(ctx, getUserConfigurationById, arg.ID, arg.Username)
	var i KeWcnfConf
	err := row.Scan(
		&i.ID,
		&i.CnfgIdconfig,
		&i.CnfgClase,
		&i.CnfgTipo,
		&i.CnfgValnum,
		&i.CnfgValsino,
		&i.CnfgValtxt,
		&i.CnfgLentxt,
		&i.CnfgValfch,
		&i.CnfgActiva,
		&i.CnfgEtiq,
		&i.CnfgTtip,
		&i.Username,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const insertConfiguration = `-- name: InsertConfiguration :exec
INSERT INTO ke_wcnf_conf (
    id,
    cnfg_idconfig,
    cnfg_clase,
    cnfg_tipo,
    cnfg_valnum,
    cnfg_valsino,
    cnfg_valtxt,
    cnfg_lentxt,
    cnfg_valfch,
    cnfg_activa,
    cnfg_etiq,
    cnfg_ttip,
    username,
    created_at,
    updated_at
)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, NOW(), NOW())
`

type InsertConfigurationParams struct {
	ID           string
	CnfgIdconfig string
	CnfgClase    string
	CnfgTipo     string
	CnfgValnum   string
	CnfgValsino  bool
	CnfgValtxt   string
	CnfgLentxt   int16
	CnfgValfch   time.Time
	CnfgActiva   bool
	CnfgEtiq     string
	CnfgTtip     string
	Username     string
}

func (q *Queries) InsertConfiguration(ctx context.Context, arg InsertConfigurationParams) error {
	_, err := q.db.ExecContext(ctx, insertConfiguration,
		arg.ID,
		arg.CnfgIdconfig,
		arg.CnfgClase,
		arg.CnfgTipo,
		arg.CnfgValnum,
		arg.CnfgValsino,
		arg.CnfgValtxt,
		arg.CnfgLentxt,
		arg.CnfgValfch,
		arg.CnfgActiva,
		arg.CnfgEtiq,
		arg.CnfgTtip,
		arg.Username,
	)
	return err
}

const softDeleteConfiguration = `-- name: SoftDeleteConfiguration :exec
UPDATE ke_wcnf_conf
SET deleted_at = NOW()
WHERE id = ?
`

func (q *Queries) SoftDeleteConfiguration(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, softDeleteConfiguration, id)
	return err
}

const updateConfiguration = `-- name: UpdateConfiguration :exec
UPDATE ke_wcnf_conf
SET cnfg_idconfig = ?,
    cnfg_clase = ?,
    cnfg_tipo = ?,
    cnfg_valnum = ?,
    cnfg_valsino = ?,
    cnfg_valtxt = ?,
    cnfg_lentxt = ?,
    cnfg_valfch = ?,
    cnfg_activa = ?,
    cnfg_etiq = ?,
    cnfg_ttip = ?,
    username = ?,
    updated_at = NOW() 
WHERE id = ?
`

type UpdateConfigurationParams struct {
	CnfgIdconfig string
	CnfgClase    string
	CnfgTipo     string
	CnfgValnum   string
	CnfgValsino  bool
	CnfgValtxt   string
	CnfgLentxt   int16
	CnfgValfch   time.Time
	CnfgActiva   bool
	CnfgEtiq     string
	CnfgTtip     string
	Username     string
	ID           string
}

func (q *Queries) UpdateConfiguration(ctx context.Context, arg UpdateConfigurationParams) error {
	_, err := q.db.ExecContext(ctx, updateConfiguration,
		arg.CnfgIdconfig,
		arg.CnfgClase,
		arg.CnfgTipo,
		arg.CnfgValnum,
		arg.CnfgValsino,
		arg.CnfgValtxt,
		arg.CnfgLentxt,
		arg.CnfgValfch,
		arg.CnfgActiva,
		arg.CnfgEtiq,
		arg.CnfgTtip,
		arg.Username,
		arg.ID,
	)
	return err
}
