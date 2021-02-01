# gRPC calculator

This is a gRPC implementation of simple math tasks.
![gRPC Calculator](img/gRpc_2.png)

<!-- TABLE OF CONTENTS -->
<details open="open">
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
    </li>
    <li><a href="#further-possible-improvements">Further Possible Improvements</a></li>
  </ol>
</details>

## About The Project 

An example gRPC calculator usage could look like the following. By running the server in a first shell user will be notified, that it is serving and waiting for calls from client. By starting the client in the second shell window, user will be prompted to enter a simple equation, f.e `1 + 2 `. All allowed operations are `+`, `-`,`*` and `/`.

### Built With

* [Go](https://golang.org)
* [gRPC](https://grpc.io/docs/languages/go/quickstart/)


## Getting Started


1. Install dependencies
    ```shell script
    go mod download
    ```

1. In the first shell, start the server `gServer.go`
    ```shell script
    go run gServer.go
    ```

1. In the second shell, start the client `gClient.go`

    ```shell script
    go run gClient.go
    ```
1. Input some math operation
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

## Use container-container communication

Made according to this [description](https://stackoverflow.com/questions/44642241/use-grpc-to-communicate-between-containers)

1. Create network 

    ```shell script
    docker network create --driver bridge go-net
    ```
   
1. Build and run the server. It is important to provide here a name for the server, so the client knows to whom he has to speak

    ```shell script
    docker build -t server -f server/Dockerfile .
    docker run -d --network go-net  --name  server server
    ```

1. Build and run the client via a shell script, that simulates a simple input.

    ```shell script
    docker build -t client -f client/Dockerfile .
    docker run -i --log-driver=none -a stdin -a stdout -a stderr --network go-net -p 8080:8080 client
    ```

### Further possible improvements


- parsing complex strings
- implementing additional math equations
- ✅ error handling



