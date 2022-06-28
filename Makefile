MAINFILE=cmd/qemantra/main.go
BINARY=qemantra
VERSION=0.0.1
LDFLAGS="-X main.VERSION='$(VERSION)'"

BUILD_DIR=build

tidy:
	go mod tidy
test:
	go test -v ./...
	go test -cover ./...

build: $(MAINFILE) tidy
	go build -o $(BINARY) -ldflags $(LDFLAGS) $(MAINFILE)
install: tidy build
	install -Dm 755 $(BINARY) ~/.local/share/bin/$(BINARY)
format:
	go fmt ./...
build-release: build-clean
	make compile-release
	make compress-release

compress-release:
	tar -C $(BUILD_DIR) -cvzf  build/$(BINARY)-$(VERSION)-linux-amd64.tar.gz $(BINARY)-$(VERSION)-linux-amd64
	tar -C $(BUILD_DIR) -cvzf  build/$(BINARY)-$(VERSION)-linux-arm.tar.gz $(BINARY)-$(VERSION)-linux-arm
	tar -C $(BUILD_DIR) -cvzf  build/$(BINARY)-$(VERSION)-linux-i386.tar.gz $(BINARY)-$(VERSION)-linux-i386
compile-release:
	GOOS=linux GOARCH=amd64 go build  -ldflags $(LDFLAGS) -o build/$(BINARY)-$(VERSION)-linux-amd64 $(MAINFILE)
	GOOS=linux GOARCH=arm go build -ldflags $(LDFLAGS) -o build/$(BINARY)-$(VERSION)-linux-arm $(MAINFILE)
	GOOS=linux GOARCH=386 go build -ldflags $(LDFLAGS) -o build/$(BINARY)-$(VERSION)-linux-i386 $(MAINFILE)
build-clean:
	rm -rf $(BUILD_DIR)
	mkdir $(BUILD_DIR)
.PHONY: build install tidy test build-clean compile-release compress-release
