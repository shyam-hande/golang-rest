# Testing Guide

This document explains the testing infrastructure for the Go REST API.

## 1. Running Tests
You can run tests using the standard `go test` command.

| Command | Description |
| :--- | :--- |
| `go test ./...` | Run all tests in the project (recursive). |
| `go test -v ./...` | Run all tests with verbose output (shows logs and test names). |
| `go test -v ./utils` | Run tests for a specific package (e.g., `utils`). |
| `go test -cover ./...` | Run tests and show code coverage percentage. |

## 2. Tools & Libraries
We use the following libraries to make testing easier and more robust:

- **`testing`**: The standard Go library for writing tests.
- **`github.com/stretchr/testify`**: A toolkit that provides:
    - **`assert`**: Readable assertions (e.g., `assert.Equal`, `assert.NoError`).
    - **`require`**: Similar to `assert` but stops the test immediately on failure.
- **`net/http/httptest`**: A standard library for mocking HTTP requests/responses, essential for testing handlers and middleware.

## 3. Testing Strategies

### A. Unit Tests (Pure Functions)
For logic that doesn't depend on external systems (like databases or HTTP), we use simple unit tests.

**Example: Password Hashing**
*File: `utils/hash.go`*
```go
func TestHashPassword(t *testing.T) {
    password := "secret"
    hash, err := HashPassword(password)
    
    // Assertions check values directly
    assert.NoError(t, err)
    assert.NotEqual(t, password, hash)
}
```

### B. Middleware & Handler Tests (HTTP)
Testing HTTP middleware requires simulating a web server environment without actually spinning one up. We do this using `httptest`.

**How it works:**
1.  **Response Recorder**: `httptest.NewRecorder()` acts as a "fake" `ResponseWriter` that captures the response code, body, and headers.
2.  **Request Construction**: `http.NewRequest(...)` creates a request object manually.
3.  **Gin Context**: We use `gin.CreateTestContext` or simply `r.ServeHTTP(w, req)` to pass the request through the Gin engine.

**Example: Auth Middleware**
*File: `middleware/auth_test.go`*
```go
func TestAuthenticate(t *testing.T) {
    // 1. Setup Router with Middleware
    r := gin.New()
    r.Use(Authenticate)
    r.GET("/protected", func(c *gin.Context) {
        c.Status(200) // If middleware passes, this handler runs
    })

    // 2. Create the Recorder (catches the response)
    w := httptest.NewRecorder()
    
    // 3. Create the Request (simulating a client)
    req, _ := http.NewRequest("GET", "/protected", nil)
    req.Header.Set("Authorization", "valid-token")

    // 4. Serve the Request
    r.ServeHTTP(w, req)

    // 5. Assert valid code
    assert.Equal(t, 200, w.Code)
}
```

## 4. Best Practices
- **Test File Location**: Keep `_test.go` files in the same package as the code they test.
- **Table-Driven Tests**: Use table-driven tests for validating multiple input/output scenarios (seen in `add_test.go`).
- **Mocking**: For external dependencies (like DBs), interfaces and mocks should be used (though we currently use an in-memory or file-based DB setup for simplicity).
