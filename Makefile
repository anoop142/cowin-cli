BIN="bin"
REL="release"
BUILD_FLAGS="-s -w"
INSTALL_DIR="/usr/local/bin"

build:
	mkdir -p bin
	go build -o $(BIN)

compile:
	mkdir -p $(BIN)
	env GOOS=linux GOARCH=amd64 go build -o $(BIN)/cowin-cli -ldflags $(BUILD_FLAGS)
	env GOOS=windows GOARCH=amd64 go build -o $(BIN)/cowin-cli.exe -ldflags $(BUILD_FLAGS)
	env GOOS=windows GOARCH=386 go build -o $(BIN)/cowin-cli_x86.exe -ldflags $(BUILD_FLAGS)

release: compile
	mkdir -p release
	zip -9 $(REL)/cowin-cli_linux_64.zip $(BIN)/cowin-cli
	zip -9 $(REL)/cowin-cli_Windows_64.zip $(BIN)/cowin-cli.exe
	zip -9 $(REL)/cowin-cli_Windows_i386.zip $(BIN)/cowin-cli_x86.exe

clean:
	rm -rf bin release
	go clean

run: 	build
	./$(BIN)/cowin-cli

install:	build
	cp $(BIN)/cowin-cli $(INSTALL_DIR)

all:	build
