# Code Map

Code Map explains how the codes are structured. This document will list down all folders or packages and their purpose.

## `.github`

This folder is used to define the [Github Actions](https://docs.github.com/en/actions).

## `app`

This folder contains the `main.go`. The usecase may be run and serve in multi forms, such as API, cron, or fullstack web.
To cater that case, `app` folder can contains subfolders with each folder named accordingly to the form and contain only main package.
e.g: `app/api/main.go`, `app/cron/main.go`, and `app/web/main.go`

## `db/migrations`

This folder contains all database migration files.

## `doc`

This folder contains documents for some topics.

## `entity`

This folder contains the domain of the application. Mostly, this folders contains only structs, constants, global variables, and enumerations.

## `internal`

All APIs in the internal folder (and all if its subfolders) are designed to [not be able to be imported](https://golang.org/doc/go1.4#internalpackages).
This folder contains all detail implementation specified in the `usecase` folder.

## `internal/builder`

This folder contains the [builder design pattern](https://sourcemaking.com/design_patterns/builder).
It composes all codes needed to build a full usecase.

## `internal/config`

This folder contains configuration for the application.

## `internal/http`

This folder and all of its subfolders are the place to put all codes related to REST HTTP.

## `internal/http/handler`

This folder contains the HTTP handlers.

## `internal/http/middleware`

This folder contains the HTTP middlewares.

## `internal/http/response`

This folder contains the form of HTTP response.

## `internal/http/router`

This folder contains the HTTP routes.

## `internal/http/server`

This folder contains the [Echo](https://echo.labstack.com/) HTTP server.

## `internal/repository`

This folder contains codes that connect to the database.

## `internal/tool`

This folder contains all codes that can support code the system.

## `test`

This folder contains test related stuffs.

## `test/fixture`

This folder contains a well defined support for test.

## `test/mock`

This folder contains mock for testing.

## `usecase`

This folder contains the main business logic of the application.
All interfaces and the business logic flows are defined here.
If someone wants to know the flow of the application, they better start to open this folder.