.PHONY: install frontend backend

install:
	cd frontend && npm i
	cd backend && go mod tidy

frontend:
	cd frontend && npm run start

backend:
	cd backend && go run cmd/app/main.go
