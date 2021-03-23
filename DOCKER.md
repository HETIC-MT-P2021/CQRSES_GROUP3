#  Docker Usage

## Services 

The stack uses Docker and Docker-compose to run 5 services:

API - Back-end container tp run the application endpoints and produce events and retrieve the read model.

Consumer app - Back-end container to run the consumer that will retrieve the produced event from the message broker (RabbitMQ) 

Postgres - Relational Database system

ElasticSearch - Document oriented database system to store the events.

RabbitMQ - Message broker used to dispatch the different events for the dedicated consumers.

Adminer - Database management system with gui interface

Kibana - Visualization dashboard for Elasticsearch 

## Setup

1. Move your terminal in the root directory of the project.
2. Run `docker-compose up --build -d` The flag `-d ` stands for detached if you don't want to have all the docker logs. 
3. Access to the api app through `http://localhost:8000`
> Note: If you already build the project once you don't need to add the `--build` flag. You will only need it some build changes are made


> Disclaimer: sometime the ElasticSearch and RabbitMQ services starts after the api or the consumer app and causes crash
### In case of crash

- For the api:

Change a file in the api app under `/app` directory and the hotreload will do the rest

- For the consumer:

1. log into your terminal at the root directory of the project
2. Run `docker-compose exec app-consumer /bin/sh`, you should be in the container in the directory `consumer` then run `air -c .consumer.toml` 

## Ports

|Container|External Port|Internal Port|
|---------|-------------|-------------|
|API|8000|8000|
|Consumer|8082|8082|
|Postgres|5432|5432|
|ElasticSearch|9200|9200|
|RabbitMQ|5672|5672|
|RabbitMQ GUI|15672|15672|
|Kibana|5601|5601|
|Adminer|8080|8080|

> For rabbitMQ login, ID: user, Password: bitnami