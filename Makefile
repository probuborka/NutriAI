# See: https://www.gnu.org/software/make/

# Переменные.
DOCKER_COMPOSE := docker compose

# Сборка Docker-образа.
.PHONY: build run-local
build:
	@echo "Building Docker image for the application..."
	$(DOCKER_COMPOSE) up --build

# Запуска сервисов с использованием docker-compose.
run-local:
	@echo "Starting services with docker-compose..."
	$(DOCKER_COMPOSE) up 

# Задача для остановки и удаления контейнеров.
down:
	@echo "Stopping and removing services..."
	$(DOCKER_COMPOSE) down

# Задача для перезапуска контейнеров.
restart:
	@echo "Restarting services..."
	$(DOCKER_COMPOSE) down && $(DOCKER_COMPOSE) up --build

# Задача для очистки неиспользуемых Docker-ресурсов.
prune:
	@echo "Pruning unused Docker resources (images, containers, volumes, networks)..."
	docker system prune -af	