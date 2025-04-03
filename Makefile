gen:
	go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen -config ./config.yaml ./spec.yaml
db-up:
	migrate --path db/migrations --database 'postgres://postgres:postgres@localhost:5432/api?sslmode=disable' -verbose up
db-down:
	migrate --path db/migrations --database 'postgres://postgres:postgres@localhost:5432/api?sslmode=disable' -verbose down
