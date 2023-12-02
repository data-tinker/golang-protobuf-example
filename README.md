# Go Protocol Buffers Example
Protocol Buffers usage example in Go

## Running Locally
* Build the protocol buffer definition: `protoc --go_out=. *.proto`
* Run the program: `go run main.go`
* Run tests: `go test ./...`
* Alternatively, run the app using Docker: `make build && make run`

## Technologies Used
* [Go](https://go.dev/)
* [Protocol Buffer](https://protobuf.dev/getting-started/gotutorial/)
* [Docker](https://www.docker.com/)
