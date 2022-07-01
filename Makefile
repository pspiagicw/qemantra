MAINFILE=cmd/qemantra/main.go
BINARY=qemantra
VERSION=0.0.1
LDFLAGS="-X main.VERSION='$(VERSION)'"
BUILD_DIR=build
.DEFAULT_GOAL := build

.PHONY: build install tidy test build-clean compile-release compress-release

tidy:
	go mod tidy
test:
	go test  ./...
	go test -cover ./...

build: $(MAINFILE) tidy
	go build -o $(BINARY) -ldflags $(LDFLAGS) $(MAINFILE)

install: tidy build
	install -Dm 755 $(BINARY) ~/.local/share/bin/$(BINARY)
format:
	go fmt ./...
release: build-clean
	cp README.md build/
	cp LICENSE build/
	make compile-release
	make compress-release
compress-release:
	tar -C $(BUILD_DIR) -cvzf  build/$(BINARY)-$(VERSION)-linux-amd64.tar.gz $(BINARY)-$(VERSION)-linux-amd64 LICENSE README.md
	tar -C $(BUILD_DIR) -cvzf  build/$(BINARY)-$(VERSION)-linux-arm.tar.gz $(BINARY)-$(VERSION)-linux-arm LICENSE README.md
	tar -C $(BUILD_DIR) -cvzf  build/$(BINARY)-$(VERSION)-linux-i386.tar.gz $(BINARY)-$(VERSION)-linux-i386 LICENSE README.md
compile-release:
	GOOS=linux GOARCH=amd64 go build  -ldflags $(LDFLAGS) -o build/$(BINARY)-$(VERSION)-linux-amd64 $(MAINFILE)
	GOOS=linux GOARCH=arm go build -ldflags $(LDFLAGS) -o build/$(BINARY)-$(VERSION)-linux-arm $(MAINFILE)
	GOOS=linux GOARCH=386 go build -ldflags $(LDFLAGS) -o build/$(BINARY)-$(VERSION)-linux-i386 $(MAINFILE)
clean:
	rm -rf $(BUILD_DIR)
	mkdir $(BUILD_DIR)
