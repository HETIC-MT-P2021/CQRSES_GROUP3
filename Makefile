test:
	go test -v ./tests

run:
	docker-compose up --build -d

init:
	openssl genrsa -des3 -out private.pem 2048
	openssl rsa -in private.pem -outform PEM -pubout -out public.pem
	make run
stop:
	docker-compose stop