ROOT=github.com/partkyle/go-udp

all: target/server-linux-amd64 target/server-darwin-amd64 target/client-linux-amd64 target/client-darwin-amd64 

target/server-linux-amd64:
	GOOS=linux GOARCH=amd64 go build -o target/server-linux-amd64 $(ROOT)/server

target/server-darwin-amd64:
	GOOS=darwin GOARCH=amd64 go build -o target/server-darwin-amd64 $(ROOT)/server

target/client-linux-amd64:
	GOOS=linux GOARCH=amd64 go build -o target/client-linux-amd64 $(ROOT)/client

target/client-darwin-amd64:
	GOOS=darwin GOARCH=amd64 go build -o target/client-darwin-amd64 $(ROOT)/client

clean:
	rm -rf target/

