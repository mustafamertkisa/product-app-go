# Product App Go

This is a simple CRUD (Create, Read, Update, Delete) API built with Go using Fiber framework and PostgreSQL for database management.

## Setup Instructions

To run this project, you need to have Docker installed on your machine.

1.  Clone this repository to your local machine.

    `git clone https://github.com/your-username/product-app-go.git`

2.  Navigate to the project directory.

    `cd product-app-go`

3.  Build and start the Docker containers.

    `docker-compose up --build`

4.  Once the containers are up and running, you can access the API at `http://localhost:8000/api`.

## Project Structure

- `Dockerfile`: Contains instructions to build the Docker image for the Go application.
- `docker-compose.yml`: Defines the services needed for the project, including PostgreSQL and the Go API.
- `internal/application`: Contains controllers and router setup.
- `internal/domain`: Contains domain logic including models, repositories, services.
- `internal/infrastructure/config`: Contains configurations for database connection and environment variables.
- `main.go`: Entry point of the application where server setup and initialization occur.

## Dependencies

- **Fiber**: Web framework for Go.
- **Validator**: Struct validation library.
- **PostgreSQL**: Database management system.
- **GORM**: ORM library for database interaction.

## Endpoints

- **GET /api/products**: Get all products.
- **GET /api/products/:id**: Get a single product by ID.
- **POST /api/products**: Create a new product.
- **PUT /api/products/:id**: Update an existing product.
- **DELETE /api/products/:id**: Delete a product by ID.
- **GET /api/users**: Get all users.
- **GET /api/users/:id**: Get a single user by ID.
- **POST /api/users**: Create a new user.
- **PUT /api/users/:id**: Update an existing user.
- **DELETE /api/users/:id**: Delete a user by ID.
- **GET /api/orders**: Get all orders.
- **GET /api/orders/:id**: Get a single order by ID.
- **POST /api/orders**: Create a new order.
- **PUT /api/orders/:id**: Update an existing order.
- **DELETE /api/orders/:id**: Delete an order by ID.
