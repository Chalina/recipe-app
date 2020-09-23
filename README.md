# Recipe App

## Prerequisites
- Go 1.14
- Docker 

## Setup
```
docker-compose up
```

## Running tests
```
go test ./...
```
## Example request
```
curl -X POST -d '{"ingredients": ["chocolate"]}' 'localhost:8080/search'
```