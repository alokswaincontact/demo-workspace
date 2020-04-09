/*
 * This code tests the Hello World gRPC Server that communicates with a gRPC
 * Client and returns a "Hello World" string in reply.
 *
 * Reference:
 * The implementation contained herein has been adapted from gRPC example code.
 * https://grpc.io/docs/quickstart/go/
 */

package main

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	helloworld "github.com/alokswaincontact/demo-workspace/helloworld"
)


// rpcMsg implements the gomock.Matcher interface
type rpcMsg struct {
	msg proto.Message
}


func (r *rpcMsg) Matches(msg interface{}) bool {
	m, ok := msg.(proto.Message)
	if !ok {
		return false
	}
	return proto.Equal(m, r.msg)
}


func (r *rpcMsg) String() string {
	return fmt.Sprintf("is %s", r.msg)
}

// Mock of GreeterClient interface
type MockGreeterClient struct {
	ctrl     *gomock.Controller
	recorder *_MockGreeterClientRecorder
}

// Recorder for MockGreeterClient (not exported)
type _MockGreeterClientRecorder struct {
	mock *MockGreeterClient
}

func NewMockGreeterClient(ctrl *gomock.Controller) *MockGreeterClient {
	mock := &MockGreeterClient{ctrl: ctrl}
	mock.recorder = &_MockGreeterClientRecorder{mock}
	return mock
}

func (_m *MockGreeterClient) EXPECT() *_MockGreeterClientRecorder {
	return _m.recorder
}

func (_m *MockGreeterClient) SayHello(_param0 context.Context, _param1 *helloworld.HelloRequest, _param2 ...grpc.CallOption) (*helloworld.HelloReply, error) {
	_s := []interface{}{_param0, _param1}
	for _, _x := range _param2 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "SayHello", _s...)
	ret0, _ := ret[0].(*helloworld.HelloReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}


func (_mr *_MockGreeterClientRecorder) SayHello(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SayHello", _s...)
}

func TestSayHello(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockGreeterClient := NewMockGreeterClient(ctrl)
	req := &helloworld.HelloRequest{Name: "unit_test"}
	mockGreeterClient.EXPECT().SayHello(
		gomock.Any(),
		&rpcMsg{msg: req},
	).Return(&helloworld.HelloReply{Message: "Mocked Interface"}, nil)
	testSayHello(t, mockGreeterClient)
}

func testSayHello(t *testing.T, client helloworld.GreeterClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := client.SayHello(ctx, &helloworld.HelloRequest{Name: "unit_test"})
	if err != nil || r.Message != "Mocked Interface" {
		t.Errorf("mocking failed")
	}
	t.Log("Reply : ", r.Message)
}
