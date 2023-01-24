package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	db "github.com/tobslob/video-manager/db/sqlc"
	"github.com/tobslob/video-manager/token"
	"github.com/tobslob/video-manager/utils"
)

type Server struct {
	config     utils.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

// NewServer creates a new HTTP server and setup routing.
func NewServer(config utils.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewJWTMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	server.setupRouter()
	return server, nil
}

// Start runs the HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

	authRoutes.POST("/videos", server.createVideoWithMetadata)
	authRoutes.GET("/video/:id", server.getVideoWithMetadata)

	authRoutes.POST("/annotation", server.createAnnotation)
	authRoutes.GET("/annotation/:video_id", server.getAnnotation)
	authRoutes.GET("/annotations", server.listAnnotations)
	authRoutes.DELETE("/annotation/:video_id", server.deleteAnnotation)
	authRoutes.PATCH("/annotation", server.updateAnnotation)

	server.router = router
}
