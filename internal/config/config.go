package config

import (
	"github.com/wenpiner/rabbitmq-go/conf"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	RabbitConf conf.RabbitConf
	APIKey     string
	BaseURL    string
	MQ         struct {
		BlockExchangeName       string
		BlockRouteKey           string
		TransactionExchangeName string
		TransactionRouteKey     string
	}
	// 需要监听的合约函数
	TriggerSmartContractFunctions map[string]bool
	TransferContract              bool
}
