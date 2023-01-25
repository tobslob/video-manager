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
WHERE id = $1 AND user_id = $2 LIMIT 1;

-- name: UpdateAnnotation :one
UPDATE annotations
SET note = $4,
  title = $5,
  label = $6,
  pause = $7,
  start_time =$8,
  end_time =$9,
  type = $10
WHERE id = $1 AND user_id = $2 AND video_id = $3 RETURNING *;

-- name: ListAnnotations :many
SELECT * FROM annotations
WHERE video_id = $1 AND user_id = $2
ORDER BY id
LIMIT $3
OFFSET $4;

-- name: DeleteAnnotations :exec
DELETE FROM annotations WHERE video_id = $1 AND user_id = $2;

-- name: DeleteAnnotation :exec
DELETE FROM annotations WHERE id = $1 AND user_id = $2;
