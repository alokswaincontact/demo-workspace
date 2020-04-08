# VDart Digital's Coding Challenge using GoLang & gRPC
This contains the code for the VDart Digital's Coding Challenge using GoLang and gRPC.


# Execution
The implementation has two parts the server and the client.

It is recommended to start the server first. 
You can run the server using command shown below
**go run main_server.go**

```
~/go/src/github.com/alokswaincontact/demo-workspace$ go run main_server.go 
2020/04/08 17:03:49 Hello World gRPC Server Started
2020/04/08 17:03:54 Message Received from client.
^Csignal: interrupt
```

Then you can start the client to connect to the server as shown below.
**go run main_client.go**

```
~/go/src/github.com/alokswaincontact/demo-workspace$ go run main_client.go 
2020/04/08 17:00:13 Message Received: Hello World
```
The client connects to the server and receives a *Hello World* message.


# Reference
The implementation contained herein has been adapted from gRPC example code.
 *Go Quick Start
    https://grpc.io/docs/quickstart/go/
