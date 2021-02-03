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

### Development Guide

- Fork the project

- Create a meaningful branch

    ```
    git checkout -b <your-goal>
    ```
    e.g:
    ```
    git checkout -b strenghten-security-on-sign-in-process
    ```

- Create some changes and their tests (unit test and any test if any).

- Make sure to have unit test coverage at least 90%. There will be times when the code is quite hard to test. Please, explain it in your Pull Request.

- Push the changes to repository.

- Create Pull Request (PR) for your branch. In your PR's description, please explain the goal of the PR and its changes.

- Ask the other contributors to review.

- Once your PR is approved and its pipeline status is green, ask the owner to merge your PR.

## Request Flows, Endpoints, and Dependencies

## Request Flows

### Sign-in with Google

![orvosi-sign_in](https://user-images.githubusercontent.com/4661221/106688293-65ddcd00-6600-11eb-8a57-63b5feca37df.png)

### Create a medical record

![orvosi-create_medical_record](https://user-images.githubusercontent.com/4661221/106689102-008adb80-6602-11eb-8965-3f3546c41c0f.png)

### Get user's medical records

![orvosi-get_medical_records](https://user-images.githubusercontent.com/4661221/106690208-fe298100-6603-11eb-8cc3-36c92458a9ca.png)

### Get medical record's detail

![orvosi-get_medical_record](https://user-images.githubusercontent.com/4661221/106690781-fb7b5b80-6604-11eb-894b-41e6a65d3409.png)

### Update medical record's detail

![orvosi-update_medical_record](https://user-images.githubusercontent.com/4661221/106691051-6cbb0e80-6605-11eb-9dfa-a99805df52d9.png)
