package eth

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Wallet struct {
	privateKey *ecdsa.PrivateKey
	PublicKey  common.Address
}

func InitWallet(privateHexKeys string) *Wallet {
	privateKey, err := crypto.HexToECDSA(privateHexKeys)
	if err != nil {
		panic(err)
	}
	crypto.GenerateKey()

	return &Wallet{
		privateKey: privateKey,
		PublicKey:  crypto.PubkeyToAddress(privateKey.PublicKey),
	}
}

func (w *Wallet) GetOpts(client *ethclient.Client, nonce *big.Int) (*bind.TransactOpts, error) {
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return nil, err
	}

	opts, err := bind.NewKeyedTransactorWithChainID(w.privateKey, chainID)
	if err != nil {
		return nil, err
	}
	if nonce != nil {
		opts.Nonce = nonce
	} else {
		var uint64Nonce uint64
		uint64Nonce, err = client.NonceAt(context.Background(), opts.From, nil)
		if err != nil {
			return nil, err
		}
		opts.Nonce = new(big.Int)
		opts.Nonce.SetUint64(uint64Nonce)
	}
	opts.Context = context.TODO()
	return opts, nil
}
