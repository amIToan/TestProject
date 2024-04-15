-- name : CreateBookGenres 
INSERT INTO book_genres (
    book_id, 
    genre_id
) VALUES (
    $1, $2
) RETURNING *;
--name : DeleteBookGenre
DELETE FROM book_genres WHERE genre_id = $1; 
