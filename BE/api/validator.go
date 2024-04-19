package api

import (
	"errors"
	"net/http"

	"github.com/amIToan/book_library/token"
	"github.com/amIToan/book_library/util"
	"github.com/gin-gonic/gin"
)

func CheckAuthAndExit(ctx *gin.Context) bool {
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if authPayload.Role != util.Admin {
		err := errors.New("NO AUTHORIZATION")
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return false
	}
	return true
}
