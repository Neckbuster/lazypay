build:
	docker build -t neckbuster/lazypay .

run:
	docker run --env-file .env.production -t neckbuster/lazypay

all: build run