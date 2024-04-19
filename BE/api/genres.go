package api

import (
	"net/http"

	db "github.com/amIToan/book_library/db/gen_sqlc"
	"github.com/amIToan/book_library/util"
	"github.com/gin-gonic/gin"
)

type CreateGenreReq struct {
	Name string `json:"name" binding:"required"`
}

func (server *Server) CreateGenre(ctx *gin.Context) {
	var req CreateGenreReq
	isValid := CheckAuthAndExit(ctx)
	//log.Debug().Msgf("check why invalid %v", isValid)
	if !isValid {
		return
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}
	newGenre, err := server.store.CreateGenre(ctx, req.Name)
	if err != nil {
		if util.ErrorCode(err) == util.UniqueViolation || util.ErrorCode(err) == util.ForeignKeyViolation {
			ctx.JSON(http.StatusForbidden, util.ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, newGenre)
}
func (server *Server) GetGenres(ctx *gin.Context) {
	isValid := CheckAuthAndExit(ctx)
	if !isValid {
		return
	}
	genres, err := server.store.GetGenres(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, genres)
}

type UpdateGenreReq struct {
	GenreID int64  `json:"genre_id" binding:"required,number"`
	Name    string `json:"name" binding:"required"`
}

func (server *Server) UpdateGenre(ctx *gin.Context) {
	isValid := CheckAuthAndExit(ctx)
	if !isValid {
		return
	}
	var req UpdateGenreReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}
	updateGenre, err := server.store.UpdateGenre(ctx, db.UpdateGenreParams(req))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, updateGenre)
}
