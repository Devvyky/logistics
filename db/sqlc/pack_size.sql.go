// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: pack_size.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const createPackSize = `-- name: CreatePackSize :one
INSERT INTO product_pack_sizes (
  product_line,
  pack_size
) VALUES (
  $1, $2
) RETURNING id, product_line, pack_size, updated_at, created_at
`

type CreatePackSizeParams struct {
	ProductLine string `json:"product_line"`
	PackSize    int64  `json:"pack_size"`
}

func (q *Queries) CreatePackSize(ctx context.Context, arg CreatePackSizeParams) (ProductPackSize, error) {
	row := q.db.QueryRowContext(ctx, createPackSize, arg.ProductLine, arg.PackSize)
	var i ProductPackSize
	err := row.Scan(
		&i.ID,
		&i.ProductLine,
		&i.PackSize,
		&i.UpdatedAt,
		&i.CreatedAt,
	)
	return i, err
}

const deletePackSize = `-- name: DeletePackSize :exec
DELETE FROM product_pack_sizes WHERE id = $1
`

func (q *Queries) DeletePackSize(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deletePackSize, id)
	return err
}

const getPackSize = `-- name: GetPackSize :one
SELECT id, product_line, pack_size, updated_at, created_at FROM product_pack_sizes
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetPackSize(ctx context.Context, id uuid.UUID) (ProductPackSize, error) {
	row := q.db.QueryRowContext(ctx, getPackSize, id)
	var i ProductPackSize
	err := row.Scan(
		&i.ID,
		&i.ProductLine,
		&i.PackSize,
		&i.UpdatedAt,
		&i.CreatedAt,
	)
	return i, err
}

const listPackSizes = `-- name: ListPackSizes :many
SELECT id, product_line, pack_size, updated_at, created_at FROM product_pack_sizes
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListPackSizesParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListPackSizes(ctx context.Context, arg ListPackSizesParams) ([]ProductPackSize, error) {
	rows, err := q.db.QueryContext(ctx, listPackSizes, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ProductPackSize{}
	for rows.Next() {
		var i ProductPackSize
		if err := rows.Scan(
			&i.ID,
			&i.ProductLine,
			&i.PackSize,
			&i.UpdatedAt,
			&i.CreatedAt,
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

const listPackSizesByProductLine = `-- name: ListPackSizesByProductLine :many
SELECT id, product_line, pack_size, updated_at, created_at FROM product_pack_sizes
WHERE product_line = $1
ORDER by pack_size ASC
`

func (q *Queries) ListPackSizesByProductLine(ctx context.Context, productLine string) ([]ProductPackSize, error) {
	rows, err := q.db.QueryContext(ctx, listPackSizesByProductLine, productLine)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ProductPackSize{}
	for rows.Next() {
		var i ProductPackSize
		if err := rows.Scan(
			&i.ID,
			&i.ProductLine,
			&i.PackSize,
			&i.UpdatedAt,
			&i.CreatedAt,
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

const listProductLines = `-- name: ListProductLines :many
SELECT DISTINCT product_line
FROM product_pack_sizes
`

func (q *Queries) ListProductLines(ctx context.Context) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, listProductLines)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []string{}
	for rows.Next() {
		var product_line string
		if err := rows.Scan(&product_line); err != nil {
			return nil, err
		}
		items = append(items, product_line)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updatePackSizes = `-- name: UpdatePackSizes :one
UPDATE product_pack_sizes
SET pack_size = $2,
    product_line = $3
WHERE id = $1
RETURNING id, product_line, pack_size, updated_at, created_at
`

type UpdatePackSizesParams struct {
	ID          uuid.UUID `json:"id"`
	PackSize    int64     `json:"pack_size"`
	ProductLine string    `json:"product_line"`
}

func (q *Queries) UpdatePackSizes(ctx context.Context, arg UpdatePackSizesParams) (ProductPackSize, error) {
	row := q.db.QueryRowContext(ctx, updatePackSizes, arg.ID, arg.PackSize, arg.ProductLine)
	var i ProductPackSize
	err := row.Scan(
		&i.ID,
		&i.ProductLine,
		&i.PackSize,
		&i.UpdatedAt,
		&i.CreatedAt,
	)
	return i, err
}
