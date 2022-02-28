build: cmd/qemantra/main.go
	go build -o qemantra cmd/qemantra/main.go

install: 
	go install ./...

.PHONY: build install
