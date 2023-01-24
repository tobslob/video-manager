 -- name: CreateVideo :one
INSERT INTO videos (
  id,
  url,
  user_id,
  duration,
  title
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;

-- name: GetVideo :one
SELECT * FROM videos WHERE id = $1 AND user_id = $2
LIMIT 1;

-- name: GetAVideoAndMetadata :one
SELECT * FROM videos V
INNER JOIN metadatas M ON V.id = M.video_id WHERE V.id = $1 AND V.user_id = $2
LIMIT 1;

-- name: ListVideos :many
SELECT * FROM videos V INNER JOIN metadatas M ON V.id = M.video_id
WHERE user_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: DeleteVideo :exec
DELETE FROM videos WHERE user_id = $1 AND id = $2;

-- name: UpdateVideo :one
UPDATE videos
SET title = $2
WHERE id = $1
RETURNING *;
