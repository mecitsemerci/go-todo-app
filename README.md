![Go](https://github.com/mecitsemerci/go-todo-app/workflows/Go/badge.svg?branch=master)

# Go + Angular Todo APP Project Template

This repository is a todo sample go and angular web project built according to Clean Architecture.  

## Technologies
* Go Web Framework ([gin-gonic](https://github.com/gin-gonic/gin))
* Containerize ([docker](https://www.docker.com/))
* Swagger ([swaggo](https://github.com/swaggo/swag))
* Database
    * [MongoDB](https://www.mongodb.com/) (default)
    * [Redis](https://redis.io/)
* Dependency Injection ([google wire](https://github.com/google/wire))
* Unit/Integration Tests ([testify](https://github.com/stretchr/testify))
* Tracing ([opentracing](https://github.com/opentracing/opentracing-go))
* Logger ([logrus](https://github.com/sirupsen/logrus))
* Error Wrapper ([pkg errors](https://github.com/pkg/errors))
* WebUI ([Angular 11](https://angular.io/))

### Web UI Preview
![GitHub Logo](https://github.com/mecitsemerci/blog/blob/master/src/images/angular_ui.gif?raw=true)

### Open API Doc Preview
![GitHub Logo](https://github.com/mecitsemerci/blog/blob/master/src/images/swagger_ui.jpg?raw=true)


## Layers and Dependencies

### `cmd` (application run)
Main application executive folder. Don't put a lot of code in the application directory.
The directory name for each application should match the name of the executable you want to have (e.g., /cmd/myapp).
It's common to have a small main function that imports and invokes the code from the /internal and /pkg directories and nothing else.

### `internal` (application codes)
Private application and library code. This is the code you don't want others importing in their applications or libraries.
* **core** includes application core files (domain objects, interfaces). It has no dependencies on another layer. 
* **pkg** includes external dependencies files and implementation of core interfaces.

### `test` (integration tests)
Application integration test folder.

### `web` (web ui)
Web application specific components: static web assets, server side templates and SPAs.

### `docs` (openapi docs)
open api (swagger) docs files. Swaggo generates automatically. 

    swag init -g ./cmd/api/main.go -o ./docs


## Usage

Open your terminal and clone the repository

    git clone https://github.com/mecitsemerci/go-todo-app.git

The application uses mongodb for default database so run makefile command

    make docker-mongo-start

This command builds all docker services so if it's ok check that application urls.  

Application | URL | Purpose
------------ | -------------| -------------
Angular UI | http://localhost:5000 | Todo APP Project
Swagger UI | http://localhost:8080/swagger/index.html | Todo API OpenAPI Docs
Jaeger UI | http://localhost:16686 | Opentracing Dashboard


By the way the application supports redis, if you use redis run that command

    make docker-redis-start

This command builds docker services so if it's ok check same application urls.

## Local Development
  ### Configuration
  The application uses environment variables. Environment variable names and values as follows by default. 
  ```
    # MONGO
    MONGO_URL=mongodb://127.0.0.1:27017
    MONGO_TODO_DB=TodoDb
    MONGO_CONNECTION_TIMEOUT=20
    MONGO_MAX_POOL_SIZE=10
    
    # REDIS
    REDIS_URL=127.0.0.1:6379
    REDIS_TODO_DB=0
    REDIS_CONNECTION_TIMEOUT=20
    REDIS_MAX_POOL_SIZE=10
    
    # JAEGER
    JAEGER_AGENT_HOST=localhost
    JAEGER_AGENT_PORT=6831
    JAEGER_SAMPLER_PARAM=1
    JAEGER_SAMPLER_TYPE=probabilistic
    JAEGER_SERVICE_NAME=go-todo-app
    JAEGER_DISABLED=false
  ```  

  ### Dependency Injection

  The project uses google wire for compile time dependency injection. The project is set for **MongoDB** by default. 
  Docker compose files generates automatically **wire_gen.go** in containers but, it must be created manually for local development. 
  
  Wire dependency file is `/internal/wired/wire_gen.go`
  
    make wire-redis
  
  This command generates **wire_gen.go** with redis provider. When `wire_gen.go` file is checked, the following change will be seen.

  ```go
  // Injectors from redis.go:
  
  func InitializeTodoHandler() (handler.TodoHandler, error) {
      client, err := redisdb.ProvideRedisClient()
      if err != nil {
          return handler.TodoHandler{}, err
      }
      todoRepository := redisdb.ProvideTodoRepository(client)
      idGenerator := redisdb.ProvideIDGenerator()
      todoService := services.ProvideTodoService(todoRepository, idGenerator)
      todoHandler := handler.ProvideTodoHandler(todoService)
      return todoHandler, nil
  }
  ```

  The following command can be run for **MongoDB** again

    make wire-mongo
  
  All changes can be observed in `/internal/wired/wire_gen.go` 
  
  ### Swagger
  
  The command that generates the open api document to `/docs` folder.

    make swag

  ### Tests
  Existing tests are for demonstration purposes only

  Unit Test run command  

    make unit-test
  
  Integration Test run command
  
    make integration-test