-- name: CreateGenre
INSERT INTO genres (
    name
) VALUES (
    $1
) RETURNING *;

-- name: GetAuthorsWithLimitation 
SELECT * FROM genres LIMIT 50; 

-- name: GetGenreByNames :many
SELECT * FROM genres
WHERE name LIKE ('%' || sqlc.arg(name) || '%');

-- name: UpdateGenre :one
UPDATE genres
SET
  name = $2
WHERE
  genre_id = $1
RETURNING *;

--name: DeleteGenre: one 
DELETE FROM genres WHERE genre_id = $1; 