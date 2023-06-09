up:
	docker compose up -d && docker exec -it receiver-management-api bash -c "go run -v ./database/seed/"

seed:
	docker exec -it receiver-management-api bash -c "go run -v ./database/seed/"

run:
	docker exec -it receiver-management-api bash -c "go run -v ./cmd/server/"

test:
	go test ./... -v --short

app:
	docker exec -it receiver-management-api bash

test-integration:
	docker exec -it receiver-management-api bash -c "go test ./... -v --run Integration" && make seed

db:
	docker exec -it receiver-management-db bash -c "psql -d receiver-management -U app"

unit-test-cover:
	go test ./... -v --short -coverprofile coverage.out

cover-report:
	go tool cover -html=coverage.out