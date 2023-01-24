 -- name: CreateAnnotation :one
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
RETURNING *;

-- name: GetAnnotation :one
SELECT * FROM annotations
WHERE video_id = $1 AND user_id = $2 LIMIT 1;

-- name: UpdateAnnotation :one
UPDATE annotations
SET note = $3,
  title = $4,
  label = $5,
  pause = $6,
  start_time =$7,
  end_time =$8,
  type = $9
WHERE video_id = $1 AND user_id = $2 RETURNING *;

-- name: ListAnnotations :many
SELECT * FROM annotations
WHERE video_id = $1 AND user_id = $2
ORDER BY id
LIMIT $3
OFFSET $4;

-- name: DeleteAnnotation :exec
DELETE FROM annotations WHERE video_id = $1 AND user_id = $2;
