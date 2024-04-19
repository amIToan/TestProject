package api

import (
	"net/http"

	"github.com/amIToan/book_library/util"
	"github.com/gin-gonic/gin"
)

type GetBookGenresReq struct {
	BookId int64 `uri:"book_id" binding:"required"`
}

func (server *Server) GetBookGenres(ctx *gin.Context) {
	var req GetBookGenresReq
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}
	data, err := server.store.GetBookGenres(ctx, req.BookId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, data)
}
