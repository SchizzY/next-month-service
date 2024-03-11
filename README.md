## next-month-service

**next-month-service** is the `golang` backend for the next-month budgeting app.

## Pre-requisites

- [`golang`](https://go.dev/) 1.16 or higher
- `docker` 20.10.5 or higher
- [`justfile`](https://github.com/casey/just)

## Running the app

1. Make a copy of the `.env.example` file and rename it to `.env`.
2. Set the `CON` variable to the connection string of your database.
3. Install all required dependencies
4. Run the app using `just run` (as a local server) or using `just watch` as a local docker docker container.
