package eth

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetGasLimit(client *ethclient.Client, from, to common.Address, value *big.Int, input []byte) (uint64, error) {
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return 0, err
	}

	msg := ethereum.CallMsg{
		From:      from,
		To:        &to,
		GasPrice:  gasPrice,
		GasTipCap: nil,
		GasFeeCap: nil,
		Value:     value,
		Data:      input,
	}
	return client.EstimateGas(context.Background(), msg)
}
