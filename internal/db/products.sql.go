// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: products.sql

package db

import (
	"context"
	"time"
)

const adminGetProductByCode = `-- name: AdminGetProductByCode :one
select id, codigo, grupo, subgrupo, nombre, referencia, marca, unidad, existencia, precio1, precio2, precio3, precio4, precio5, precio6, precio7, discont, vta_max, vta_min, dctotope, enpreventa, comprometido, vta_minenx, vta_solofac, vta_solone, codbarras, detalles, cantbulto, costo_prom, util1, util2, util3, fchultcomp, qtyultcomp, created_at, updated_at, deleted_at
from articulo
where codigo = ?
`

func (q *Queries) AdminGetProductByCode(ctx context.Context, codigo string) (Articulo, error) {
	row := q.db.QueryRowContext(ctx, adminGetProductByCode, codigo)
	var i Articulo
	err := row.Scan(
		&i.ID,
		&i.Codigo,
		&i.Grupo,
		&i.Subgrupo,
		&i.Nombre,
		&i.Referencia,
		&i.Marca,
		&i.Unidad,
		&i.Existencia,
		&i.Precio1,
		&i.Precio2,
		&i.Precio3,
		&i.Precio4,
		&i.Precio5,
		&i.Precio6,
		&i.Precio7,
		&i.Discont,
		&i.VtaMax,
		&i.VtaMin,
		&i.Dctotope,
		&i.Enpreventa,
		&i.Comprometido,
		&i.VtaMinenx,
		&i.VtaSolofac,
		&i.VtaSolone,
		&i.Codbarras,
		&i.Detalles,
		&i.Cantbulto,
		&i.CostoProm,
		&i.Util1,
		&i.Util2,
		&i.Util3,
		&i.Fchultcomp,
		&i.Qtyultcomp,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const adminGetProductById = `-- name: AdminGetProductById :one
select id, codigo, grupo, subgrupo, nombre, referencia, marca, unidad, existencia, precio1, precio2, precio3, precio4, precio5, precio6, precio7, discont, vta_max, vta_min, dctotope, enpreventa, comprometido, vta_minenx, vta_solofac, vta_solone, codbarras, detalles, cantbulto, costo_prom, util1, util2, util3, fchultcomp, qtyultcomp, created_at, updated_at, deleted_at
from articulo
where id = ?
`

func (q *Queries) AdminGetProductById(ctx context.Context, id string) (Articulo, error) {
	row := q.db.QueryRowContext(ctx, adminGetProductById, id)
	var i Articulo
	err := row.Scan(
		&i.ID,
		&i.Codigo,
		&i.Grupo,
		&i.Subgrupo,
		&i.Nombre,
		&i.Referencia,
		&i.Marca,
		&i.Unidad,
		&i.Existencia,
		&i.Precio1,
		&i.Precio2,
		&i.Precio3,
		&i.Precio4,
		&i.Precio5,
		&i.Precio6,
		&i.Precio7,
		&i.Discont,
		&i.VtaMax,
		&i.VtaMin,
		&i.Dctotope,
		&i.Enpreventa,
		&i.Comprometido,
		&i.VtaMinenx,
		&i.VtaSolofac,
		&i.VtaSolone,
		&i.Codbarras,
		&i.Detalles,
		&i.Cantbulto,
		&i.CostoProm,
		&i.Util1,
		&i.Util2,
		&i.Util3,
		&i.Fchultcomp,
		&i.Qtyultcomp,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const adminGetProducts = `-- name: AdminGetProducts :many
select id, codigo, grupo, subgrupo, nombre, referencia, marca, unidad, existencia, precio1, precio2, precio3, precio4, precio5, precio6, precio7, discont, vta_max, vta_min, dctotope, enpreventa, comprometido, vta_minenx, vta_solofac, vta_solone, codbarras, detalles, cantbulto, costo_prom, util1, util2, util3, fchultcomp, qtyultcomp, created_at, updated_at, deleted_at
from articulo
`

func (q *Queries) AdminGetProducts(ctx context.Context) ([]Articulo, error) {
	rows, err := q.db.QueryContext(ctx, adminGetProducts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Articulo
	for rows.Next() {
		var i Articulo
		if err := rows.Scan(
			&i.ID,
			&i.Codigo,
			&i.Grupo,
			&i.Subgrupo,
			&i.Nombre,
			&i.Referencia,
			&i.Marca,
			&i.Unidad,
			&i.Existencia,
			&i.Precio1,
			&i.Precio2,
			&i.Precio3,
			&i.Precio4,
			&i.Precio5,
			&i.Precio6,
			&i.Precio7,
			&i.Discont,
			&i.VtaMax,
			&i.VtaMin,
			&i.Dctotope,
			&i.Enpreventa,
			&i.Comprometido,
			&i.VtaMinenx,
			&i.VtaSolofac,
			&i.VtaSolone,
			&i.Codbarras,
			&i.Detalles,
			&i.Cantbulto,
			&i.CostoProm,
			&i.Util1,
			&i.Util2,
			&i.Util3,
			&i.Fchultcomp,
			&i.Qtyultcomp,
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

const createProduct = `-- name: CreateProduct :exec
insert into articulo (
id,
codigo,
grupo,
subgrupo,
nombre,
referencia,
marca,
unidad,
existencia,
precio1,
precio2,
precio3,
precio4,
precio5,
precio6,
precio7,
discont,
vta_max,
vta_min,
dctotope,
enpreventa,
comprometido,
vta_minenx,
vta_solofac,
vta_solone,
codbarras,
detalles,
cantbulto,
costo_prom,
util1,
util2,
util3,
fchultcomp,
qtyultcomp,
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

type CreateProductParams struct {
	ID           string
	Codigo       string
	Grupo        int32
	Subgrupo     int32
	Nombre       string
	Referencia   string
	Marca        string
	Unidad       string
	Existencia   int32
	Precio1      string
	Precio2      string
	Precio3      string
	Precio4      string
	Precio5      string
	Precio6      string
	Precio7      string
	Discont      bool
	VtaMax       int32
	VtaMin       int32
	Dctotope     string
	Enpreventa   bool
	Comprometido int32
	VtaMinenx    int32
	VtaSolofac   bool
	VtaSolone    bool
	Codbarras    string
	Detalles     string
	Cantbulto    int32
	CostoProm    string
	Util1        string
	Util2        string
	Util3        string
	Fchultcomp   time.Time
	Qtyultcomp   int32
}

func (q *Queries) CreateProduct(ctx context.Context, arg CreateProductParams) error {
	_, err := q.db.ExecContext(ctx, createProduct,
		arg.ID,
		arg.Codigo,
		arg.Grupo,
		arg.Subgrupo,
		arg.Nombre,
		arg.Referencia,
		arg.Marca,
		arg.Unidad,
		arg.Existencia,
		arg.Precio1,
		arg.Precio2,
		arg.Precio3,
		arg.Precio4,
		arg.Precio5,
		arg.Precio6,
		arg.Precio7,
		arg.Discont,
		arg.VtaMax,
		arg.VtaMin,
		arg.Dctotope,
		arg.Enpreventa,
		arg.Comprometido,
		arg.VtaMinenx,
		arg.VtaSolofac,
		arg.VtaSolone,
		arg.Codbarras,
		arg.Detalles,
		arg.Cantbulto,
		arg.CostoProm,
		arg.Util1,
		arg.Util2,
		arg.Util3,
		arg.Fchultcomp,
		arg.Qtyultcomp,
	)
	return err
}

const getAllProducts = `-- name: GetAllProducts :many
select id, codigo, grupo, subgrupo, nombre, referencia, marca, unidad, existencia, precio1, precio2, precio3, precio4, precio5, precio6, precio7, discont, vta_max, vta_min, dctotope, enpreventa, comprometido, vta_minenx, vta_solofac, vta_solone, codbarras, detalles, cantbulto, costo_prom, util1, util2, util3, fchultcomp, qtyultcomp, created_at, updated_at, deleted_at
from articulo
where deleted_at is null
`

func (q *Queries) GetAllProducts(ctx context.Context) ([]Articulo, error) {
	rows, err := q.db.QueryContext(ctx, getAllProducts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Articulo
	for rows.Next() {
		var i Articulo
		if err := rows.Scan(
			&i.ID,
			&i.Codigo,
			&i.Grupo,
			&i.Subgrupo,
			&i.Nombre,
			&i.Referencia,
			&i.Marca,
			&i.Unidad,
			&i.Existencia,
			&i.Precio1,
			&i.Precio2,
			&i.Precio3,
			&i.Precio4,
			&i.Precio5,
			&i.Precio6,
			&i.Precio7,
			&i.Discont,
			&i.VtaMax,
			&i.VtaMin,
			&i.Dctotope,
			&i.Enpreventa,
			&i.Comprometido,
			&i.VtaMinenx,
			&i.VtaSolofac,
			&i.VtaSolone,
			&i.Codbarras,
			&i.Detalles,
			&i.Cantbulto,
			&i.CostoProm,
			&i.Util1,
			&i.Util2,
			&i.Util3,
			&i.Fchultcomp,
			&i.Qtyultcomp,
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

const getProductByCode = `-- name: GetProductByCode :one
select id, codigo, grupo, subgrupo, nombre, referencia, marca, unidad, existencia, precio1, precio2, precio3, precio4, precio5, precio6, precio7, discont, vta_max, vta_min, dctotope, enpreventa, comprometido, vta_minenx, vta_solofac, vta_solone, codbarras, detalles, cantbulto, costo_prom, util1, util2, util3, fchultcomp, qtyultcomp, created_at, updated_at, deleted_at
from articulo
where codigo = ? and deleted_at is null
`

func (q *Queries) GetProductByCode(ctx context.Context, codigo string) (Articulo, error) {
	row := q.db.QueryRowContext(ctx, getProductByCode, codigo)
	var i Articulo
	err := row.Scan(
		&i.ID,
		&i.Codigo,
		&i.Grupo,
		&i.Subgrupo,
		&i.Nombre,
		&i.Referencia,
		&i.Marca,
		&i.Unidad,
		&i.Existencia,
		&i.Precio1,
		&i.Precio2,
		&i.Precio3,
		&i.Precio4,
		&i.Precio5,
		&i.Precio6,
		&i.Precio7,
		&i.Discont,
		&i.VtaMax,
		&i.VtaMin,
		&i.Dctotope,
		&i.Enpreventa,
		&i.Comprometido,
		&i.VtaMinenx,
		&i.VtaSolofac,
		&i.VtaSolone,
		&i.Codbarras,
		&i.Detalles,
		&i.Cantbulto,
		&i.CostoProm,
		&i.Util1,
		&i.Util2,
		&i.Util3,
		&i.Fchultcomp,
		&i.Qtyultcomp,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getProductById = `-- name: GetProductById :one
select id, codigo, grupo, subgrupo, nombre, referencia, marca, unidad, existencia, precio1, precio2, precio3, precio4, precio5, precio6, precio7, discont, vta_max, vta_min, dctotope, enpreventa, comprometido, vta_minenx, vta_solofac, vta_solone, codbarras, detalles, cantbulto, costo_prom, util1, util2, util3, fchultcomp, qtyultcomp, created_at, updated_at, deleted_at
from articulo
where id = ? and deleted_at is null
`

func (q *Queries) GetProductById(ctx context.Context, id string) (Articulo, error) {
	row := q.db.QueryRowContext(ctx, getProductById, id)
	var i Articulo
	err := row.Scan(
		&i.ID,
		&i.Codigo,
		&i.Grupo,
		&i.Subgrupo,
		&i.Nombre,
		&i.Referencia,
		&i.Marca,
		&i.Unidad,
		&i.Existencia,
		&i.Precio1,
		&i.Precio2,
		&i.Precio3,
		&i.Precio4,
		&i.Precio5,
		&i.Precio6,
		&i.Precio7,
		&i.Discont,
		&i.VtaMax,
		&i.VtaMin,
		&i.Dctotope,
		&i.Enpreventa,
		&i.Comprometido,
		&i.VtaMinenx,
		&i.VtaSolofac,
		&i.VtaSolone,
		&i.Codbarras,
		&i.Detalles,
		&i.Cantbulto,
		&i.CostoProm,
		&i.Util1,
		&i.Util2,
		&i.Util3,
		&i.Fchultcomp,
		&i.Qtyultcomp,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const softDeleteProduct = `-- name: SoftDeleteProduct :exec
update articulo
set deleted_at = NOW()
where id = ?
`

func (q *Queries) SoftDeleteProduct(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, softDeleteProduct, id)
	return err
}

const updateProduct = `-- name: UpdateProduct :exec
update articulo
set grupo = ?,
    subgrupo = ?,
    nombre = ?,
    referencia = ?,
    marca = ?,
    unidad = ?,
    existencia = ?,
    precio1 = ?,
    precio2 = ?,
    precio3 = ?,
    precio4 = ?,
    precio5 = ?,
    precio6 = ?,
    precio7 = ?,
    discont = ?,
    vta_max = ?,
    vta_min = ?,
    dctotope = ?,
    enpreventa = ?,
    comprometido = ?,
    vta_minenx = ?,
    vta_solofac = ?,
    vta_solone = ?,
    codbarras = ?,
    detalles = ?,
    cantbulto = ?,
    costo_prom = ?,
    util1 = ?,
    util2 = ?,
    util3 = ?,
    fchultcomp = ?,
    qtyultcomp = ?,
    updated_at = NOW()
where id = ?
`

type UpdateProductParams struct {
	Grupo        int32
	Subgrupo     int32
	Nombre       string
	Referencia   string
	Marca        string
	Unidad       string
	Existencia   int32
	Precio1      string
	Precio2      string
	Precio3      string
	Precio4      string
	Precio5      string
	Precio6      string
	Precio7      string
	Discont      bool
	VtaMax       int32
	VtaMin       int32
	Dctotope     string
	Enpreventa   bool
	Comprometido int32
	VtaMinenx    int32
	VtaSolofac   bool
	VtaSolone    bool
	Codbarras    string
	Detalles     string
	Cantbulto    int32
	CostoProm    string
	Util1        string
	Util2        string
	Util3        string
	Fchultcomp   time.Time
	Qtyultcomp   int32
	ID           string
}

func (q *Queries) UpdateProduct(ctx context.Context, arg UpdateProductParams) error {
	_, err := q.db.ExecContext(ctx, updateProduct,
		arg.Grupo,
		arg.Subgrupo,
		arg.Nombre,
		arg.Referencia,
		arg.Marca,
		arg.Unidad,
		arg.Existencia,
		arg.Precio1,
		arg.Precio2,
		arg.Precio3,
		arg.Precio4,
		arg.Precio5,
		arg.Precio6,
		arg.Precio7,
		arg.Discont,
		arg.VtaMax,
		arg.VtaMin,
		arg.Dctotope,
		arg.Enpreventa,
		arg.Comprometido,
		arg.VtaMinenx,
		arg.VtaSolofac,
		arg.VtaSolone,
		arg.Codbarras,
		arg.Detalles,
		arg.Cantbulto,
		arg.CostoProm,
		arg.Util1,
		arg.Util2,
		arg.Util3,
		arg.Fchultcomp,
		arg.Qtyultcomp,
		arg.ID,
	)
	return err
}
