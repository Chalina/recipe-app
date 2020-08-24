# Recipe App

## Prerequisites
- Go 1.14
## Setup
```
go mod download
go run main.go
```
## Running tests
```
go test ./...
```
## Example request
```
curl -X POST -d '{"ingredients": ["chocolate"]}' 'localhost:8080/search'
```