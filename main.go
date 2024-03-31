package main

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/event"
	"log"
	"os"
	"os/signal"
)

var (
	bridges []*Bridge
)

func main() {
	count := len(bridgeContracts)
	bridges = make([]*Bridge, count)

	for i, bridgeContract := range bridgeContracts {
		bridges[i] = &Bridge{
			Contract:   bridgeContract,
			Peer:       bridgeContracts[i^1],
			Transactor: signers[i],
			Channel:    make(chan *BridgeContractActionRequested),
		}

		sub, err := bridges[i].Contract.WatchActionRequested(&bind.WatchOpts{}, bridges[i].Channel, nil)
		if err != nil {
			log.Fatal("Error subscribing to events! ", err.Error())
		}

		go bridges[i].Run(sub)
	}

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	<-signalChan
}

func (bridge *Bridge) Run(sub event.Subscription) {
	log.Println("Started listening for ActionCreate")
	for {
		select {
		case eventLog := <-bridge.Channel:
			bridge.HandleLog(eventLog)
		case err := <-sub.Err():
			log.Fatal(err)
		}
	}
}

func (bridge *Bridge) HandleLog(event *BridgeContractActionRequested) {
	IdString := hexutil.Encode(event.Id[:])
	log.Println("Received event", IdString)
	action, err := bridge.Peer.BridgeContractCaller.Actions(&bind.CallOpts{}, event.Id)
	if err != nil {
		log.Println("Error calling peer contract [Actions]", err.Error())
		return
	}

	if !action {
		log.Println(IdString, "not found in peer contract")
		return
	}

	actionConsumed, err := bridge.Contract.BridgeContractCaller.ConsumedActions(&bind.CallOpts{}, event.Id)
	if err != nil {
		log.Println("Error calling us contract [ConsumedActions]", err.Error())
		return
	}

	if actionConsumed {
		log.Println(IdString, "already consumed")
		return
	}

	log.Println("Authorizing action", IdString)
	transcation, err := bridge.Contract.BridgeContractTransactor.AuthorizeAction(bridge.Transactor, event.Id)
	if err != nil {
		log.Println("Error authorizing action", err.Error())
		return
	}

	log.Println("Action authorized", transcation.Hash().Hex())
}
