package main

import (
	"github.com/BurntSushi/toml"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

type Rpc struct {
	ChainId *big.Int
	Url     string
}

type Config struct {
	RpcUrls         []Rpc
	ContractAddress *common.Address
	PrivateKey      string
}

var (
	config          Config
	clients         [2]*ethclient.Client
	bridgeContracts [2]*BridgeContract
	signers         [2]*bind.TransactOpts
	bridges         [2]*Bridge
)

func init() {
	_, err := toml.DecodeFile("config.toml", &config)
	if err != nil {
		log.Fatal("Error decoding config file", err.Error())
	}

	privateKey, err := crypto.HexToECDSA(config.PrivateKey)
	if err != nil {
		log.Fatal("Error creating private key")
	}

	for i, rpc := range config.RpcUrls {
		clients[i], err = ethclient.Dial(rpc.Url)
		if err != nil {
			log.Fatal("Error connecting to rpc url: ", rpc.Url)
		}

		signers[i], err = bind.NewKeyedTransactorWithChainID(privateKey, rpc.ChainId)
		if err != nil {
			log.Fatal("Error creating transactor", err.Error())
		}

		bridgeContracts[i], err = NewBridgeContract(*config.ContractAddress, clients[i])
		if err != nil {
			log.Fatal("Error creating bridge contract")
		}
	}
}
