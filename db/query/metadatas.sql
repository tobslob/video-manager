 -- name: CreateMetadata :one
INSERT INTO metadatas (
  id,
  video_id,
  width,
  height,
  file_type,
  file_size,
  last_modify,
  accessed_date,
  resolutions,
  keywords
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10
)
RETURNING *;

-- name: GetMetadata :one
SELECT * FROM metadatas
WHERE video_id = $1 LIMIT 1;

-- name: UpdateMetadata :one
UPDATE metadatas
SET width = $2, height = $3, keywords = $4
WHERE id = $1
RETURNING *;
