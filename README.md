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
curl -v -d '{"ingredients": ["Potato salad"]}' 'localhost:8080/search' | jq
```