.PHONY: run, dc

dc:
	docker-compose up --remove-orphans --build

run:
	go run cmd/main.go