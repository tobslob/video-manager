// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: videos.sql

package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const createVideo = `-- name: CreateVideo :one
INSERT INTO videos (
  id,
  url,
  user_id,
  duration,
  title
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING id, url, user_id, duration, title, created_at, updated_at
`

type CreateVideoParams struct {
	ID       uuid.UUID `json:"id"`
	Url      string    `json:"url"`
	UserID   uuid.UUID `json:"user_id"`
	Duration string    `json:"duration"`
	Title    string    `json:"title"`
}

func (q *Queries) CreateVideo(ctx context.Context, arg CreateVideoParams) (Video, error) {
	row := q.queryRow(ctx, q.createVideoStmt, createVideo,
		arg.ID,
		arg.Url,
		arg.UserID,
		arg.Duration,
		arg.Title,
	)
	var i Video
	err := row.Scan(
		&i.ID,
		&i.Url,
		&i.UserID,
		&i.Duration,
		&i.Title,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteVideo = `-- name: DeleteVideo :exec
DELETE FROM videos WHERE user_id = $1 AND id = $2
`

type DeleteVideoParams struct {
	UserID uuid.UUID `json:"user_id"`
	ID     uuid.UUID `json:"id"`
}

func (q *Queries) DeleteVideo(ctx context.Context, arg DeleteVideoParams) error {
	_, err := q.exec(ctx, q.deleteVideoStmt, deleteVideo, arg.UserID, arg.ID)
	return err
}

const getAVideoAndMetadata = `-- name: GetAVideoAndMetadata :one
SELECT v.id, url, user_id, duration, title, v.created_at, v.updated_at, m.id, video_id, width, height, file_type, file_size, last_modify, accessed_date, resolutions, keywords, m.created_at, m.updated_at FROM videos V
INNER JOIN metadatas M ON V.id = M.video_id WHERE V.id = $1 AND V.user_id = $2
LIMIT 1
`

type GetAVideoAndMetadataParams struct {
	ID     uuid.UUID `json:"id"`
	UserID uuid.UUID `json:"user_id"`
}

type GetAVideoAndMetadataRow struct {
	ID           uuid.UUID      `json:"id"`
	Url          string         `json:"url"`
	UserID       uuid.UUID      `json:"user_id"`
	Duration     string         `json:"duration"`
	Title        string         `json:"title"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	ID_2         uuid.UUID      `json:"id_2"`
	VideoID      uuid.UUID      `json:"video_id"`
	Width        int32          `json:"width"`
	Height       int32          `json:"height"`
	FileType     string         `json:"file_type"`
	FileSize     sql.NullString `json:"file_size"`
	LastModify   time.Time      `json:"last_modify"`
	AccessedDate time.Time      `json:"accessed_date"`
	Resolutions  int32          `json:"resolutions"`
	Keywords     string         `json:"keywords"`
	CreatedAt_2  time.Time      `json:"created_at_2"`
	UpdatedAt_2  time.Time      `json:"updated_at_2"`
}

func (q *Queries) GetAVideoAndMetadata(ctx context.Context, arg GetAVideoAndMetadataParams) (GetAVideoAndMetadataRow, error) {
	row := q.queryRow(ctx, q.getAVideoAndMetadataStmt, getAVideoAndMetadata, arg.ID, arg.UserID)
	var i GetAVideoAndMetadataRow
	err := row.Scan(
		&i.ID,
		&i.Url,
		&i.UserID,
		&i.Duration,
		&i.Title,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.ID_2,
		&i.VideoID,
		&i.Width,
		&i.Height,
		&i.FileType,
		&i.FileSize,
		&i.LastModify,
		&i.AccessedDate,
		&i.Resolutions,
		&i.Keywords,
		&i.CreatedAt_2,
		&i.UpdatedAt_2,
	)
	return i, err
}

const getVideo = `-- name: GetVideo :one
SELECT id, url, user_id, duration, title, created_at, updated_at FROM videos WHERE id = $1 AND user_id = $2
LIMIT 1
`

type GetVideoParams struct {
	ID     uuid.UUID `json:"id"`
	UserID uuid.UUID `json:"user_id"`
}

func (q *Queries) GetVideo(ctx context.Context, arg GetVideoParams) (Video, error) {
	row := q.queryRow(ctx, q.getVideoStmt, getVideo, arg.ID, arg.UserID)
	var i Video
	err := row.Scan(
		&i.ID,
		&i.Url,
		&i.UserID,
		&i.Duration,
		&i.Title,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listVideos = `-- name: ListVideos :many
SELECT v.id, url, user_id, duration, title, v.created_at, v.updated_at, m.id, video_id, width, height, file_type, file_size, last_modify, accessed_date, resolutions, keywords, m.created_at, m.updated_at FROM videos V INNER JOIN metadatas M ON V.id = M.video_id
WHERE user_id = $1
ORDER BY id
LIMIT $2
OFFSET $3
`

type ListVideosParams struct {
	UserID uuid.UUID `json:"user_id"`
	Limit  int32     `json:"limit"`
	Offset int32     `json:"offset"`
}

type ListVideosRow struct {
	ID           uuid.UUID      `json:"id"`
	Url          string         `json:"url"`
	UserID       uuid.UUID      `json:"user_id"`
	Duration     string         `json:"duration"`
	Title        string         `json:"title"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	ID_2         uuid.UUID      `json:"id_2"`
	VideoID      uuid.UUID      `json:"video_id"`
	Width        int32          `json:"width"`
	Height       int32          `json:"height"`
	FileType     string         `json:"file_type"`
	FileSize     sql.NullString `json:"file_size"`
	LastModify   time.Time      `json:"last_modify"`
	AccessedDate time.Time      `json:"accessed_date"`
	Resolutions  int32          `json:"resolutions"`
	Keywords     string         `json:"keywords"`
	CreatedAt_2  time.Time      `json:"created_at_2"`
	UpdatedAt_2  time.Time      `json:"updated_at_2"`
}

func (q *Queries) ListVideos(ctx context.Context, arg ListVideosParams) ([]ListVideosRow, error) {
	rows, err := q.query(ctx, q.listVideosStmt, listVideos, arg.UserID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListVideosRow{}
	for rows.Next() {
		var i ListVideosRow
		if err := rows.Scan(
			&i.ID,
			&i.Url,
			&i.UserID,
			&i.Duration,
			&i.Title,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.ID_2,
			&i.VideoID,
			&i.Width,
			&i.Height,
			&i.FileType,
			&i.FileSize,
			&i.LastModify,
			&i.AccessedDate,
			&i.Resolutions,
			&i.Keywords,
			&i.CreatedAt_2,
			&i.UpdatedAt_2,
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

const updateVideo = `-- name: UpdateVideo :one
UPDATE videos
SET title = $2
WHERE id = $1
RETURNING id, url, user_id, duration, title, created_at, updated_at
`

type UpdateVideoParams struct {
	ID    uuid.UUID `json:"id"`
	Title string    `json:"title"`
}

func (q *Queries) UpdateVideo(ctx context.Context, arg UpdateVideoParams) (Video, error) {
	row := q.queryRow(ctx, q.updateVideoStmt, updateVideo, arg.ID, arg.Title)
	var i Video
	err := row.Scan(
		&i.ID,
		&i.Url,
		&i.UserID,
		&i.Duration,
		&i.Title,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
