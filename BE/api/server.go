package api

import (
	"context"
	"fmt"
	"net/http"

	db "github.com/amIToan/book_library/db/gen_sqlc"
	"github.com/amIToan/book_library/token"
	"github.com/amIToan/book_library/util"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	config     util.Config
	store      db.StorageInterface
	tokenMaker token.Maker
	router     *gin.Engine
	srv        *http.Server
}

// NewServer creates a new HTTP server and set up routing.
func NewServer(config util.Config, store db.StorageInterface) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("role_force", util.ValidateRole)
		v.RegisterValidation("AtLeastOneFile", util.AtLeastOneFile)
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {

	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Hello")
	})

	//router.POST("/admin", server.CreateAdmin)
	router.Static("/assets", "./assets")
	router.POST("/users/login", server.Login)
	router.POST("/users", server.CreateUser)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))
	//users
	authRoutes.POST("/users/update", server.UpdateUser)
	//authors
	authRoutes.POST("/author", server.CreateAuthor)
	authRoutes.GET("/authors", server.GetFiftyAuthors)
	authRoutes.GET("/authors/:author_name", server.GetAuthorsByName)
	authRoutes.PUT("/authors/update", server.UpdateAuthor)
	authRoutes.DELETE("/authors/:id", server.DeleteAuthorById)
	//genres
	authRoutes.POST("/genre", server.CreateGenre)
	authRoutes.GET("/genres", server.GetGenres)
	authRoutes.PUT("/genres/update", server.UpdateGenre)
	//book
	authRoutes.POST("/book", server.CreateBook)
	authRoutes.GET("/book/:id", server.GetBookById)
	authRoutes.GET("/books", server.GetBooks)
	authRoutes.PUT("/book", server.UpdateBook)
	authRoutes.DELETE("/book/delete/:id", server.DeleteBooks)
	authRoutes.GET("books/title", server.GetBooksByTitle)
	//book_genres
	authRoutes.GET("/books_genres/:book_id", server.GetBookGenres)

	server.router = router
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	srv := &http.Server{
		Addr:    address,
		Handler: server.router,
	}
	server.srv = srv
	return srv.ListenAndServe()
}

func (server *Server) GraceFullyShutdown(ctx context.Context) error {
	return server.srv.Shutdown(ctx)
}
