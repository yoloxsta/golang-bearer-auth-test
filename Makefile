.PHONY: help build up down restart logs clean test

help: ## Show this help
	@echo "Available commands:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}'

build: ## Build Docker images
	docker-compose build

up: ## Start all services
	docker-compose up -d
	@echo ""
	@echo "✅ Services started!"
	@echo "📡 API Server: http://localhost:8080"
	@echo "🎨 Frontend UI: http://localhost:3000"
	@echo ""
	@echo "Run 'make logs' to see logs"
	@echo "Run 'make down' to stop"

down: ## Stop all services
	docker-compose down

restart: ## Restart all services
	docker-compose restart

logs: ## Show logs
	docker-compose logs -f

logs-api: ## Show API logs only
	docker-compose logs -f api

logs-frontend: ## Show frontend logs only
	docker-compose logs -f frontend

ps: ## Show running containers
	docker-compose ps

clean: ## Remove containers, images, and volumes
	docker-compose down -v --rmi all

test: ## Test the API
	@echo "Testing API endpoints..."
	@curl -s http://localhost:8080/health | jq .
	@echo ""
	@curl -s -H "Authorization: Bearer secret_token_12345" http://localhost:8080/users/1 | jq .

shell-api: ## Open shell in API container
	docker-compose exec api sh

shell-frontend: ## Open shell in frontend container
	docker-compose exec frontend sh
