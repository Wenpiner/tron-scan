package tron

import (
	"golang.org/x/crypto/sha3"
	"testing"
)

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

func TestSignature(t *testing.T) {
	// 代币转账函数原型
	functionPrototype := "transfer(address,uint256)"
	// 计算函数签名
	signature := GetFunctionSignature(functionPrototype)
	t.Logf("Function Signature: 0x%x", signature)
	// Output:
	// Function Signature: 0xa9059cbb
}
