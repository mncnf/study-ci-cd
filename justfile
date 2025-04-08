build:
    docker-compose up -d --build

up:
    docker-compose up -d

down:
    docker-compose down --remove-orphans

restart:
    @just down
    @just up

exec:
    docker-compose exec app sh

run:
    docker compose run --rm app sh
