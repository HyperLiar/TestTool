package main

import "net/rpc"

const HelloServiceName = "path/to/pkg.HelloeService"

// 定义规范
type HelloServiceInterface = interface {
	Hello(request string, reply *string) error
}

func RegisterHelloService(svc HelloServiceInterface) error {
	return rpc.RegisterName(HelloServiceName, svc)
}

// 进一步封装
type HelloServiceClient struct {
	*rpc.Client
}

var _ HelloServiceInterface = (*HelloServiceClient)(nil)

func DialHelloService(network, address string) (*HelloServiceClient, error) {
	c, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}

	return &HelloServiceClient{Client:c}, nil
}

func (p *HelloServiceClient) Hello (request string, reply *string) error {
	return p.Client.Call(HelloServiceName+".Hello", request, reply)
}