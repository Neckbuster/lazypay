build:
	docker build -t neckbuster/lazypay .

run:
	docker run -p 8080:8080 --env-file .env.production -t neckbuster/lazypay

all: build run