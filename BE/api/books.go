package api

import (
	"mime/multipart"
	"net/http"
	"path/filepath"

	db "github.com/amIToan/book_library/db/gen_sqlc"
	"github.com/amIToan/book_library/util"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/rs/zerolog/log"
)

type CreateBookReq struct {
	Title           string                  `form:"title" binding:"required"`
	Description     string                  `form:"description" binding:"required"`
	ImageLinks      []*multipart.FileHeader `form:"image_links" binding:"required,AtLeastOneFile"`
	Price           string                  `form:"price" binding:"required"`
	PublicationDate string                  `form:"publication_date" binding:"required"`
	AuthorID        int64                   `form:"author_id" binding:"required"`
	Genre           []int                   `form:"genre_id" binding:"required"`
}

func (server *Server) CreateBook(ctx *gin.Context) {
	isValid := CheckAuthAndExit(ctx)
	if !isValid {
		return
	}
	var form CreateBookReq
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}
	log.Info().Msgf("Genre ints : %T", form.Genre)
	book_genre, err := server.store.SetBookGenreAuthorTx(ctx, db.BookGenreAuthorTxParams{
		Title:           form.Title,
		Description:     form.Description,
		ImageLinks:      form.ImageLinks,
		Price:           form.Price,
		PublicationDate: form.PublicationDate,
		AuthorID:        form.AuthorID,
		Genre:           form.Genre,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, book_genre)
}

type GetBooksById struct {
	Id int64 `uri:"id" binding:"required"`
}

func (server *Server) GetBookById(ctx *gin.Context) {
	var req GetBooksById
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}
	book, err := server.store.GetBookById(ctx, req.Id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, book)
}

func (server *Server) GetBooks(ctx *gin.Context) {
	books, err := server.store.GetBookLists(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, books)
}

type UpdateBookReq struct {
	Title           string                  `form:"title"`
	Description     string                  `form:"description"`
	ImageLinks      []*multipart.FileHeader `form:"image_links"`
	Price           string                  `form:"price"`
	PublicationDate string                  `form:"publication_date"`
	AuthorID        int64                   `form:"author_id"`
	BookID          int64                   `form:"book_id" binding:"required"`
}

func (server *Server) UpdateBook(ctx *gin.Context) {
	isValid := CheckAuthAndExit(ctx)
	if !isValid {
		return
	}
	var form UpdateBookReq
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	arg := db.UpdateBookByIdParams{}
	if form.Title != "" {
		arg.Title = pgtype.Text{
			String: form.Title,
		}
	}
	if form.Description != "" {
		arg.Description = pgtype.Text{
			String: form.Title,
		}
	}
	log.Debug().Msgf("vo dy k ??/ %v", form.ImageLinks)
	if len(form.ImageLinks) > 0 {
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
				return
			}
		}
		arg.ImageLinks = imgSlice
	}
	if form.Price != "" {
		var pgPrice pgtype.Numeric
		_ = pgPrice.Scan(form.Price)
		arg.Price = pgPrice
	}
	if form.PublicationDate != "" {
		var pgDate pgtype.Date
		_ = pgDate.Scan(form.PublicationDate)
		arg.PublicationDate = pgDate
	}
	if form.AuthorID != 0 {
		arg.AuthorID = pgtype.Int8{Int64: form.AuthorID, Valid: true}
	}
	if form.BookID != 0 {
		arg.BookID = form.BookID
	}
	updatedBook, err := server.store.UpdateBookById(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, updatedBook)
}

type DeleteByIdReq struct {
	Id int64 `uri:"id" binding:"required"`
}

func (server *Server) DeleteBooks(ctx *gin.Context) {
	var IdsReq DeleteByIdReq
	isValid := CheckAuthAndExit(ctx)
	if !isValid {
		return
	}
	if err := ctx.ShouldBindUri(&IdsReq); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}
	err := server.store.DeleteBook(ctx, IdsReq.Id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "success",
	})
}

type Keywords struct {
	Keywords string `form:"keyword"`
}

func (server *Server) GetBooksByTitle(ctx *gin.Context) {
	var Key Keywords
	isValid := CheckAuthAndExit(ctx)
	if !isValid {
		return
	}
	if err := ctx.ShouldBind(&Key); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}
	pgKey := pgtype.Text{
		String: Key.Keywords,
		Valid:  true,
	}
	books, err := server.store.GetBooksByTitle(ctx, pgKey)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, books)
}
