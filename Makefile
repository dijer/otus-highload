.PHONY: install frontend backend start stop restart

install:
	cd frontend && npm i
	cd backend && go mod tidy

frontend:
	cd frontend && npm run start

backend:
	cd backend && go run cmd/app/main.go

start:
	docker-compose up --build

stop:
	docker-compose down -v

restart:
	$(MAKE) stop 
	$(MAKE) start
