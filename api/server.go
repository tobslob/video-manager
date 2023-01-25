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

	unsecureRoutes := router.Group("/api/v1")

	unsecureRoutes.POST("/users", server.createUser)
	unsecureRoutes.POST("/users/login", server.loginUser)

	authRoutes := router.Group("/api/v1").Use(authMiddleware(server.tokenMaker))

	authRoutes.POST("/videos", server.createVideoWithMetadata)
	authRoutes.GET("/videos/:id", server.getVideoWithMetadata)
	authRoutes.DELETE("/videos/:id", server.deleteVideo)

	authRoutes.POST("/annotations", server.createAnnotation)
	authRoutes.GET("/annotations/:id", server.getAnnotation)
	authRoutes.GET("/annotations", server.listAnnotations)
	authRoutes.DELETE("/annotations/:id", server.deleteAnnotation)
	authRoutes.PATCH("/annotations/:id/:video_id", server.updateAnnotation)

	server.router = router
}
