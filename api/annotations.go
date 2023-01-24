package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lib/pq"
	db "github.com/tobslob/video-manager/db/sqlc"
	"github.com/tobslob/video-manager/token"
	"github.com/tobslob/video-manager/utils"
)

type createAnnotationRequest struct {
	VideoID   uuid.UUID `json:"video_id" binding:"required"`
	Type      string    `json:"type" binding:"required"`
	Note      string    `json:"note" binding:"required"`
	Title     string    `json:"title" binding:"required,max=150"`
	Label     string    `json:"label" binding:"required,max=50"`
	Pause     bool      `json:"pause"`
	StartTime string    `json:"start_time" binding:"required"`
	EndTime   string    `json:"end_time" binding:"required"`
}

func (server *Server) createAnnotation(ctx *gin.Context) {
	var req createAnnotationRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	video, err := server.store.GetVideo(ctx, db.GetVideoParams{ID: req.VideoID, UserID: authPayload.UserID})
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	endTime := utils.ParseTimeToStringRepresentation(req.EndTime)
	startTime := utils.ParseTimeToStringRepresentation(req.StartTime)

	if !utils.ValidateTime(video.Duration, endTime) {
		err := fmt.Errorf("anotation is out of bounds of video duration")
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	id, _ := uuid.NewRandom()

	arg := db.CreateAnnotationParams{
		ID:        id,
		VideoID:   req.VideoID,
		UserID:    authPayload.UserID,
		Type:      req.Type,
		Note:      req.Note,
		Title:     req.Title,
		Label:     req.Label,
		Pause:     req.Pause,
		StartTime: startTime,
		EndTime:   endTime,
	}

	annotation, err := server.store.CreateAnnotation(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "foreign_key_violation", "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, annotation)
}

type getAnnotationRequest struct {
	VideoID string `uri:"video_id" binding:"required"`
}

func (server *Server) getAnnotation(ctx *gin.Context) {
	var req getAnnotationRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	id, _ := uuid.Parse(req.VideoID)
	annotation, err := server.store.GetAnnotation(ctx, db.GetAnnotationParams{
		VideoID: id, UserID: authPayload.UserID,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, annotation)
}

type listAnnotationsRequest struct {
	VideoID  string `form:"video_id" binding:"required"`
	PageID   int32  `form:"page_id" binding:"required,min=1"`
	PageSize int32  `form:"page_size" binding:"required,min=1"`
}

func (server *Server) listAnnotations(ctx *gin.Context) {
	var req listAnnotationsRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	id, _ := uuid.Parse(req.VideoID)
	arg := db.ListAnnotationsParams{
		VideoID: id,
		UserID:  authPayload.UserID,
		Limit:   req.PageSize,
		Offset:  (req.PageID - 1) * req.PageSize,
	}

	accounts, err := server.store.ListAnnotations(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, accounts)
}

func (server *Server) deleteAnnotation(ctx *gin.Context) {
	var req getAnnotationRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	id, _ := uuid.Parse(req.VideoID)
	deleteArg := db.DeleteAnnotationParams{VideoID: id, UserID: authPayload.UserID}

	err := server.store.DeleteAnnotation(ctx, deleteArg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

type UpdateAnnotationRequest struct {
	VideoID   string `json:"video_id"`
	Note      string `json:"note"`
	Title     string `json:"title" binding:"max=150"`
	Label     string `json:"label" binding:"max=50"`
	Pause     bool   `json:"pause"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	Type      string `json:"type"`
}

func (server *Server) updateAnnotation(ctx *gin.Context) {
	var req UpdateAnnotationRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	id, _ := uuid.Parse(req.VideoID)
	updateArg := db.UpdateAnnotationParams{
		VideoID:   id,
		UserID:    authPayload.UserID,
		Note:      req.Note,
		Title:     req.Title,
		Label:     req.Label,
		Pause:     req.Pause,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
		Type:      req.Type,
	}

	updateArg.EndTime = utils.ParseTimeToStringRepresentation(req.EndTime)
	updateArg.StartTime = utils.ParseTimeToStringRepresentation(req.StartTime)

	video, err := server.store.GetVideo(ctx, db.GetVideoParams{ID: id, UserID: authPayload.UserID})
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if !utils.ValidateTime(video.Duration, updateArg.EndTime) {
		err := fmt.Errorf("anotation is out of bounds of video duration")
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	annotion, err := server.store.UpdateAnnotation(ctx, updateArg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, annotion)
}
