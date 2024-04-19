package api

import (
	"errors"
	"net/http"
	"time"

	db "github.com/amIToan/book_library/db/gen_sqlc"
	"github.com/amIToan/book_library/util"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

type CreatedUserParams struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
	Email    string `json:"email" binding:"required,email"`
	Role     string `json:"role" binding:"required,role_force"`
}

type UserRes struct {
	Username          string    `json:"username"`
	Email             string    `json:"email" binding:"required,email"`
	Role              string    `json:"role" binding:"required,role_force"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
}

func (server *Server) CreateUser(ctx *gin.Context) {
	isValid := CheckAuthAndExit(ctx)
	if !isValid {
		return
	}
	var req CreatedUserParams
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}
	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
		return
	}
	var pgText pgtype.Text
	_ = pgText.Scan(req.Role)
	argument := db.CreateUserParams{
		Username:       req.Username,
		HashedPassword: hashedPassword,
		Email:          req.Email,
		Role:           pgText,
	}
	createdUser, err := server.store.CreateUser(ctx, argument)
	if err != nil {
		if util.ErrorCode(err) == util.UniqueViolation {
			ctx.JSON(http.StatusForbidden, util.ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, UserRes{
		Username:          createdUser.Username,
		Email:             createdUser.Email,
		Role:              createdUser.Role.String,
		PasswordChangedAt: createdUser.PasswordChangedAt,
		CreatedAt:         createdUser.CreatedAt,
	})
}

type LoginUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
}
type LoginUserResponse struct {
	AccessToken          string    `json:"access_token"`
	AccessTokenExpiresAt time.Time `json:"access_token_expires_at"`
	User                 UserRes   `json:"user"`
}

func (server *Server) Login(ctx *gin.Context) {
	var req LoginUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}
	user, err := server.store.GetUser(ctx, req.Username)
	if err != nil {
		if errors.Is(err, util.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, util.ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
		return
	}
	err = util.CheckPassword(req.Password, user.HashedPassword)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, util.ErrorResponse(err))
		return
	}
	accessToken, accessPayload, err := server.tokenMaker.CreateToken(
		user.Username,
		user.Role.String,
		server.config.AccessTokenDuration,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
		return
	}
	rsp := LoginUserResponse{
		AccessToken:          accessToken,
		AccessTokenExpiresAt: accessPayload.ExpiredAt,
		User: UserRes{
			Username:          user.Username,
			Email:             user.Email,
			Role:              user.Role.String,
			PasswordChangedAt: user.PasswordChangedAt,
			CreatedAt:         user.CreatedAt,
		},
	}
	ctx.JSON(http.StatusOK, rsp)
}

type UpdateUserReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (server *Server) UpdateUser(ctx *gin.Context) {
	var req UpdateUserReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}
	isValid := CheckAuthAndExit(ctx)
	if !isValid {
		return
	}
	arg := db.UpdateUserParams{
		Username: req.Username,
	}
	if req.Password != "" {
		hashedPass, err := util.HashPassword(req.Password)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
			return
		}
		arg.HashedPassword = pgtype.Text{
			String: hashedPass,
			Valid:  true,
		}
	}
	if req.Email != "" {
		arg.Email = pgtype.Text{
			String: req.Email,
			Valid:  true,
		}
	}
	updatedUser, err := server.store.UpdateUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, UserRes{
		Username:          updatedUser.Username,
		Email:             updatedUser.Email,
		Role:              updatedUser.Role.String,
		PasswordChangedAt: updatedUser.PasswordChangedAt,
		CreatedAt:         updatedUser.CreatedAt,
	})
}

/* func (server *Server) CreateAdmin(ctx *gin.Context) {

	var req CreatedUserParams
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}
	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
		return
	}
	var pgText pgtype.Text
	_ = pgText.Scan(req.Role)
	argument := db.CreateUserParams{
		Username:       req.Username,
		HashedPassword: hashedPassword,
		Email:          req.Email,
		Role:           pgText,
	}
	createdUser, err := server.store.CreateUser(ctx, argument)
	if err != nil {
		if util.ErrorCode(err) == util.UniqueViolation {
			ctx.JSON(http.StatusForbidden, util.ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, UserResult{
		Username:          createdUser.Username,
		Email:             createdUser.Email,
		Role:              createdUser.Role.String,
		PasswordChangedAt: createdUser.PasswordChangedAt,
		CreatedAt:         createdUser.CreatedAt,
	})
} */
