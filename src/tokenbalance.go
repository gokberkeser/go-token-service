package src

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"net/http"
)

type BalanceRequest struct {
	ContractAddress string `json:"contractAddress"`
	WalletAddress   string `json:"walletAddress"`
}

func BalanceHandler(w http.ResponseWriter, r *http.Request) {
	var req BalanceRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	client, err := ethclient.Dial("https://api.avax-test.network/ext/bc/C/rpc")
	if err != nil {
		http.Error(w, "Failed to connect to the Ethereum client: "+err.Error(), http.StatusInternalServerError)
		return
	}

	tokenAddress := common.HexToAddress(req.ContractAddress)
	instance, _ := NewMain(tokenAddress, client)

	walletAddress := common.HexToAddress(req.WalletAddress)
	balance, err := instance.BalanceOf(&bind.CallOpts{}, walletAddress)
	if err != nil {
		http.Error(w, "Error getting balance: "+err.Error(), http.StatusInternalServerError)
		return
	}

	balanceEther := new(big.Float)
	wei := big.NewFloat(1e18)
	balanceEther.SetString(balance.String())
	balanceEther.Quo(balanceEther, wei)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"balance": balanceEther.String(),
	})
}
