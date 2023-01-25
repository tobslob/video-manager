postgres:
		docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=kazeem -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
		docker exec -it postgres12 createdb --username=kazeem --owner=kazeem video-manager

dropdb:
		docker exec -it postgres12 dropdb --username=kazeem video-manager

migrateup:
		migrate -path db/migration -database "postgresql://kazeem:secret@0.0.0.0:5432/video-manager?sslmode=disable" -verbose up

migratedown:
		migrate -path db/migration -database "postgresql://postgres:secret@0.0.0.0:5432/video-manager?sslmode=disable" -verbose down

sqlc:
		sqlc generate

test:
		go test -v -cover ./...
	
server:
		go run main.go

mock:
		mockgen -package mockdb -destination db/mock/store.go github.com/tobslob/video-manager/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server mock