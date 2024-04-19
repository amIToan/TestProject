package db

import (
	"errors"
	"mime/multipart"
	"net/http"
	"path/filepath"

	"github.com/amIToan/book_library/util"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
)

// to assert that structs have to implement
type StorageInterface interface {
	Querier
	SetBookGenreAuthorTx(ctx *gin.Context, arg BookGenreAuthorTxParams) (BookGenreAuthorTxResult, error)
}

// real struct to manipulate database
type SQLStorage struct {
	conn *pgxpool.Pool
	*Queries
}

// NewStore creates a new store
func NewStore(connPool *pgxpool.Pool) StorageInterface {
	return &SQLStorage{
		conn:    connPool,
		Queries: New(connPool),
	}
}

type BookGenreAuthorTxParams struct {
	Title           string
	Description     string
	ImageLinks      []*multipart.FileHeader
	Price           string
	PublicationDate string
	AuthorID        int64
	Genre           []int
}
type BookGenreAuthorTxResult struct {
	Book       Book  `json:"book"`
	Book_genre int64 `json:"book_genre"`
}

func (sqlStorage *SQLStorage) SetBookGenreAuthorTx(ctx *gin.Context, form BookGenreAuthorTxParams) (BookGenreAuthorTxResult, error) {
	var _BookGenreAuthorTxResult BookGenreAuthorTxResult
	err := sqlStorage.ExecTx(ctx, func(q *Queries) error {
		imgSlice := []string{}
		for _, file := range form.ImageLinks {
			tail := filepath.Ext(file.Filename)
			uuid := uuid.New()
			newFileName := form.Title + "_" + file.Filename + "_" + uuid.String() + tail
			imgLink := "/assets/imgs/" + newFileName
			imgSlice = append(imgSlice, imgLink)
			filePath := filepath.Join("assets/imgs", newFileName)
			err := ctx.SaveUploadedFile(file, filePath)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
				return err
			}
		}
		var pgPrice pgtype.Numeric
		_ = pgPrice.Scan(form.Price)
		var pgDate pgtype.Date
		_ = pgDate.Scan(form.PublicationDate)
		arg := CreateBookParams{
			Title:           form.Title,
			Description:     form.Description,
			ImageLinks:      imgSlice,
			Price:           pgPrice,
			PublicationDate: pgDate,
			AuthorID:        form.AuthorID,
		}
		book, err := sqlStorage.CreateBook(ctx, arg)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
			return err
		}
		_BookGenreAuthorTxResult.Book = book
		if len(form.Genre) > 0 {
			_CreateBookGenresParams := []CreateBookGenresParams{}
			for _, g := range form.Genre {
				_CreateBookGenresParams = append(_CreateBookGenresParams, CreateBookGenresParams{BookID: book.BookID, GenreID: int64(g)})
			}
			book_genre, err := sqlStorage.CreateBookGenres(ctx, _CreateBookGenresParams)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
				return err
			}
			_BookGenreAuthorTxResult.Book_genre = book_genre
		} else {
			return errors.New("no genres")
		}
		log.Info().Msgf("Inside to check _BookGenreAuthorTxResult : %v", _BookGenreAuthorTxResult)
		return nil
	})
	return _BookGenreAuthorTxResult, err
}
