package types

type Parameter struct {
	TypeURL string                 `json:"type_url"`
	Value   map[string]interface{} `json:"value"`
}

type Transaction struct {
	Signature  []string `json:"signature"`
	TxID       string   `json:"txID"`
	RawDataHex string   `json:"rawDataHex"`
	Rets       []struct {
		// DEFAULT, SUCCESS, REVERT, BAD_JUMP_DESTINATION, OUT_OF_MEMORY, PRECOMPILED_CONTRACT, STACK_TOO_SMALL, STACK_TOO_LARGE, ILLEGAL_OPERATION, STACK_OVERFLOW, OUT_OF_ENERGY, OUT_OF_TIME, JVM_STACK_OVER_FLOW, UNKNOWN, TRANSFER_FAILED, UNRECOGNIZED;
		ContractRet string `json:"contractRet"`
		// SUCCESS, FAILED, UNRECOGNIZED
		Ret string `json:"ret"`
	} `json:"ret"`
	RawData struct {
		RefBlockBytesHex string `json:"ref_block_bytes"`
		RefBlockHash     string `json:"ref_block_hash"`
		Expiration       int64  `json:"expiration"`
		FeeLimit         int64  `json:"fee_limit"`
		Timestamp        int64  `json:"timestamp"`
		Contract         []struct {
			Parameter Parameter `json:"parameter"`
			//AccountCreateContract, TransferContract, TransferAssetContract, VoteAssetContract, VoteWitnessContract,
			//WitnessCreateContract, AssetIssueContract, WitnessUpdateContract, ParticipateAssetIssueContract, AccountUpdateContract, FreezeBalanceContract, UnfreezeBalanceContract, WithdrawBalanceContract, UnfreezeAssetContract, UpdateAssetContract, ProposalCreateContract, ProposalApproveContract, ProposalDeleteContract, SetAccountIdContract, CustomContract, CreateSmartContract, TriggerSmartContract, GetContract, UpdateSettingContract, ExchangeCreateContract, ExchangeInjectContract, ExchangeWithdrawContract, ExchangeTransactionContract, UpdateEnergyLimitContract, AccountPermissionUpdateContract, ClearABIContract, UpdateBrokerageContract, ShieldedTransferContract, MarketSellAssetContract, MarketCancelOrderContract, FreezeBalanceV2Contract, UnfreezeBalanceV2Contract, WithdrawExpireUnfreezeContract, DelegateResourceContract, UnDelegateResourceContract, CancelAllUnfreezeV2Contract
			Type string `json:"type"`
		} `json:"contract"`
	} `json:"raw_data"`
}
