tidy:
	go mod tidy
test:
	go test ./...
build: cmd/qemantra/main.go tidy
	go build -o qemantra cmd/qemantra/main.go
install: tidy build
	install -Dm 755 qemantra ~/.local/share/bin/qemantra
.PHONY: build install tidy test
