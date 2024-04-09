package tron

import (
	"context"
	"encoding/json"
	"tronScan/internal/config"
	"tronScan/internal/svc"
	"tronScan/internal/types"

	"github.com/rabbitmq/amqp091-go"
	"github.com/wenpiner/rabbitmq-go"
	"github.com/zeromicro/go-zero/core/logx"
)

type HandleBlock struct {
	// 块消息通道
	blockChan chan types.Block
	// done
	done chan struct{}
	// 消息队列
	Rabbit *rabbitmq.RabbitMQ
	// 配置信息
	config *config.Config
	// svc
	svc *svc.ServiceContext
}

// NewHandleBlock creates a new HandleBlock instance
func NewHandleBlock(blockChan chan types.Block, Rabbit *rabbitmq.RabbitMQ, config *config.Config, svc *svc.ServiceContext) *HandleBlock {
	return &HandleBlock{
		blockChan: blockChan,
		Rabbit:    Rabbit,
		config:    config,
		svc:       svc,
	}
}

// Start StartPolling starts polling the TronScan API
func (t *HandleBlock) Start() {
	for {
		select {
		case <-t.done:
			return
		case block := <-t.blockChan:
			// Handle the block
			go t.Handle(block)
		}
	}
}

func (t *HandleBlock) Handle(block types.Block) {
	// 更新健康信息
	t.svc.UpdateBlockInfo(block)

	channel, err := t.Rabbit.Channel()
	if err != nil {
		logx.Errorf("获取消息队列通道失败: %v", err)
		return
	}
	defer channel.Close()

	// 发送块消息
	blockMessage := &types.BlockMessage{
		BlockNum:       block.BlockHeader.RawData.Number,
		BlockHash:      block.BlockID,
		Timestamp:      block.BlockHeader.RawData.Timestamp,
		WitnessAddress: block.BlockHeader.RawData.WitnessAddress.Address.String(),
	}
	baseMessage := &types.Message{
		Type: "block",
		Data: blockMessage,
	}

	blockByte, err := json.Marshal(baseMessage)
	if err != nil {
		logx.Errorf("序列化消息失败: %v", err)
		return
	}
	logx.Infof("处理块:%d", block.BlockHeader.RawData.Number)
	err = channel.PublishWithContext(
		context.Background(), t.config.MQ.BlockExchangeName, t.config.MQ.BlockRouteKey, false, false, amqp091.Publishing{
			ContentType: "application/json",
			Body:        blockByte,
		},
	)
	if err != nil {
		logx.Errorf("发送块消息失败: %v", err)
		return
	}

	// 检查所有交易
	for _, tx := range block.Transactions {
		if len(tx.Rets) == 1 {
			if tx.Rets[0].ContractRet == "SUCCESS" {
				baseMessageTx := &types.Message{
					Type: "transaction",
				}
				// 成功交易
				contract := tx.RawData.Contract[0]
				switch contract.Type {
				case "TransferContract":
					c := t.HandleTransferContract(block.BlockHeader.RawData.Number, block.BlockID, tx)
					if c != nil {
						baseMessageTx.Data = c
					}
					break
				case "TriggerSmartContract":
					t := t.HandleTriggerSmartContract(block.BlockHeader.RawData.Number, block.BlockID, tx)
					if t != nil {
						baseMessageTx.Data = t
					}
					break
				default:
					continue
				}
				transactionByte, _ := json.Marshal(baseMessageTx)
				err := channel.PublishWithContext(
					context.Background(), t.config.MQ.TransactionExchangeName, t.config.MQ.TransactionRouteKey, false, false, amqp091.Publishing{
						ContentType: "application/json",
						Body:        transactionByte,
					},
				)
				if err != nil {
					logx.Errorf("发送交易消息失败: %v", err)
				}
				continue
			}
		}
	}
}

func (t *HandleBlock) HandleTransferContract(blockNum uint64, blockHash string, transaction types.Transaction) *types.MQTransactionMessage {
	contractVal := types.NewTransferContract(transaction.RawData.Contract[0].Parameter.Value)
	amount := contractVal.Amount
	fromAddr := contractVal.OwnerAddress.String()
	toAddr := contractVal.ToAddress.String()
	return &types.MQTransactionMessage{
		Hash:      transaction.TxID,
		BlockNum:  blockNum,
		Amount:    amount,
		Contract:  "",
		BlockHash: blockHash,
		FromAddr:  fromAddr,
		ToAddr:    toAddr,
		Datetime:  transaction.RawData.Timestamp,
	}
}

func (t *HandleBlock) HandleTriggerSmartContract(blockNum uint64, blockHash string, transaction types.Transaction) *types.MQTransactionMessage {
	contractVal := types.NewTriggerSmartContract(transaction.RawData.Contract[0].Parameter.Value)
	// 过滤TRC20交易
	if contractVal.DataInfo.FunctionName != "a9059cbb" {
		return nil
	}
	amount := contractVal.DataInfo.Amount
	fromAddr := contractVal.OwnerAddress.String()
	toAddr := contractVal.DataInfo.ToAddress.String()
	return &types.MQTransactionMessage{
		Hash:      transaction.TxID,
		BlockNum:  blockNum,
		Amount:    amount,
		Contract:  contractVal.ContractAddress.String(),
		BlockHash: blockHash,
		FromAddr:  fromAddr,
		ToAddr:    toAddr,
		Datetime:  transaction.RawData.Timestamp,
	}
}

func (t *HandleBlock) Stop() {
	close(t.blockChan)
}
