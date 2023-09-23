all:
	DOCKER_BUILDKIT=1 docker-compose up --build --force-recreate --remove-orphans -d

test: all
	docker-compose run test-client

clean:
	docker-compose down --remove-orphans -v

-include include.mk
