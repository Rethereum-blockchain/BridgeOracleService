package main

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
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
			Peer:       bridgeContracts[i%(count-1)],
			Transactor: signers[i],
			Channel:    make(chan *BridgeContractActionRequested),
		}

		_, err := bridges[i].Contract.BridgeContractFilterer.WatchActionRequested(&bind.WatchOpts{}, bridges[i].Channel, nil)
		if err != nil {
			log.Println("Error creating filter", err.Error())
		}

		go bridges[i].Run()
	}

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	<-signalChan
}

func (bridge *Bridge) Run() {
	for {
		select {
		case eventLog := <-bridge.Channel:
			bridge.HandleLog(eventLog)
		}
	}
}

func (bridge *Bridge) HandleLog(event *BridgeContractActionRequested) {
	log.Println("Received event", event.Id)
	action, err := bridge.Peer.BridgeContractCaller.Actions(&bind.CallOpts{}, event.Id)
	if err != nil {
		log.Println("Error calling peer contract [Actions]", err.Error())
		return
	}

	if !action {
		log.Println(event.Id, "not found in peer contract")
		return
	}

	actionConsumed, err := bridge.Contract.BridgeContractCaller.ConsumedActions(&bind.CallOpts{}, event.Id)
	if err != nil {
		log.Println("Error calling us contract [ConsumedActions]", err.Error())
		return
	}

	if actionConsumed {
		log.Println(event.Id, "already consumed")
		return
	}

	log.Println("Authorizing action", event.Id)
	transcation, err := bridge.Contract.BridgeContractTransactor.AuthorizeAction(bridge.Transactor, event.Id)
	if err != nil {
		log.Println("Error authorizing action", err.Error())
		return
	}

	log.Println("Action authorized", transcation.Hash().Hex())
}
