-- name: CreateProduct :one
INSERT INTO products (nome, descricao, preco, data_venda) VALUES ($1, $2, $3, $4) RETURNING *;

-- name: GetProduct :one
SELECT * FROM products WHERE id = $1;

-- name: ListProduct :many
SELECT * FROM products;

-- name: UpdateProduct :one
UPDATE products SET nome = $2, descricao = $3, preco = $4, data_venda = $5 WHERE id = $1 RETURNING *;

-- name: DeleteProduct :exec
DELETE FROM products WHERE id = $1;