// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: annotations.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const createAnnotation = `-- name: CreateAnnotation :one
INSERT INTO annotations (
  id,
  video_id,
  type,
  note,
  title,
  label,
  pause,
  start_time,
  end_time,
  user_id
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10
)
RETURNING id, video_id, user_id, type, note, title, label, pause, start_time, end_time, created_at, updated_at
`

type CreateAnnotationParams struct {
	ID        uuid.UUID `json:"id"`
	VideoID   uuid.UUID `json:"video_id"`
	Type      string    `json:"type"`
	Note      string    `json:"note"`
	Title     string    `json:"title"`
	Label     string    `json:"label"`
	Pause     bool      `json:"pause"`
	StartTime string    `json:"start_time"`
	EndTime   string    `json:"end_time"`
	UserID    uuid.UUID `json:"user_id"`
}

func (q *Queries) CreateAnnotation(ctx context.Context, arg CreateAnnotationParams) (Annotation, error) {
	row := q.queryRow(ctx, q.createAnnotationStmt, createAnnotation,
		arg.ID,
		arg.VideoID,
		arg.Type,
		arg.Note,
		arg.Title,
		arg.Label,
		arg.Pause,
		arg.StartTime,
		arg.EndTime,
		arg.UserID,
	)
	var i Annotation
	err := row.Scan(
		&i.ID,
		&i.VideoID,
		&i.UserID,
		&i.Type,
		&i.Note,
		&i.Title,
		&i.Label,
		&i.Pause,
		&i.StartTime,
		&i.EndTime,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteAnnotation = `-- name: DeleteAnnotation :exec
DELETE FROM annotations WHERE video_id = $1 AND user_id = $2
`

type DeleteAnnotationParams struct {
	VideoID uuid.UUID `json:"video_id"`
	UserID  uuid.UUID `json:"user_id"`
}

func (q *Queries) DeleteAnnotation(ctx context.Context, arg DeleteAnnotationParams) error {
	_, err := q.exec(ctx, q.deleteAnnotationStmt, deleteAnnotation, arg.VideoID, arg.UserID)
	return err
}

const getAnnotation = `-- name: GetAnnotation :one
SELECT id, video_id, user_id, type, note, title, label, pause, start_time, end_time, created_at, updated_at FROM annotations
WHERE video_id = $1 AND user_id = $2 LIMIT 1
`

type GetAnnotationParams struct {
	VideoID uuid.UUID `json:"video_id"`
	UserID  uuid.UUID `json:"user_id"`
}

func (q *Queries) GetAnnotation(ctx context.Context, arg GetAnnotationParams) (Annotation, error) {
	row := q.queryRow(ctx, q.getAnnotationStmt, getAnnotation, arg.VideoID, arg.UserID)
	var i Annotation
	err := row.Scan(
		&i.ID,
		&i.VideoID,
		&i.UserID,
		&i.Type,
		&i.Note,
		&i.Title,
		&i.Label,
		&i.Pause,
		&i.StartTime,
		&i.EndTime,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listAnnotations = `-- name: ListAnnotations :many
SELECT id, video_id, user_id, type, note, title, label, pause, start_time, end_time, created_at, updated_at FROM annotations
WHERE video_id = $1 AND user_id = $2
ORDER BY id
LIMIT $3
OFFSET $4
`

type ListAnnotationsParams struct {
	VideoID uuid.UUID `json:"video_id"`
	UserID  uuid.UUID `json:"user_id"`
	Limit   int32     `json:"limit"`
	Offset  int32     `json:"offset"`
}

func (q *Queries) ListAnnotations(ctx context.Context, arg ListAnnotationsParams) ([]Annotation, error) {
	rows, err := q.query(ctx, q.listAnnotationsStmt, listAnnotations,
		arg.VideoID,
		arg.UserID,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Annotation{}
	for rows.Next() {
		var i Annotation
		if err := rows.Scan(
			&i.ID,
			&i.VideoID,
			&i.UserID,
			&i.Type,
			&i.Note,
			&i.Title,
			&i.Label,
			&i.Pause,
			&i.StartTime,
			&i.EndTime,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateAnnotation = `-- name: UpdateAnnotation :one
UPDATE annotations
SET note = $3,
  title = $4,
  label = $5,
  pause = $6,
  start_time =$7,
  end_time =$8,
  type = $9
WHERE video_id = $1 AND user_id = $2 RETURNING id, video_id, user_id, type, note, title, label, pause, start_time, end_time, created_at, updated_at
`

type UpdateAnnotationParams struct {
	VideoID   uuid.UUID `json:"video_id"`
	UserID    uuid.UUID `json:"user_id"`
	Note      string    `json:"note"`
	Title     string    `json:"title"`
	Label     string    `json:"label"`
	Pause     bool      `json:"pause"`
	StartTime string    `json:"start_time"`
	EndTime   string    `json:"end_time"`
	Type      string    `json:"type"`
}

func (q *Queries) UpdateAnnotation(ctx context.Context, arg UpdateAnnotationParams) (Annotation, error) {
	row := q.queryRow(ctx, q.updateAnnotationStmt, updateAnnotation,
		arg.VideoID,
		arg.UserID,
		arg.Note,
		arg.Title,
		arg.Label,
		arg.Pause,
		arg.StartTime,
		arg.EndTime,
		arg.Type,
	)
	var i Annotation
	err := row.Scan(
		&i.ID,
		&i.VideoID,
		&i.UserID,
		&i.Type,
		&i.Note,
		&i.Title,
		&i.Label,
		&i.Pause,
		&i.StartTime,
		&i.EndTime,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
