package dump

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

type BvmDumpAccount struct {
	Address  common.Address         `json:"address"`
	Code     string                 `json:"code"`
	CodeHash string                 `json:"codeHash"`
	Storage  map[common.Hash]string `json:"storage"`
	ABI      abi.ABI                `json:"abi"`
	Nonce    uint64                 `json:"nonce"`
}

type BvmDump struct {
	Accounts map[string]BvmDumpAccount `json:"accounts"`
}
