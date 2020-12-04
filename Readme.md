# gRPC calculator

This is a gRPC implementation of a calculator.

1. initialise go.mod

    ```shell script
    go mod init testAssignmment
    ```

2. create a description `.proto` file, which will describe the communication between client and server and the methods, that will be called.


## Generating the gRPC client and server interface from `.proto` file
1. Install instrumments for OS X with Homebrew
    ```shell script
    brew install protobuf
    ```

2. installing `protoc-gen-go` according to the description on [StackOverflow](https://stackoverflow.com/questions/57700860/protoc-gen-go-program-not-found-or-is-not-executable)
    ```shell script
    go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
    ```

3. Export GOPATH
    ```shell script
    echo 'export GOPATH=$HOME/Go' >> $HOME/.bashrc
    source $HOME/.bashrc
    ```
Either if you use `zsh` instead write and source the `.zshrc`  

4. Generating the `calc.pb.go` file with for gRPC
    ```shell script
    protoc --go_out=plugins=grpc:./gen --go_opt=paths=source_relative calc.proto
    ```
[//]: <Like> (<--protoc -I . --go_out=/Users/chernykh_alexander/go/src/tutorial/ calc.proto-->)

5. Generated `calc.pb.go` will contain protocols and data structures, that will be used during the gRPC calls.


### Get Started 

An example gRPC calculator usage could look like the following. By running the server in a first shell user will be noootified, that it is serving and waiting for calls from client. By starting the client in the second shell window, user will be prompted to enter a simple equation, f.e `1 + 2 `. Additional allowed operations are `-`,`*` and `/`.


1. In the first shell, start the server `gServer.go`
    ```shell script
    go run gServer.go
    ```

2. In the second shell, start the client `gClient.go`

    ```shell script
    go run gClient.go
    ```
3. 