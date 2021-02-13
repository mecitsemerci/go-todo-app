![Go](https://github.com/mecitsemerci/go-todo-app/workflows/Go/badge.svg?branch=master)

# Clean Architecture Go Project Template

This repository is a sample go lang web project built according to Clean Architecture.  

## Technologies
* Web Framework (gin-gonic)
* Docker
* Swagger (swaggo)
* NoSQL Database
    * Mongodb (default)
    * Redis (soon)
* Dependency Injection (wire)
* Unit/Integration Tests (testify)

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