### comands
go mod tidy
go get github.com/google/uuid

go build main.go
go env
GOOS=linux go build main.go
GOOS=linux GOARCH=amd64 linux go build main.go