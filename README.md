# Service application

## Run application

    go run main.go

## Compile application

    go build -o main main.go

## Start and run application locally with modules

    go mod init service
    go get .
    go run .

## Create and run docker container

    docker build -t localhost/app .
    docker run --rm -p 7080:7080 localhost/app
    docker exec -it app /bin/sh

## Create and run docker-compose environment

    docker-compose down
    docker-compose up -d db
    docker-compose up app