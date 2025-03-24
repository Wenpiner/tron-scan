package types

type (
	Message struct {
		// 消息类型
		Type string `json:"type"`
		// 消息内容
		Data interface{} `json:"data"`
	}

	BlockMessage struct {
		// 区块高度
		BlockNum uint64 `json:"block_num"`
		// 区块哈希
		BlockHash string `json:"block_hash"`
		// 时间戳
		Timestamp uint64 `json:"timestamp"`
		// 块生产者
		WitnessAddress string `json:"witness_address"`
	}

	MQTransactionMessage struct {
		Hash         string        `json:"hash"`
		BlockNum     uint64        `json:"blockNum"`
		Amount       int64         `json:"amount"`
		Contract     string        `json:"contract"`
		BlockHash    string        `json:"blockHash"`
		FromAddr     string        `json:"fromAddr"`
		ToAddr       string        `json:"toAddr"`
		Datetime     int64         `json:"datetime"`
		FunctionName string        `json:"functionName"`
		Values       []interface{} `json:"values"`
	}
)
