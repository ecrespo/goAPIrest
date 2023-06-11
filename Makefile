# Variables
DOCKER_COMPOSE_FILE_TEST = docker-compose.test.yaml
DOCKER_COMPOSE_FILE = docker-compose.yaml
DOCKER_COMPOSE = docker compose -f $(DOCKER_COMPOSE_FILE)
DOCKER_COMPOSE_TEST = docker compose -f $(DOCKER_COMPOSE_FILE_TEST)

# Levantar Docker y ejecutar pruebas unitarias

build:
	@$(DOCKER_COMPOSE) build

up:
	@$(DOCKER_COMPOSE) up --build

ps:
	@$(DOCKER_COMPOSE) ps

logs:
	@$(DOCKER_COMPOSE) logs -f

unittest:
	@$(DOCKER_COMPOSE_TEST) up --build  --rm unittest --abort-on-container-exit

# Borrar contenedores e imágenes
clean:
	@$(DOCKER_COMPOSE) down --volumes --remove-orphans

purge: clean
	@docker system prune -a --volumes --force

# Ejecutar make clean y levantar Docker
restart: clean up