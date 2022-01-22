build:
	docker build -t neckbuster/lazypay docker/

run:
	docker run --env-file .env.production -t neckbuster/lazypay

all: build run