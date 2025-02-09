## Base project hexagonal GoFr

### Framework GoFr

go get gofr.dev

### Hexagonal

go mod init hexagonal-gofr
go mod tidy
go run cmd/api/main.go

### Test

Create user
curl -X POST -H "Content-Type: application/json" -d '{"id": 1, "name": "John Doe", "email": "john@example.com"}' http://localhost:8000/users

Get user
curl http://localhost:8000/users/1