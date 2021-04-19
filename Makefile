.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(firstword $(MAKEFILE_LIST)) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

test:	## Runs all test contained in the tests folder
	docker-compose exec api sh ./tests.sh

run: ## Start the app containers
	docker-compose up --build -d

init: ## refresh the env file, generate public and private RSA keys and start the app
	cp .env.dist .env
	openssl genrsa -des3 -out private.pem 2048
	openssl rsa -in private.pem -outform PEM -pubout -out public.pem
	make run
stop: ## stop all running containers of the application
	docker-compose stop