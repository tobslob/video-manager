package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	db "github.com/tobslob/video-manager/db/sqlc"
	"github.com/tobslob/video-manager/token"
	"github.com/tobslob/video-manager/utils"
)

type VideoAndMetadataRequest struct {
	Url         string         `json:"url" binding:"required"`
	Duration    string         `json:"duration" binding:"required"`
	Title       string         `json:"title" binding:"required"`
	Width       int32          `json:"width" binding:"required"`
	Height      int32          `json:"height" binding:"required"`
	FileType    string         `json:"file_type" binding:"required"`
	FileSize    sql.NullString `json:"file_size"`
	Resolutions int32          `json:"resolutions"`
	Keywords    sql.NullString `json:"keywords"`
}

func (server *Server) createVideoWithMetadata(ctx *gin.Context) {
	var req VideoAndMetadataRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	duration := utils.ParseTimeToStringRepresentation(req.Duration)

	arg := db.CreateVideoWithMetadata{
		CreateVideoParams: db.CreateVideoParams{
			Url:      req.Url,
			UserID:   authPayload.UserID,
			Duration: duration,
			Title:    req.Title,
		},
		CreateMetadataParams: db.CreateMetadataParams{
			Width:       req.Width,
			Height:      req.Height,
			FileType:    req.FileType,
			FileSize:    req.FileSize,
			Resolutions: req.Resolutions,
			Keywords:    req.Keywords,
		},
	}

	result, err := server.store.CreatVideoWithMetadataTx(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, result)
}

type GetVideoRequest struct {
	Id string `uri:"id" binding:"required"`
}

func (server *Server) getVideoWithMetadata(ctx *gin.Context) {
	var req GetVideoRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	id, _ := uuid.Parse(req.Id)
	result, err := server.store.GetAVideoAndMetadata(ctx, db.GetAVideoAndMetadataParams{
		ID: id, UserID: authPayload.UserID,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, result)
}
