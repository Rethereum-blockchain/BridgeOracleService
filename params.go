package BridgeOracleService

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
	Rpcs            []Rpc
	ContractAddress *common.Address
	PrivateKey      string
}

type Bridge struct {
	Contract   *BridgeContract
	Peer       *BridgeContract
	Transactor *bind.TransactOpts
	Channel    chan *BridgeContractActionRequested
}

var (
	config          Config
	clients         []*ethclient.Client
	bridgeContracts []*BridgeContract
	signers         []*bind.TransactOpts
)

func init() {
	_, err := toml.DecodeFile("config.toml", &config)
	if err != nil {
		log.Fatal("Error decoding config file")
	}
	privateKey, err := crypto.HexToECDSA(config.PrivateKey)
	if err != nil {
		log.Fatal("Error creating private key")
	}

	clients = make([]*ethclient.Client, len(config.Rpcs))
	for i, rpc := range config.Rpcs {
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
