all: server client

server: cmd/server/main.go
	go build go-tcp-server/cmd/server
client: cmd/client/main.go
	go build go-tcp-server/cmd/client

clean:
	rm -fr ./server
	rm -fr ./client