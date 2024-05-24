# Product App Go

This project is a simple CRUD application developed in Go. It allows you to perform basic management operations for products.

## Features

- Create, Read, Update, and Delete (CRUD) products
- Docker support for containerization
- Environment configuration with `app.env`
- Middleware for authentication
- Organized routing for products, users, orders, authentication, and logs

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/mustafamertkisa/product-app-go.git
   cd product-app-go
   ```

2. Set up the environment variables:

   ```bash
   cp app.env.example app.env
   ```

3. Navigate to the `cmd` directory:

   ```bash
   cd cmd
   ```

4. Build and run the application using Docker:
   ```bash
   docker-compose up --build
   ```

## Usage

After starting the application, it will be accessible at `http://localhost:8080`.

## Router Setup

The router is configured to handle different groups of routes:

- **Health Check**: `GET /healthchecker`
- **Products**:
  - `POST /products`
  - `GET /products`
  - `DELETE /products/:productId`
  - `GET /products/:productId`
  - `PUT /products/:productId`
- **Users**:
  - `POST /users`
  - `GET /users`
  - `DELETE /users/:userId` (requires auth)
  - `GET /users/:userId`
  - `PUT /users/:userId` (requires auth)
- **Orders**:
  - `POST /orders`
  - `GET /orders`
  - `DELETE /orders/:orderId`
  - `GET /orders/:orderId`
  - `PUT /orders/:orderId`
- **Authentication**:
  - `POST /auth/register`
  - `POST /auth/login`
  - `POST /auth/user`
  - `POST /auth/logout`
- **Logs**:
  - `DELETE /logs`
  - `GET /logs`
  - `GET /logs/:id`
  - `DELETE /logs/:id`
  - `GET /logs/user/:userId`

## Project Structure

- `cmd`: Contains the main application code.
- `internal`: Contains internal packages and business logic.
- `test`: Contains tests for the application.

## Dependencies

- Go
- Docker
- Fiber
