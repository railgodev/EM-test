run:
	@cd backend && \
	POSTGRES_HOST=localhost go run ./cmd/main.go

up:
	@docker compose --env-file ./backend/.env up -d --build

up-fg:
	@docker compose --env-file ./backend/.env up --build

up-postgres:
	@docker compose --env-file ./backend/.env up -d --build postgres

up-postgres-fg:
	@docker compose --env-file ./backend/.env up --build postgres

down:
	@docker compose --env-file ./backend/.env down

stop:
	@docker compose --env-file ./backend/.env stop

swagger:
	@swagger generate server -A SubscriptionService -f ./swagger.yaml --exclude-main