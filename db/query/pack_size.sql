-- name: CreatePackSize :one
INSERT INTO product_pack_sizes (
  product_line,
  pack_size
) VALUES (
  $1, $2
) RETURNING *;