package pb

import (
	"github.com/micro/go-micro"
	"testgopb/conf"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/selector"
	"github.com/micro/go-micro/server"
	"github.com/micro/go-micro/client"
)

const (
	ResolveSVCName = "testResolve"
)


func NewService(csm *conf.ConfigServiceMicro, opts ...micro.Option) micro.Service{
	reg := consul.NewRegistry(registry.Addrs(csm.ConsulAddress))
	sel := selector.NewSelector(selector.Registry(reg))

	serverOpts := []server.Option{}
	serverOpts = append(serverOpts, server.Registry(reg))
	serverOpts = append(serverOpts, server.Name(csm.ServiceName))
	serverOpts = append(serverOpts, server.Address(csm.ServiceAddress))

	clientOpts := []client.Option{}
	clientOpts = append(clientOpts, client.Selector(sel))

	options := []micro.Option{
		micro.Registry(reg),
		micro.Server(server.NewServer(serverOpts...)),
		micro.Client(client.NewClient(clientOpts...)),
	}

	myService := micro.NewService(options...)
	return myService
}