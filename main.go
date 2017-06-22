package main

import (
	"flag"
	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/go-redis/redis"
	"github.com/katsew/core"
	"github.com/katsew/kuji"
	"github.com/katsew/kuji-redis"
	"github.com/katsew/kuji-server/services/kuji"
	"os"
)

const (
	APP_MODE     = "MODE"
	DEFAULT_MODE = "thrift"
	MODE_GRPC    = "grpc"
	MODE_THRIFT  = "thrift"
)

func main() {

	config := core.LoadConfigs()

	var mode string
	mode = os.Getenv("MODE")

	if mode == "" {
		flag.StringVar(&mode, APP_MODE, DEFAULT_MODE, "Get run mode")
	}

	switch mode {
	case MODE_GRPC:

		panic("Not implemented.")

	case MODE_THRIFT:

		protocolFactory := thrift.NewTJSONProtocolFactory()
		transportFactory := thrift.NewTTransportFactory()
		transport, err := thrift.NewTServerSocket("localhost:" + config.Port)
		if err != nil {
			panic(err)
		}

		processor := thrift.NewTMultiplexedProcessor()
		strategy := kuji_redis.NewSimpleStrategy(&redis.Options{
			Addr:     "127.0.0.1:6379",
			Password: "",
			DB:       0,
		})
		config := kuji.KujiStrategyConfig{
			Strategy: strategy,
			FailOver: nil,
		}
		kujiProcessor := services.NewThriftGachaService(transport, transportFactory, protocolFactory, config)
		processor.RegisterProcessor(
			"KujiService",
			kujiProcessor,
		)
		server := thrift.NewTSimpleServer4(processor, transport, transportFactory, protocolFactory)
		server.Serve()

	}
}
