ifndef ENV
ENV := development
endif

ifeq ($(filter $(ENV),test debug development staging production),)
$(error The ENV variable is invalid.)
endif

ifeq ($(ENV), production)
COMPOSE_FILE_PATH := -f docker/docker-compose.yml -f docker/docker-compose.prod.yml
else
COMPOSE_FILE_PATH := -f docker/docker-compose.yml
endif

ifeq ($(ENV), debug)
FILE := debug
else
FILE := dev
endif

ACTION := APP_ENV=$(ENV) FILE=$(FILE) docker-compose $(COMPOSE_FILE_PATH)


build:
	$(info Make: Building "$(ENV)" environment images.)
	@$(ACTION) build --no-cache

start:
	$(info Make: Starting "$(ENV)" environment containers.)
	@$(ACTION) up

start-build:
	$(info Make: Starting "$(ENV)" environment containers.)
	@$(ACTION) up --build

stop:
	$(info Make: Stopping "$(ENV)" environment containers.)
	@$(ACTION) stop

clean:
	@docker system prune --volumes --force
