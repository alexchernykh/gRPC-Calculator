# gRPC calculator

This is a gRPC implementation of simple math tasks.
![gRPC Calculator](img/gRpc_2.png)

## Description

The aim is to show the calculation of simple math tasks like `addition`, `subtraction`, `multiplication` and `division` based on a client server architecture with the use of the gRPC.
This Task requires several major steps.

1. Initialise go.mod

    ```shell script
    go mod init gRPC-Calculator
    ```

1. create a description `.proto` file, which will describe the communication between client and server and the methods, that will be called.
    - Here need to be defined the `message` structure to be send(`Request`) and received(`Response`) by the **client** and the rpc function(`CalcResult`) in the service(`MessageService`).
    - For detailed information see [here](https://grpc.io/docs/what-is-grpc/introduction/)
1. Use `protoc` to generate interface code. Follow the [installation and generation process](#protogenerration)
 The generated interface `CalcResult` need to be overwritten in the **server** part.

1. Create a server and client logic

## Generating the gRPC client and server interface from `.proto` file <a id="protogenerration"></a>
1. Install instrumments for OS X with Homebrew
    ```shell script
    brew install protobuf
    ```

1. Installing `protoc-gen-go` according to the description on [StackOverflow](https://stackoverflow.com/questions/57700860/protoc-gen-go-program-not-found-or-is-not-executable)
    ```shell script
    go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
    ```

1. Export GOPATH
    ```shell script
    echo 'export GOPATH=$HOME/Go' >> $HOME/.bashrc
    source $HOME/.bashrc
    ```
    Either if you use `zsh` instead, write and source the `.zshrc` files

1. Generating the `calc.pb.go` file with for gRPC
    ```shell script
    protoc --go_out=plugins=grpc:./gen --go_opt=paths=source_relative calc.proto
    ```
    [//]: <Like> (<--protoc -I . --go_out=/Users/chernykh_alexander/go/src/tutorial/ calc.proto-->)

1. Generated `calc.pb.go` will contain protocols and data structures, that will be used during the gRPC calls.


### Get Started 

An example gRPC calculator usage could look like the following. By running the server in a first shell user will be notified, that it is serving and waiting for calls from client. By starting the client in the second shell window, user will be prompted to enter a simple equation, f.e `1 + 2 `. All allowed operations are `+`, `-`,`*` and `/`.


1. In the first shell, start the server `gServer.go`
    ```shell script
    go run gServer.go
    ```

2. In the second shell, start the client `gClient.go`

    ```shell script
    go run gClient.go
    ```
3. Input some math operation
    ```
    1 + 1
    ```

    ```
    1000 / 13
    ```

    ```
    1000 * 13
    ```

### Tests

For tests execution:
```shell script
go test
```

To atumate the tests on every commit was added a Github Actions workflow.


### Docker

#### Server

to run the server as a Dockerfile:

1. Build the Dockerfile:
    
    ```shell script
    docker build -t server -f server/Dockerfile .
    ```

1. Run the container

    ```shell script
    docker run -d -p 8080:8080 server
    ```
1. Run the client.

    ```shell script
    go run clien/gClient.go
    ```

### Further possible improvements

- parsing complex strings
- implementing additional math equations
✅ error handling 



