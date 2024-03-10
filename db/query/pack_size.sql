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

-- name: ListPackSizesByProductLine :many
SELECT * FROM product_pack_sizes
WHERE product_line = $1
ORDER by pack_size ASC;

-- name: UpdatePackSizes :one
UPDATE product_pack_sizes
SET pack_size = $2,
    product_line = $3
WHERE id = $1
RETURNING *;

-- name: DeletePackSize :exec
DELETE FROM product_pack_sizes WHERE id = $1;

-- name: ListProductLines :many
SELECT DISTINCT product_line
FROM product_pack_sizes;