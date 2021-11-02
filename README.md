# Deporvillage - Go TCP API - **Feeder Service (product SKU)**

This repository contains the code examples used on the technical test for Deporvillage.

This documentation could be accessed in a [Notion page](https://manurua.notion.site/Technical-Test-Deporvillage-885b3520a43b488992dc7ac0bd061e3f).

## Requirements

- Go v1.17+

## Contents

- This project has been designed as a single Go module.
- Go Modules is provided to manage dependencies.
- Docker folder to deploy service in containers.
- Top level folders are created following standard project layout for Golang:
    - [https://github.com/golang-standards/project-layout](https://github.com/golang-standards/project-layout)
    - [https://blog.friendsofgo.tech/posts/como_estructurar_tus_aplicaciones_go/](https://blog.friendsofgo.tech/posts/como_estructurar_tus_aplicaciones_go/)
- Internal folder structure follows DDD and Hexagonal Architecture principles:
    - Shared and Product contexts.
    - Domain, application and infrastructure layers.
    - Product aggregate as main data model.
    - Repository pattern for data management.
    - Value objects to validate SKU format.
- Tests are close to executable code.

## Assumptions

- Due to these requirements:
    - *The Application must write a de-duplicated list of these numbers to a log file in no particular order.*
    - *Your Application may not for any part of its operation use or require the use of external systems, for example Apache Kafka or Redis.*

    I have not implemented a solution with an external database like MySQL or PostgreSQL and persist info with a log file.

    The sentence *"Use only the standard lib, except for the database drivers."* was a little confusing for me with that.

- The service intends to continue growing and expanding, so a clean architecture is important to allow this in the future.
- The service will be exec in a Linux environment with Docker containers. This is important for above line of code, because Windows OS not support Kill function.

    ```go
    syscall.Kill(syscall.Getpid(), syscall.SIGINT)
    ```

- Input with zeroes that exceed 9 characters are consider invalids.

## Dependencies - Go Modules

Go Modules is provided to manage dependencies, but the project in the current state does not require any external dependency.

Service use standard library only.

For installing a go module, just exec:

```bash
go get package_url
```

For clean and install dependencies in the code, exec:

```bash
go mod tidy
```

## Simple usage (Linux OS)

To execute the application, just run:

```bash
go run cmd/api/main.go
```

To build it into a executable:

```bash
go build cmd/api/main.go
./main
```

## Recommend usage: Docker

The project provides a **Makefile** with some commands to easy-way management of Docker files.

It calculates which Docker-Compose files are executed in base to ENV variable.

If not provided, ENV variable is *development* by default. Possible values are *test, debug, development, staging* and *production.*

To run it, just execute:

```bash
make start ENV=env
```

To run it, building images, execute:

```bash
make start-build ENV=env
```

To stop it, execute:

```bash
make stop ENV=env
```

To build containers without cache, execute:

```bash
make build ENV=env
```

### Development

For developing more code into project, it is recommended to launch Makefile commands with ENV=development.

It gets *dev-api.Dockerfile* to create **api** service, which provides a Golang image where **air** package is installed to exec service with live reload.

To see more: [https://github.com/cosmtrek/air](https://github.com/cosmtrek/air)

```bash
make start ENV=development
```

Makefile commands set *development* ENV by default if not provided.

### Debug

For debugging existing code from project, it is recommended to launch Makefile commands with ENV=debug.

It gets *debug-api.Dockerfile* to create **api** service, which provides a Golang image where **delve** package is installed to exec service with a debugger on port 40000.

To see more: [https://github.com/go-delve/delve](https://github.com/go-delve/delve)

```bash
make start ENV=debug
```

For VS Code users, a configuration is provided to debug with **Delve** into *launch.json* of *.vscode* folder.

### Production

For launching project in a production environment, it is recommended to launch Makefile commands with ENV=production.

It gets *prod-api.Dockerfile* to create **api** service, which provides a Golang **alpine** image where a **single executable** file is generated to be launched.

```bash
make start ENV=production
```

## Environment variables

- Important service configuration parameters is taken from environment variables.
- If some param is not defined in environment variables, Docker Compose files set a default value.
- In last instance, if the service can not take the value from env variable (launched without Docker for example), a default value is defined in code.

To set a env variable, just run:

```bash
export API_PORT=value
```

## Tests

To execute all tests with coverage, just run:

```bash
go test -cover ./...
```

### Extra: GoConvey

To watch live results of tests in the browser with a very cool UI.

[GoConvey - Go testing in the browser](http://goconvey.co/)

## Client Provider

### Manual

To experience the client side, any TCP client is valid to connect to service, send data and watch complete workflow.

For example, **netcat** utility is used in tests:

```bash
nc 127.0.0.1 4000
```

Then, test to write some product SKU and see how the client is disconnect.

- If the input has a valid format and it does not exist before, it will be persisted in the log file into *tmp/products.log*, and will count as unique product.
- If the input has a valid format but is exists before, it will count as duplicated product.
- If the input has an invalid format, it will count as invalid product.
- If the input is exactly *terminate*, all clients are disconnected and server is shutting down immediately.

### Auto

In `cmd/client`, I provide a short go program with a TCP client that generates random strings and send it to server, reconnecting all time and repeating the process every 2 seconds.

There is a flag for generating valid product SKUs. Otherwise, inputs will be invalid by default.

```bash
go run cmd/client/main.go
```
