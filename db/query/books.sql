-- name: CreateABook
INSERT INTO books (
    title, 
    description, 
    price, 
    publication_date, 
    author_id
) VALUES (
  $1, $2, $3, $4, $5
) RETURNING *;

-- name: BookList :many
SELECT * FROM books
ORDER BY publication_date
LIMIT $2
OFFSET $3;

-- name : GetBookById 
SELECT * from books
WHERE book_id = $1; 

-- name : UpdateBookById 
UPDATE books
SET
    title = COALESCE(sqlc.narg(title), title), 
    description = COALESCE(sqlc.narg(description), description), 
    price = COALESCE(sqlc.narg(price), price), 
    publication_date = COALESCE(sqlc.narg(publication_date), publication_date), 
    author_id = COALESCE(sqlc.narg(author_id), author_id)
WHERE 
    book_id = sqlc.arg(book_id)
RETURNING *;

--name: DeleteBook: one 
DELETE FROM genres WHERE book_id = $1; 

--name: DeleteManyBooks: many 
DELETE FROM genres
WHERE book_id IN (sqlc.slice('book_ids'));