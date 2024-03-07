-- name: CreatePackSize :one
INSERT INTO product_pack_sizes (
  product_line,
  pack_size
) VALUES (
  $1, $2
) RETURNING *;

-- name: GetPackSize :one
SELECT * FROM product_pack_sizes
WHERE id = $1 LIMIT 1;


-- name: ListPackSizes :many
SELECT * FROM product_pack_sizes
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdatePackSizes :one
UPDATE product_pack_sizes
SET pack_size = $2,
    product_line = $3
WHERE id = $1
RETURNING *;