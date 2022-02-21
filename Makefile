build: cmd/qemantra/main.go
	go build -o qemantra cmd/qemantra/main.go

install: build
	go install ./...

.PHONY: build install
