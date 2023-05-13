package src

import (
	"context"
	"encoding/json"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type TokenContract struct {
	Transfer func(options *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error)
}

type TransferRequest struct {
	PrivateKey      string `json:"privateKey"`
	ContractAddress string `json:"contractAddress"`
	ToAddress       string `json:"toAddress"`
	Amount          string `json:"amount"`
}

func TransferToken(w http.ResponseWriter, r *http.Request) {
	var req TransferRequest
	json.NewDecoder(r.Body).Decode(&req)

	client, _ := ethclient.Dial("https://api.avax-test.network/ext/bc/C/rpc")

	privateKey, _ := crypto.HexToECDSA(req.PrivateKey)
	fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)

	nonce, _ := client.PendingNonceAt(context.Background(), fromAddress)
	gasPrice, _ := client.SuggestGasPrice(context.Background())

	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(43113))
	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasLimit = uint64(300000)
	auth.GasPrice = gasPrice

	tokenAddress := common.HexToAddress(req.ContractAddress)
	instance, _ := NewMain(tokenAddress, client)
	toAddress := common.HexToAddress(req.ToAddress)
	tokenAmount := new(big.Int)

	tokenAmount.SetString(req.Amount, 10)

	weiMultiplier := big.NewInt(1e18)

	tokenAmountInWei := new(big.Int)
	tokenAmountInWei.Mul(tokenAmount, weiMultiplier)

	result, err := instance.Transfer(auth, toAddress, tokenAmountInWei)
	print(result)
	if err != nil {
		http.Error(w, "Error transferring token: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
