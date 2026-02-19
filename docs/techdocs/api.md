# API Reference

The full interactive API reference is available via the Swagger UI when the service is running locally.

## Swagger UI

```
http://localhost:8080/swagger
```

## Raw OpenAPI Spec (YAML)

```
http://localhost:8080/swagger/doc.yaml
```

Or view the hosted spec directly:

```
https://github.com/shyam-hande/golang-rest/blob/main/docs/swagger.yaml
```

## Authentication

Protected endpoints require a **Bearer JWT token** in the `Authorization` header:

```
Authorization: Bearer <token>
```

Obtain the token by calling `POST /login` with valid credentials.

## Models

### Event

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| ID | integer | — | Auto-assigned unique ID |
| Name | string | ✓ | Event name |
| Description | string | ✓ | Event description |
| Location | string | ✓ | Event location |
| DateTime | datetime | ✓ | ISO 8601 date/time |
| UserID | integer | — | Owner's user ID |

### User

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| ID | integer | — | Auto-assigned unique ID |
| Email | string | ✓ | Unique email address |
| Password | string | ✓ | Plain text (hashed on save) |
