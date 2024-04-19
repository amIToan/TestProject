-- name: CreateBookGenres :copyfrom
INSERT INTO
  book_genres (book_id, genre_id)
VALUES
  ($1, $2);


-- name: DeleteBookGenre :exec
DELETE FROM
  book_genres
WHERE
  genre_id = $1;


-- name: DeleteGenreByBookIds :exec
DELETE FROM
  book_genres
WHERE
  book_id = ANY($1::int[]);
  
-- name: GetBookGenres :many
SELECT
  bg.*,
  g.name
from
  book_genres bg
  join genres g on bg.genre_id = g.genre_id
where
  bg.book_id = $1;