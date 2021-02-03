# api

[![Go Report Card](https://goreportcard.com/badge/github.com/orvosi/api)](https://goreportcard.com/report/github.com/orvosi/api)
[![Workflow](https://github.com/orvosi/api/workflows/Test/badge.svg)](https://github.com/orvosi/api/actions)
[![codecov](https://codecov.io/gh/orvosi/api/branch/master/graph/badge.svg?token=WA9A65NFR9)](https://codecov.io/gh/orvosi/api)
[![Maintainability](https://api.codeclimate.com/v1/badges/3fa0f93762298b7ae7bc/maintainability)](https://codeclimate.com/github/orvosi/api/maintainability)
[![Go Reference](https://pkg.go.dev/badge/github.com/orvosi/api.svg)](https://pkg.go.dev/github.com/orvosi/api)

## Description

API provides HTTP REST API for Orvosi application.

## SLI and SLO

- Availability: TBD
- Average response time
    - POST /sign-in: TBD
    - POST /medical-records: TBD
    - GET /medical-records: TBD
    - GET /medical-records/:id: TBD
    - PUT /medical-records: TBD

## Architecture Diagram

![orvosi-api](https://user-images.githubusercontent.com/4661221/106680454-43908300-65f1-11eb-9f60-c92e900d99f9.png)

## Owner

[Indra Saputra](https://github.com/indrasaputra)

## Onboarding and Development Guide

### How to Run

- Install Go

    We use version 1.15. Follow [Golang installation guideline](https://golang.org/doc/install).

- Install PostgreSQL

    Follow [PostgreSQL installation guideline](https://www.postgresql.org/download/).

- Go to project folder

    Usually, it would be `cd go/src/github.com/orvosi/api`.

- Run the database

    - Make sure to run PostgreSQL.

- Fill in the environment variables

    Copy the sample env file.
    ```
    cp env.sample .env
    ```
    Then, fill the values according to your setting in `.env` file.

- Download the dependencies

    ```
    make dep-download
    ```
    or run this command if you don't have `make` installed in your local.
    ```
    go mod download 
    ```

- Run the database migration

    Install `golang-migrate`. Follow [Golang Migrate CLI](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate).

    Run the migration command.
    ```
    make migrate url=<database url>
    ```

    e.g:
    ```
    make migrate url=postgres://user:password@localhost:5432/orvosi
    ```