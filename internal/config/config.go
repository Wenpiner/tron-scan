package config

import (
	"github.com/wenpiner/rabbitmq-go/conf"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	RabbitConf conf.RabbitConf
	APIKey     string
	MQ         struct {
		BlockExchangeName       string
		BlockRouteKey           string
		TransactionExchangeName string
		TransactionRouteKey     string
	}
}
