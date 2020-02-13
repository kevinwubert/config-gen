all: build

build:
	go build -o bin/config-gen cmd/config-gen/*.go

clean:
	rm -rf bin