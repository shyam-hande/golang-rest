# Testing

This project uses Go's standard `testing` package along with [testify](https://github.com/stretchr/testify) for assertions.

## Running Tests

```bash
# Run all tests
go test ./...

# Run with verbose output
go test ./... -v

# Run with coverage report
go test ./... -cover
```

## Test Structure

| Package | Test File | What It Tests |
|---------|-----------|---------------|
| `middleware` | `auth_test.go` | JWT auth middleware |
| `middleware` | `logger_test.go` | Logger middleware |
| `utils` | `jwt_test.go` | Token generation & validation |
| `utils` | `hash_test.go` | Password hashing |
| `root` | `add_test.go` | Example unit test |

## More Details

See [TESTING.md](https://github.com/shyam-hande/golang-rest/blob/main/TESTING.md) in the repository for detailed testing setup and examples.
