COMPOSE := docker-compose
default:
	$(COMPOSE) up --build

up:
	$(COMPOSE) up

down:
	$(COMPOSE) down

api:
	$(COMPOSE) up api

db:
	$(COMPOSE) up db

db-d:
	$(COMPOSE) up -d db

mysql:
	$(COMPOSE) exec db mysql -uroot -p db -pPASSWORD 

build:
	$(COMPOSE) build

.PHONY: default api db build
