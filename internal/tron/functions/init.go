package functions

import (
	"errors"
	"fmt"
	"github.com/wenpiner/tron-scan/internal/tron/functions/params"
	"github.com/wenpiner/tron-scan/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/sha3"
)

type FunctionHandle interface {
	HandleMessage(id string, num uint64, hash string, timestamp int64, c *types.TriggerSmartContract) (*types.MQTransactionMessage, error)
	// Params 返回参数名称和实际数据大小
	Params() []params.FunctionParam
	// FunctionName 函数名称
	FunctionName() string
}

var functions = map[string]FunctionHandle{}

func RegisterFunction(handle FunctionHandle) {
	//
	var args string
	for i := range handle.Params() {
		args += handle.Params()[i].Type
		if i != len(handle.Params())-1 {
			args += ","
		}
	}

	name := handle.FunctionName() + "(" + args + ")"
	sign := GetFunctionSignature(name)
	n := fmt.Sprintf("%x", sign)
	functions[n] = handle
	logx.Infof("register function name:%s sign:%s", name, n)
}

func HandleMessage(id string, num uint64, hash string, timestamp int64, c *types.TriggerSmartContract) (*types.MQTransactionMessage, error) {
	if handle, ok := functions[c.DataInfo.FunctionName]; ok {
		return handle.HandleMessage(id, num, hash, timestamp, c)
	}
	return nil, nil

}

// Keccak256 计算 Keccak-256 哈希
func Keccak256(data []byte) []byte {
	hash := sha3.NewLegacyKeccak256()
	hash.Write(data)
	return hash.Sum(nil)
}

// GetFunctionSignature 计算函数签名
func GetFunctionSignature(functionPrototype string) [4]byte {
	// 计算 Keccak-256 哈希
	hash := Keccak256([]byte(functionPrototype))
	// 截取前 4 个字节
	var signature [4]byte
	copy(signature[:], hash[:4])
	return signature
}

func HandleFunction(data string) ([]interface{}, error) {
	// 截取函数签名,截取前判读长度是否足够 0xa9059cbb
	if len(data) < 10 {
		return nil, errors.New("data length is not enough")
	}

	// 如果开始为0x,则去掉0x
	if data[:2] == "0x" {
		data = data[2:]
	}

	signature := data[:8]

	// 截取data
	data = data[8:]

	if handle, ok := functions[signature]; ok {
		// 循环参数
		var result []interface{}
		for _, param := range handle.Params() {
			value, err := param.GetValue(data)
			if err != nil {
				return nil, err
			}
			result = append(result, value)
			// 截取字符串
			data = data[param.Len:]
		}
		return result, nil
	}
	return nil, nil
}
