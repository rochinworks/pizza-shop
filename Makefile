DOCKER=docker
DOCKER_COMPOSE=docker-compose

default: build

build:
	docker build -t pizza-shop .
	docker run -it pizza-shop

# starts the application in watch mode
start:
	$(DOCKER_COMPOSE) up -d --build
	$(DOCKER_COMPOSE) exec app gin -i --path=src --bin=app

start-dev:
	$(DOCKER_COMPOSE) docker-compose.yml up -d
	HOST="localhost" gin -i run src/*.go
