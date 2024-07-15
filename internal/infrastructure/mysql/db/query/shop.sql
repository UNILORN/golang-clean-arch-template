-- name: ShopFindById :one
SELECT
   *
FROM
   shop
WHERE
   id = ?;

-- name: ShopFindByIds :many
SELECT
   *
FROM
   shop
WHERE
   id IN (sqlc.slice('ids'));

-- name: ShopFetch :many
SELECT
   *
FROM
  shop;

-- name: UpsertShop :exec
INSERT INTO shop (
   id,
   name,
   description,
   latitude,
   longitude
) VALUES (
   sqlc.arg(id),
   sqlc.arg(name),
   sqlc.arg(description),
   sqlc.arg(latitude),
   sqlc.arg(longitude)
) ON DUPLICATE KEY UPDATE
   name = sqlc.arg(name),
   description = sqlc.arg(description),
   latitude = sqlc.arg(latitude),
   longitude = sqlc.arg(longitude);