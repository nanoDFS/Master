BINARY_NAME = master
SRC = ./main.go

build:
	go build -o $(BINARY_NAME) $(SRC)

run:
	go run $(SRC)

test:
	go test ./...

clean:
	rm -f $(BINARY_NAME)

fmt:
	go fmt ./...

deps:
	go mod tidy

proto:
	protoc --go_out=. --go-grpc_out=. proto/*.proto


req: 
	echo "$(msg)" | nc 127.0.0.1 9000

default: build


