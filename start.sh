set -e

echo "run db migration"
/app/migration -path /app/migration -databasee "$DB_SOURCE" -verbose up

echo "start the app"
exec "$@"