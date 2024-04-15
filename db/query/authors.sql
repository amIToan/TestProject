-- name: CreateAuthor
INSERT INTO authors (
    first_name,
    last_name, 
    dob, 
    email,
    gender
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetAuthorsWithLimitation 
SELECT * FROM authors LIMIT 50; 

-- name: GetAuthorsByNames :many
SELECT * FROM authors
WHERE CONCAT(first_name, ' ', last_name) LIKE '%' || sqlc.arg(amount) || '%';

-- name: UpdateAuthor :one
UPDATE authors
SET
  first_name = COALESCE(sqlc.narg(first_name), first_name),
  last_name = COALESCE(sqlc.narg(last_name), last_name),
  dob = COALESCE(sqlc.narg(dob), dob),
  email = COALESCE(sqlc.narg(email), email),
  gender = COALESCE(sqlc.narg(gender), gender)
WHERE
  id = sqlc.arg(id)
RETURNING *;