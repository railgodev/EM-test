# Project Overview

This project is a backend service for managing subscriptions, written in Go. It provides a RESTful API for creating, reading, updating, and deleting subscriptions.

**Key Technologies:**

*   **Backend:** Go
*   **API Framework:** Gin
*   **Database:** PostgreSQL
*   **Database Driver:** pgx
*   **Migrations:** golang-migrate
*   **Containerization:** Docker

**Architecture:**

The project follows a clean architecture pattern, with a clear separation of concerns between the API handlers, use cases, and repository layers.

# Building and Running

## Prerequisites

*   Go
*   Docker
*   Docker Compose
*   make

## Configuration

The application is configured using environment variables. An example `.env` file is provided as `.example.env`. Rename this file to `.env` and modify the values as needed.

## Running the Application

There are two ways to run the application:

**1. Running the Go compiler on the host and PostgreSQL in a container:**

```bash
# Start the PostgreSQL container
make up-postgres

# Run the Go application
make run
```

**2. Running both the Go application and PostgreSQL in containers:**

```bash
# Start both containers
make up
```

To stop and remove the containers, run:

```bash
make down
```

# Development Conventions

## API Endpoints

The following API endpoints are available:

*   `GET /api/v1/subscriptions?page=1&page_size=10`: List all subscriptions.
*   `POST /api/v1/subscriptions`: Create a new subscription.
*   `GET /api/v1/subscriptions/{id}`: Get a subscription by its ID.
*   `PUT /api/v1/subscriptions/{id}`: Update a subscription.
*   `DELETE /api/v1/subscriptions/{id}`: Delete a subscription.
*   `GET /api/v1/subscriptions/total?start=07-2025&end=09-2025&service_name=string&user_id=UUID`: Get the total number of subscriptions.

## Database Migrations

Database migrations are located in the `backend/migrate/migrations` directory. To run the migrations, the application will automatically run them on startup.


