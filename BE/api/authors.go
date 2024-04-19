package api

import (
	"errors"
	"net/http"

	db "github.com/amIToan/book_library/db/gen_sqlc"
	"github.com/amIToan/book_library/util"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

type CreateAuthorReq struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name"  binding:"required"`
	Dob       string `json:"dob"  binding:"required"`
	Email     string `json:"email"  binding:"required,email"`
	Gender    string `json:"gender"  binding:"required"`
}

func (server *Server) CreateAuthor(ctx *gin.Context) {
	var req CreateAuthorReq
	isValid := CheckAuthAndExit(ctx)
	if !isValid {
		return
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}
	var pgDate pgtype.Date
	err := pgDate.Scan(req.Dob)
	if err != nil {
		newErr := errors.New("wrong date format")
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(newErr))
		return
	}
	arg := db.CreateAuthorParams{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Dob:       pgDate,
		Email:     req.Email,
		Gender:    req.Gender,
	}
	author, err := server.store.CreateAuthor(ctx, arg)
	if err != nil {
		if util.ErrorCode(err) == util.UniqueViolation || util.ErrorCode(err) == util.ForeignKeyViolation {
			ctx.JSON(http.StatusForbidden, util.ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, author)
}

func (server *Server) GetFiftyAuthors(ctx *gin.Context) {
	authors, err := server.store.GetFirstFiftyOfAuthors(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, authors)
}

type GetAuthorsByNameReq struct {
	Author_name string `uri:"author_name" binding:"required"`
}

func (server *Server) GetAuthorsByName(ctx *gin.Context) {
	var req GetAuthorsByNameReq
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}
	isValid := CheckAuthAndExit(ctx)
	if !isValid {
		return
	}
	var pgText pgtype.Text
	_ = pgText.Scan(req.Author_name)

	author, err := server.store.GetAuthorsByNames(ctx, pgText)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, author)
}

type UpdateAuthorReq struct {
	FirstName string `json:"first_name" `
	LastName  string `json:"last_name"  `
	Dob       string `json:"dob"`
	Email     string `json:"email" binding:"email"`
	Gender    string `json:"gender"`
	ID        int64  `json:"id" binding:"required"`
}

func (server *Server) UpdateAuthor(ctx *gin.Context) {
	var req UpdateAuthorReq
	isValid := CheckAuthAndExit(ctx)
	if !isValid {
		return
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}
	arg := db.UpdateAuthorParams{
		ID: req.ID,
	}
	if req.FirstName != "" {
		arg.FirstName = pgtype.Text{
			String: req.FirstName,
			Valid:  true,
		}
	}
	if req.LastName != "" {
		arg.LastName = pgtype.Text{
			String: req.LastName,
			Valid:  true,
		}
	}
	if req.Dob != "" {
		var pgDate pgtype.Date
		err := pgDate.Scan(req.Dob)
		if err != nil {
			newErr := errors.New("wrong date format")
			ctx.JSON(http.StatusBadRequest, util.ErrorResponse(newErr))
			return
		}
		arg.Dob = pgDate
	}
	if req.Email != "" {
		arg.Email = pgtype.Text{
			String: req.Email,
			Valid:  true,
		}
	}
	if req.Gender != "" {
		arg.Gender = pgtype.Text{
			String: req.Gender,
			Valid:  true,
		}
	}
	updateAuthor, err := server.store.UpdateAuthor(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, updateAuthor)
}

type DeleteReq struct {
	Id int64 `uri:"id" binding:"required"`
}

func (server *Server) DeleteAuthorById(ctx *gin.Context) {
	var req DeleteReq
	isValid := CheckAuthAndExit(ctx)
	if !isValid {
		return
	}
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}
	err := server.store.DeleteAuthor(ctx, req.Id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "success",
	})
}
