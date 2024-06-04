APP ?= app
all:
	docker-compose up --build --force-recreate --remove-orphans -d

healthcheck:
	docker inspect $(APP) --format "{{ (index (.State.Health.Log) 0).Output }}"

test: all
	docker-compose run test-client

clean:
	docker-compose down --remove-orphans -v

-include include.mk
