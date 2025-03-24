package main

import (
	"flag"
	"fmt"
	"github.com/wenpiner/rabbitmq-go"
	"github.com/wenpiner/tron-scan/internal/tron"
	"github.com/wenpiner/tron-scan/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"

	"github.com/wenpiner/tron-scan/internal/config"
	"github.com/wenpiner/tron-scan/internal/handler"
	"github.com/wenpiner/tron-scan/internal/svc"
	_ "github.com/wenpiner/tron-scan/internal/tron/functions"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/tron-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c, conf.UseEnv())

	sg := service.ServiceGroup{}
	defer sg.Stop()

	server := rest.MustNewServer(c.RestConf)
	sg.Add(server)

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	blockChan := make(chan types.Block)
	// 初始化Rabbit
	mq := rabbitmq.NewRabbitMQ(c.RabbitConf)

	handle := tron.NewHandleBlock(blockChan, mq, &c, ctx)
	sg.Add(handle)

	poll := tron.NewTronPoll(c.APIKey, blockChan)
	sg.Add(poll)

	// 打印开启的方法
	for s, v := range c.TriggerSmartContractFunctions {
		logx.Infof("tron-smart-contract function: %s open:%v", s, v)
	}

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	sg.Start()
}
