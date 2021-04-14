test:
	go test ./...
run:
	go run main.go
docker-build:
	docker build -t message-publisher .
docker-run:
	docker run -d -p 3001:3001 --name app message-publisher
docker-clean:
	docker rm -f app