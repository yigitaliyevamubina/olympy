# API Documentation

Welcome to the API documentation for the Olympic Management System! This API provides a comprehensive suite of endpoints for managing countries, events, and medals, as well as user authentication. This README will guide you through the key features and usage of the API.

## Table of Contents

- Overview
- Getting Started
- Authentication
- Events
- Countries
- Medals
- Error Handling
- API Endpoints
- Swagger Documentation

## Overview

The Olympic Management System API enables you to perform CRUD operations on countries, events, and medals, as well as manage user authentication. It is designed to be efficient, scalable, and user-friendly.

## Getting Started

To get started with the API, follow these steps:

### Clone the Repository:

```bash
git clone <https://github.com/your-repo/olympic-management-api.git>

```

### Navigate to the Project Directory:

```bash
cd olympic-management-api

```

### Install Dependencies:

Ensure you have Go installed, then run:

```bash
go mod tidy

```

### Start the Server:

```bash
go run main.go

```

### Access the API:

The API will be running at `http://localhost:9090/api/v1`.

## Authentication

The API uses token-based authentication for secure access. You need to register and log in to obtain a token.

- **Register:** `POST /api/v1/auth/register`
- **Login:** `POST /api/v1/auth/login`
- **Refresh Token:** `POST /api/v1/auth/refresh`

Include the token in the Authorization header for authenticated requests:

```
Authorization: Bearer <your-token>

```

## Events

Manage Olympic events with the following endpoints:

- **Add Event:** `POST /api/v1/events/add`
- **Edit Event:** `POST /api/v1/events/edit`
- **Delete Event:** `DELETE /api/v1/events/delete/:id`
- **Get Event:** `GET /api/v1/events/get/:id`
- **Get All Events:** `POST /api/v1/events/getall`
- **Search Events:** `POST /api/v1/events/search`

## Countries

Manage countries with the following endpoints:

- **Add Country:** `POST /api/v1/countries/add`
- **Edit Country:** `POST /api/v1/countries/edit`
- **Delete Country:** `DELETE /api/v1/countries/delete/:id`
- **Get Country:** `GET /api/v1/countries/get/:id`
- **List Countries:** `POST /api/v1/countries/getall`

## Medals

Manage medals with the following endpoints:

- **Add Medal:** `POST /api/v1/medals/add`
- **Edit Medal:** `POST /api/v1/medals/edit`
- **Delete Medal:** `DELETE /api/v1/medals/delete/:id`
- **Get Medal:** `GET /api/v1/medals/get/:id`
- **List Medals:** `POST /api/v1/medals/getall`

## Error Handling

The API returns standard HTTP status codes for errors:

- **400 Bad Request:** Invalid request data.
- **404 Not Found:** Resource not found.
- **500 Internal Server Error:** Server error.

Error responses include a message with details about the error.

## API Endpoints

For detailed information about each endpoint, including request parameters and response formats, refer to the Swagger documentation.

## Swagger Documentation

Explore and interact with the API using Swagger UI:

- **Swagger Documentation:** `http://localhost:9090/swagger/index.html`
