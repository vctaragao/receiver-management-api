run:
	go run -v ./cmd/server/

test:
	go test ./... -v

container:
	docker exec -it receiver-management-api bash