.SILENT:

build: 
	go build -o bin/main

run: build
	./bin/main

test:
	go test

clean:
	rm -rf bin/
	go clean -cache -testcache -modcache