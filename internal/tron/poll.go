package tron

import (
	"errors"
	"fmt"
	"github.com/wenpiner/tron-scan/internal/types"
	"math/rand"
	"net"
	"net/http"
	"strings"
	"time"

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
	client.SetTimeout(5 * time.Second)
	// 设置重试次数
	client.SetRetryCount(2)
	// 设置重试等待时间
	client.SetRetryWaitTime(300 * time.Millisecond)
	// 设置连接和读取超时
	client.SetTransport(&http.Transport{
		DialContext: (&net.Dialer{
			Timeout: 2 * time.Second, // 设置连接超时
		}).DialContext,
		ResponseHeaderTimeout: 5 * time.Second, // 设置响应头超时
	})

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
	time.Sleep(300 * time.Millisecond)
}

// Stop poll polls the TronScan API
func (t *Poll) Stop() {
	close(t.done)
}

// 获取当前块
func (t *Poll) getBlock(num int64) (*types.Block, error) {
	request := map[string]interface{}{
		"visible": true,
		"num":     num,
	}
	url := fmt.Sprintf(`https://go.getblock.io/%s/wallet/getblockbynum`, t.randKey())
	var block types.Block
	response, err := t.client.SetHeader("TRON-PRO-API-KEY", t.randKey()).R().SetResult(&block).SetBody(request).Post(url)
	if err != nil {
		logx.Errorf("获取块失败:%v", err)
		return nil, err
	}
	if response.IsSuccess() {
		logx.Infof("请求块:%d 当前块:%d", num, block.BlockHeader.RawData.Number)
		if block.BlockHeader.RawData.Number == 0 {
			return nil, errors.New("block not found")
		}
		return &block, nil
	} else {
		logx.Errorf("获取块失败:%v", response.String())
		return nil, errors.New(response.String())
	}
}

func (t *Poll) getNowBlock() (*types.Block, error) {
	request := map[string]interface{}{
		"visible": true,
	}
	url := fmt.Sprintf(`https://go.getblock.io/%s/wallet/getnowblock`, t.randKey())
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
