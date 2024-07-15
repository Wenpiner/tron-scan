package types

import (
	"encoding/hex"

	"github.com/fbsobreira/gotron-sdk/pkg/address"
)

type Block struct {
	BlockID     string `json:"blockID"`
	BlockHeader struct {
		WitnessSignature string `json:"witness_signature"`
		RawData          struct {
			Number         uint64      `json:"number"`
			TxTrieRoot     string      `json:"txTrieRoot"`
			WitnessAddress TronAddress `json:"witness_address"`
			ParentHash     string      `json:"parentHash"`
			Version        int64       `json:"version"`
			Timestamp      uint64      `json:"timestamp"`
		} `json:"raw_data"`
	} `json:"block_header"`

	Transactions []Transaction `json:"transactions"`
}

type TronAddress struct {
	Address address.Address
}

// UnmarshalJSON 实现 UnmarshalJSON 方法
func (t TronAddress) UnmarshalJSON(data []byte) (err error) {
	// 赋值给Address
	t.Address = data
	return
}

// MarshalJSON 实现 MarshalJSON 方法
func (t TronAddress) MarshalJSON() ([]byte, error) {
	return []byte(t.Address.String()), nil
}

func (t TronAddress) String() string {
	return t.Address.String()
}

func AddressByBase58(val string) TronAddress {
	toAddress, err := address.Base58ToAddress(val)
	if err != nil {
		return TronAddress{
			Address: nil,
		}
	}
	return TronAddress{
		Address: toAddress,
	}
}

// 除去bytes前面的0值
func removeZeroBytes(bytes []byte) []byte {
	for i, b := range bytes {
		if b != 0 {
			return bytes[i:]
		}
	}
	return nil
}

func AddressByHex(val string) TronAddress {
	decodeString, err := hex.DecodeString(val)
	if err != nil {
		return TronAddress{}
	}
	newByte := removeZeroBytes(decodeString)
	if len(newByte) == 0 {
		return TronAddress{}
	}
	// 判断开头是否为41,如果不是则加在最前面
	if newByte[0] != 0x41 {
		newByte = append([]byte{0x41}, newByte...)
	}

	return TronAddress{
		Address: newByte,
	}
}
