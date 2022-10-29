build:
	CGO_ENABLED=0 GOOS=linux

run: build
	docker-compose up --build