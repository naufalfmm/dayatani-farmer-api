SHELL:/bin/bash

DBFILENAME=
DBPATH=
SQLPATH=
ENVFILENAME?=.env
VERSION=
PORT=
TEST_RUNNING?=false
LINT_RUNNING?=false

db_migrate:
ifeq ($(DOCKER), true)
	docker run --name dayatani-farmer-api-migration -v $(SQLPATH):/usr/src/migrations/sql --env-file $(ENVFILENAME) --rm naufalfmm/dayatani-farmer-api-migration ./migrations/dayatani-farmer-api-migration migrate
else
	go run ./migrations/main.go migrate
endif

db_rollback:
ifeq ($(DOCKER), true)
	docker run --name dayatani-farmer-api-migration -v $(SQLPATH):/usr/src/migrations/sql --env-file $(ENVFILENAME) --rm naufalfmm/dayatani-farmer-api-migration ./migrations/dayatani-farmer-api-migration rollback $(if $(strip $(VERSION)), --version $(VERSION))
else
	go run ./migrations/main.go rollback $(if $(strip $(VERSION)), --version $(VERSION))
endif

db_create:
ifeq ($(DOCKER), true)
	docker run --name dayatani-farmer-api-migration -v $(SQLPATH):/usr/src/migrations/sql --env-file $(ENVFILENAME) --rm naufalfmm/dayatani-farmer-api-migration ./migrations/dayatani-farmer-api-migration create --name $(NAME)
else
	go run ./migrations/main.go create --name $(NAME)
endif

db_init:
ifeq ($(DOCKER), true)
	docker build -t naufalfmm/dayatani-farmer-api-migration:latest -f .\dockerfile\Dockerfile.migration .
endif

db:
	db_init && db_migrate

app_init:
ifeq ($(DOCKER), true)
	docker build -t naufalfmm/dayatani-farmer-api:latest --build-arg TEST_RUNNING=$(TEST_RUNNING) --build-arg LINT_RUNNING=$(LINT_RUNNING) -f .\dockerfile\Dockerfile.app .
endif

app_run:
ifeq ($(DOCKER), true)
	docker run --name dayatani-farmer-api -p $(PORT):$(PORT) --env-file $(ENVFILENAME) --rm naufalfmm/dayatani-farmer-api
else
	go run main.go
endif

app:
	app_init && app_run

run:
	db && app