# Clean Todo Rest Api with Gin Web Framework

This repository is a sample go lang web project built according to Clean Architecture.  

## Build with
* Gin Web Framework
* Docker support
* :soon: MongoDb support
* :soon: Firebase support
* :soon: Swagger UI support

### Layers and Dependencies

![image](./docs/img/layers.png)

## API (HTTP Web Api Layer)
This layer is handling all HTTP requests messages on controllers. 
Routes and DTOs (Data Transfer Objects) are defined in this layer.

## CORE (Business Layer)
This layer consists of business rules which has domain models and use cases. 
There are mostly no external dependencies in this layer, no network connections, databases, etc. allowed.
Data is transmitted to this layer via repositories and clients.

## INFRA (Infrastructure Layer)
All external dependencies are defined in this layer. 
Connections with external data resources (api, database etc.) are made through this layer via clients and database providers.
We can add configuration processes and utilities here.

## Installation
 Open your terminal and clone this repository.
 
    git clone https://github.com/mecitsemerci/clean-go-todo-api.git

If docker is running, run docker compose up command in the folder.

    docker-compose up

Check the app is running on http://localhost:8080/swagger/index.html

 ## Swagger UI Preview
 
 ![image](./docs/img/swagger_ui.png)
 
