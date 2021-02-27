# DDD-API-Go

This repository content an API based in Hexagonal Architecture.

We use the framework Gin in order to move it forward faster.

## Set new entrypoint/controller

- Create a new handler and defined in the handler folder


## Project structure

- cmd/api: Entrypoints for our service
- internal/platform: Infrastructure references 
- internal/platform/server: References to our server (gin in this case)
- internal/platform/storage: 
- internal/platform/storage/mysql/: Mysql implementation
- internal/platform/storage/storagemocks: 


## Mocks

We use the libray: [mockery](http://github.com/vektra/mockery/v2/)

You can import it using:
```go
go get github.com/vektra/mockery/v2/.../
```
