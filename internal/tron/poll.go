package tron

import (
	"errors"
	"math/rand"
	"strings"
	"time"
	"tronScan/internal/types"

	"github.com/go-resty/resty/v2"
	"github.com/zeromicro/go-zero/core/logx"
)

type Poll struct {
	// 块消息通道
	blockChan chan types.Block
	// HTTP客户端
	client *resty.Client
	// done
	done chan struct{}
	// 当前块高
	currentBlockNum int64
	// apis
	apis []string
}

func (t *Poll) randKey() string {
	index := rand.Intn(len(t.apis))

	if len(t.apis) == 1 {
		index = 0
	}
	return t.apis[index]
}

// NewTronPoll creates a new Poll instance
func NewTronPoll(apiKey string, blockChan chan types.Block) *Poll {
	keys := strings.Split(apiKey, ",")
	if len(keys) == 0 {
		logx.Error("未配置API KEY")
		return nil
	}
	client := resty.New()
	// 设置超时时间
	client.SetTimeout(3 * time.Second)
	return &Poll{
		client:    client,
		blockChan: blockChan,
		apis:      keys,
	}
}

// Start StartPolling starts polling the TronScan API
func (t *Poll) Start() {

	if t.currentBlockNum <= 0 {
		var nowBlock *types.Block
		var errCount int
		for nowBlock == nil {
			if errCount >= 10 {
				logx.Error("获取当前块失败次数过多，退出轮训")
				return
			}
			var err error
			nowBlock, err = t.getNowBlock()
			if nowBlock == nil {
				errCount++
				logx.Infof("获取当前块失败，重试次数：%d 错误信息:%v", errCount, err)
				time.Sleep(500 * time.Millisecond)
			}
		}
		t.blockChan <- *nowBlock
		t.currentBlockNum = int64(nowBlock.BlockHeader.RawData.Number + 1)
	}

	for {
		select {
		case <-t.done:
			return
		default:
			// Poll the TronScan API
			t.onRun()
		}
	}
}

func (t *Poll) onRun() {
	block, err := t.getBlock(t.currentBlockNum)
	if err == nil {
		t.blockChan <- *block
		t.currentBlockNum = int64(block.BlockHeader.RawData.Number + 1)
	}
	time.Sleep(1000 * time.Millisecond)
}

// Stop poll polls the TronScan API
func (t *Poll) Stop() {
	close(t.done)
}

// 获取当前块
func (t *Poll) getBlock(num int64) (*types.Block, error) {
	logx.Infof("获取块:%d", num)
	request := map[string]interface{}{
		"visible": true,
		"num":     num,
	}
	url := "https://api.trongrid.io/wallet/getblockbynum"
	var block types.Block
	response, err := t.client.SetHeader("TRON-PRO-API-KEY", t.randKey()).R().SetResult(&block).SetBody(request).Post(url)
	if err != nil {
		return nil, err
	}
	if response.IsSuccess() {
		if block.BlockHeader.RawData.Number == 0 {
			return nil, errors.New("block not found")
		}
		return &block, nil
	} else {
		return nil, errors.New(response.String())
	}
}

func (t *Poll) getNowBlock() (*types.Block, error) {
	request := map[string]interface{}{
		"visible": true,
	}
	url := "https://api.trongrid.io/wallet/getnowblock"
	var block types.Block
	response, err := t.client.SetHeader("TRON-PRO-API-KEY", t.randKey()).R().SetResult(&block).SetBody(request).Post(url)
	if err != nil {
		return nil, err
	}
	if response.IsSuccess() {
		if block.BlockHeader.RawData.Number == 0 {
			return nil, errors.New("block not found")
		}
		return &block, nil
	} else {
		return nil, errors.New(response.String())
	}
}
