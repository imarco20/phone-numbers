run:
	go run ./cmd/api

build:
	CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o ./bin/app ./cmd/api

test:
	go test -v ./...

build-image:
	docker build -t phone-numbers:latest .

run-container:
	docker run -d --rm -p 8080:8080 --name phone-numbers-container phone-numbers