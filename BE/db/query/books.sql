-- name: CreateBook :one
INSERT INTO
  books (
    title,
    description,
    image_links,
    price,
    publication_date,
    author_id
  )
VALUES
  ($1, $2, $3, $4, $5, $6) RETURNING *;


-- name: GetBookLists :many
SELECT
  B.*,
 (A.first_name || ' ' || A.last_name) as author_name
FROM
  books AS B
  INNER JOIN authors AS A on B.author_id = A.id
ORDER BY
  publication_date
LIMIT 100; 

-- name: GetBookById :one
SELECT
  B.*,
  (A.first_name || ' ' || A.last_name) as author_name
from
  books AS B
  INNER JOIN authors AS A on B.author_id = A.id
WHERE
  book_id = $1;

-- name: GetBooksByTitle :many
SELECT
  *
FROM
  books
WHERE
  title LIKE '%' || sqlc.arg(book_title) || '%';
  
-- name: UpdateBookById :one
UPDATE
  books
SET
  title = COALESCE(sqlc.narg(title), title),
  description = COALESCE(sqlc.narg(description), description),
  image_links = COALESCE(sqlc.narg(image_links), image_links),
  price = COALESCE(sqlc.narg(price), price),
  publication_date = COALESCE(sqlc.narg(publication_date), publication_date),
  author_id = COALESCE(sqlc.narg(author_id), author_id)
WHERE
  book_id = sqlc.arg(book_id) RETURNING *;


-- name: DeleteBook :exec
DELETE FROM
  books
WHERE
  book_id = $1;


-- name: DeleteManyBooks :exec
DELETE FROM
  books
WHERE
  book_id = ANY($1::int[]);