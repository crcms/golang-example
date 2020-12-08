package main

import (
	"context"
	"fmt"
	"golang-example/grpc/sample/example"
	"google.golang.org/grpc"
	"net"
	"sync"
)

type ExampleRequestResponseService struct {
	example.UnimplementedExampleRequestResponseServer
}

var wg1 sync.WaitGroup

func (e ExampleRequestResponseService) ExampleCall(ctx context.Context, req *example.Request) (*example.Response, error) {
	fmt.Sprintf("%#v\n", req)
	return &example.Response{
		Code:    200,
		Message: "处理成功",
		Data: []*example.ResponseDataList{
			{
				Title:   "t1",
				Content: "c1",
				Date:    "d1",
				UserId:  0,
			},
			{
				Title:   "t2",
				Content: "c2",
				Date:    "d2",
				UserId:  0,
			},
		},
	}, nil
}

func main() {
	wg1.Add(2)
	go server1()
	go client1()
	wg1.Wait()
}

func client1() {
	defer wg1.Done()
	conn, err := grpc.Dial("127.0.0.1:22188", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	var address = `address`
	client := example.NewExampleRequestResponseClient(conn)
	response, err := client.ExampleCall(context.Background(), &example.Request{
		Loves:    []int32{1, 2},
		Name:     "name",
		Age:      10,
		Address:  &address,
		Selected: false,
		Other: map[string]string{
			`k1`: `v1`,
			`k2`: `v2`,
		},
		Status: example.Status_disable,
		Avatar: &example.Request_ImageUrl{
			ImageUrl: "http://example.com",
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", response)
}

func server1() {
	defer wg1.Done()
	listen, err := net.Listen("tcp", ":22188")
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer()
	example.RegisterExampleRequestResponseServer(grpcServer, new(ExampleRequestResponseService))
	if err := grpcServer.Serve(listen); err != nil {
		panic(err)
	}
}
