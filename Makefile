tidy:
	go mod tidy
test:
	go test ./...
build: cmd/qemantra/main.go tidy
	go build -o qemantra cmd/qemantra/main.go
install: tidy
	go install ./...
.PHONY: build install tidy test
