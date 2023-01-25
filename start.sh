#!/bin/sh
set -e

echo "run db migration"
/app/migrate -path /app/db/migration -database "$DB_SOURCE" -verbose up

# migrate -path db/migration -database "postgresql://postgres:secret@postgres:5432/video-manager?sslmode=disable" -verbose up

echo "start the app"
exec "$@"