package types

import (
	"strconv"
)

// AccountCreateContract, TransferContract, TransferAssetContract, VoteAssetContract, VoteWitnessContract,
// WitnessCreateContract, AssetIssueContract, WitnessUpdateContract, ParticipateAssetIssueContract,
// AccountUpdateContract, FreezeBalanceContract, UnfreezeBalanceContract, WithdrawBalanceContract,
// UnfreezeAssetContract, UpdateAssetContract, ProposalCreateContract, ProposalApproveContract, ProposalDeleteContract,
// SetAccountIdContract, CustomContract, CreateSmartContract, TriggerSmartContract, GetContract, UpdateSettingContract,
// ExchangeCreateContract, ExchangeInjectContract, ExchangeWithdrawContract, ExchangeTransactionContract,
// UpdateEnergyLimitContract, AccountPermissionUpdateContract, ClearABIContract, UpdateBrokerageContract,
// ShieldedTransferContract, MarketSellAssetContract, MarketCancelOrderContract, FreezeBalanceV2Contract,
// UnfreezeBalanceV2Contract, WithdrawExpireUnfreezeContract, DelegateResourceContract, UnDelegateResourceContract,
// CancelAllUnfreezeV2Contract
type TransferContract struct {
	Amount       int64       `json:"amount"`
	ToAddress    TronAddress `json:"to_address"`
	OwnerAddress TronAddress `json:"owner_address"`
}

func NewTransferContract(value map[string]interface{}) *TransferContract {
	contract := &TransferContract{}
	contract.Amount = int64(value["amount"].(float64))
	contract.ToAddress = AddressByBase58(value["to_address"].(string))
	contract.OwnerAddress = AddressByBase58(value["owner_address"].(string))
	return contract
}

type TriggerSmartContract struct {
	ContractAddress TronAddress `json:"contract_address"`
	OwnerAddress    TronAddress `json:"owner_address"`
	Data            string      `json:"data"`
	DataInfo        struct {
		FunctionName string      `json:"-"`
		ToAddress    TronAddress `json:"-"`
		Amount       int64       `json:"-"`
	} `json:"-"`
}

func NewTriggerSmartContract(value map[string]interface{}) *TriggerSmartContract {
	var ok bool
	contract := &TriggerSmartContract{}
	cAddr, ok := value["contract_address"].(string)
	if !ok {
		return nil
	}
	contract.ContractAddress = AddressByBase58(cAddr)
	oAddr, ok := value["owner_address"].(string)
	if !ok {
		return nil
	}
	contract.OwnerAddress = AddressByBase58(oAddr)
	contract.Data, ok = value["data"].(string)
	if !ok {
		return nil
	}
	// 截取0 - 8位

	contract.DataInfo.FunctionName = contract.Data[0:8]
	// 如果FunctionName == a9059cbb
	if contract.DataInfo.FunctionName == "a9059cbb" {
		// 判断数据长度是否合法
		if len(contract.Data) < 136 {
			return nil
		}

		// 截取8 - 72位
		contract.DataInfo.ToAddress = AddressByHex(contract.Data[8:72])
		// 截取72位后面的64位
		amountHex := contract.Data[72 : 72+64]
		// 将16进制转换为10进制
		parseInt, err := strconv.ParseInt(amountHex, 16, 64)
		if err == nil {
			contract.DataInfo.Amount = parseInt
		}
	}
	return contract
}
