# VDart Digital's Coding Challenge using GoLang & gRPC
This contains the code for the VDart Digital's Coding Challenge using GoLang and gRPC.

# Architecture
The picture below explains how the Client and Server communicate with
each other using gRPC calls.

![architecture](https://github.com/alokswaincontact/demo-workspace/blob/master/architecture.png)

The gRPC Client connects and sends a message to the gRPC Server, and the
gRPC Server replies with a message that has a string with value `Hello World`.
The gRPC Client code is implemented in `main_client.go` and the gRPC
Server code in `main_server.go`.

The Client makes a [blocking](https://godoc.org/google.golang.org/grpc#WithBlock) call to
the Server and after receipt of the message displays it and closes the connection.

The gRPC Server waits and listens for Client connection on TCP port `50051` using
`lis, err := net.Listen("tcp", ":50051")`.
We create an instance of the gRPC server using `grpc.NewServer()`.
We then register our service implementation with the gRPC server using
`pb.RegisterGreeterServer(...)`
Finally we call `Serve()` on the server with our port details to do a blocking
wait until the process is killed or `Stop()` is called.

# Execution
The implementation has two parts the server and the client.

It is recommended to start the server first.
You can run the server using command shown below
**go run main_server.go**

```shell
~/go/src/github.com/alokswaincontact/demo-workspace$ go run main_server.go
2020/04/08 17:03:49 Hello World gRPC Server Started
2020/04/08 17:03:54 Message Received from client.
^Csignal: interrupt
```

Then you can start the client to connect to the server as shown below.
**go run main_client.go**

```shell
~/go/src/github.com/alokswaincontact/demo-workspace$ go run main_client.go
2020/04/08 17:00:13 Message Received: Hello World
```
The client connects to the server and receives a `Hello World` message.

# Setup
You need a GoLang installation to run this code. You can install GoLang
compiler in MacOS using the command below.
```shell
$ brew install golang
```

To compile the `proto` files you need to protoc compiler. In MacOS 
it can be installed using
```shell
$ brew install protobuf
```

To compile and run this code you would also need gRPC and protoc-gen-go.
```shell
$ go get -u google.golang.org/grpc
$ go get -u github.com/golang/protobuf/protoc-gen-go
```

# Testing
The Server code in `main_server.go` can be tested using command
**go test main_server_test.go** 

A verbose output is shown below.
```shell
~/go/src/github.com/alokswaincontact/demo-workspace$ go test -v main_server_test.go 
=== RUN   TestSayHello
    TestSayHello: main_server_test.go:101: Reply :  Mocked Interface
    TestSayHello: testing.go:789: In Go 1.14+ you no longer need to `ctrl.Finish()` if a *testing.T is passed to `NewController(...)`
--- PASS: TestSayHello (0.00s)
PASS
ok      command-line-arguments  0.110s
~/go/src/github.com/alokswaincontact/demo-workspace$
```

# Reference
* gRPC Quick Start  
The implementation contained herein has been adapted from gRPC example code.  
https://grpc.io/docs/quickstart/go/

* gRPC Basics  
https://grpc.io/docs/tutorials/basic/go/#generating-client-and-server-code

* protoc: Protobuf compiler installation  
http://google.github.io/proto-lens/installing-protoc.html

* Google Protocol Buffers  
https://github.com/golang/protobuf
