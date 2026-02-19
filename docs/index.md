# Go Events REST API

A RESTful API for managing events, built with **Go**, **Gin**, **JWT authentication**, and **SQLite**.

## Overview

This service provides a simple but complete event management backend:

- **Users** can sign up and log in to receive a JWT token
- **Events** can be browsed by anyone (read-only, no auth needed)
- **Authenticated users** can create, update, and delete their own events

## Architecture

| Layer | Technology |
|-------|-----------|
| Framework | [Gin](https://github.com/gin-gonic/gin) |
| Auth | JWT ([golang-jwt/jwt](https://github.com/golang-jwt/jwt)) |
| Database | SQLite ([go-sqlite3](https://github.com/mattn/go-sqlite3)) |
| Logging | `log/slog` (structured JSON) |

## Quick Start

```bash
# Clone the repository
git clone https://github.com/shyam-hande/golang-rest.git
cd golang-rest

# Run the server (starts on :8080)
go run app.go
```

## Available Endpoints

| Method | Path | Auth | Description |
|--------|------|------|-------------|
| POST | `/signup` | No | Register a new user |
| POST | `/login` | No | Login, receive JWT token |
| GET | `/events` | No | List all events |
| GET | `/events/:id` | No | Get event by ID |
| POST | `/events` | Yes | Create a new event |
| PUT | `/events/:id` | Yes | Update your event |
| DELETE | `/events/:id` | Yes | Delete your event |
| GET | `/swagger` | No | Swagger UI |
| GET | `/swagger/doc.yaml` | No | Raw OpenAPI spec |

## Project Structure

```
.
├── app.go              # Entry point
├── db/                 # Database initialisation
├── middleware/         # Auth & CORS & Logger middleware
├── models/             # Event and User models (DB logic)
├── routes/             # HTTP handlers + docs route
├── utils/              # JWT, hashing, logger utilities
└── docs/
    └── swagger.yaml    # OpenAPI 3.0 specification
```
