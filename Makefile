DIR=$(shell pwd)
DATABASE="postgres://postgres:REWQ_7AD83439wEqwR@localhost:5430/bets?sslmode=disable"

.PHONY: docker-run
docker-run:
	docker compose up --build

api_back:
	@go run $(DIR)/cmd/api_back/main.go

api_back_watch:
	@nodemon --exec 'go run $(DIR)/cmd/api_back/main.go' --signal SIGTERM -e go,conf,yml,json,mod --watch

tg_worker:
	@go run $(DIR)/cmd/tg_worker/main.go

go_mod_download:
	@echo "  >  Checking if there is any missing dependencies..."
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go mod download

go_mod_tidy:
	@echo "  >  Checking if there is any missing dependencies..."
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go mod tidy

.PHONY: migrations
migrations:
	@migrate -path $(DIR)/migrations -database $(DATABASE) up

migrations-down:
	@migrate -path $(DIR)/migrations -database $(DATABASE) down

postgres_db:
	docker run --name=wetao_db_postgres -e POSTGRES_PASSWORD="REWQ_7AD83439wEqwR" -e POSTGRES_DB=wetao_db -p 5477:5432 -d  postgres:16.0
#--rm (Удалять после закрытия)

redis_db:
	docker run --name=gd_back_redis -p 6379:6379 -d --rm redis:7.0.4
