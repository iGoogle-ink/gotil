package micro

import (
	"fmt"

	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/server"
)

func InitServer(name, version string, registry *EtcdRegistry, fn func(s server.Server)) {
	var (
		s    micro.Service
		opts []micro.Option
	)
	opts = []micro.Option{
		micro.Name(name),
		micro.WrapHandler(LogWrapper),
		micro.Version(version),
	}

	if registry != nil {
		opts = append(opts, micro.Registry(etcdRegistry(registry)))
	}
	s = micro.NewService(opts...)

	fn(s.Server())

	go func() {
		if err := s.Run(); err != nil {
			panic(fmt.Sprintf("[%s] micro server run error(%+v).", name, err))
		}
	}()
}
