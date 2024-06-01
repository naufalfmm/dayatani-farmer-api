SHELL:/bin/bash

DBFILENAME=
DBPATH=
SQLPATH=
ENVFILENAME?=.env
VERSION=
PORT=

db_migrate:
ifeq ($(DOCKER), true)
	docker run --name cryptocurrency-price-api-migration -v $(SQLPATH):/usr/src/migrations/sql -v $(DBPATH)/$(DBFILENAME):/usr/src/$(DBFILENAME) --env-file $(ENVFILENAME) --rm naufalfmm/cryptocurrency-price-api-migration ./migrations/cryptocurrency-price-api-migration migrate
else
	go run ./migrations/main.go migrate
endif

db_rollback:
ifeq ($(DOCKER), true)
	docker run --name cryptocurrency-price-api-migration -v $(SQLPATH):/usr/src/migrations/sql -v $(DBPATH)/$(DBFILENAME):/usr/src/$(DBFILENAME) --env-file $(ENVFILENAME) --rm naufalfmm/cryptocurrency-price-api-migration ./migrations/cryptocurrency-price-api-migration rollback $(if $(strip $(VERSION)), --version $(VERSION))
else
	go run ./migrations/main.go rollback $(if $(strip $(VERSION)), --version $(VERSION))
endif

db_create:
ifeq ($(DOCKER), true)
	docker run --name cryptocurrency-price-api-migration -v $(SQLPATH):/usr/src/migrations/sql --env-file $(ENVFILENAME) --rm naufalfmm/cryptocurrency-price-api-migration ./migrations/cryptocurrency-price-api-migration create --name $(NAME)
else
	go run ./migrations/main.go create --name $(NAME)
endif

db_init:
ifeq ($(DOCKER), true)
	docker build -t naufalfmm/cryptocurrency-price-api-migration:latest -f .\dockerfile\Dockerfile.migration .
endif

db:
	db_init && db_migrate

app_init:
ifeq ($(DOCKER), true)
	docker build -t naufalfmm/cryptocurrency-price-api:latest -f .\dockerfile\Dockerfile.app .
endif

app_run:
ifeq ($(DOCKER), true)
	docker run --name cryptocurrency-price-api -p $(PORT):$(PORT) -v $(DBPATH)/$(DBFILENAME):/usr/src/$(DBFILENAME) --env-file $(ENVFILENAME) --rm naufalfmm/cryptocurrency-price-api
else
	go run main.go
endif

app:
	app_init && app_run

run:
	db && app