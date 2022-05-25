tidy:
	go mod tidy
test:
	go test -v ./...
	go test -cover ./...

build: cmd/qemantra/main.go tidy
	go build -o qemantra cmd/qemantra/main.go
install: tidy build
	install -Dm 755 qemantra ~/.local/share/bin/qemantra
format:
	go fmt ./...
.PHONY: build install tidy test
