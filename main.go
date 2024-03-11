package BridgeOracleService

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"log"
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
			Peer:       bridgeContracts[i%count],
			Transactor: signers[i],
			Channel:    make(chan *BridgeContractActionRequested),
		}

		_, err = bridges[i].Contract.BridgeContractFilterer.WatchActionRequested(&bind.WatchOpts{}, bridges[i].Channel, nil)
		if err != nil {
			log.Println("Error creating filter")
		}

		go bridges[i].Run()
	}
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
	action, err := bridge.Peer.BridgeContractCaller.Actions(&bind.CallOpts{}, event.Id)
	if err != nil {
		log.Println("Error calling peer contract [ConsumedActions]", err)
		return
	}

	if action {
		log.Println(event.Id, "already consumed")
		return
	}

	actionConsumed, err := bridge.Contract.BridgeContractCaller.ConsumedActions(&bind.CallOpts{}, event.Id)
	if err != nil {
		log.Println("Error calling us contract [ConsumedActions]", err)
		return
	}

	if actionConsumed {
		log.Println(event.Id, "already consumed")
		return
	}

	log.Println("Authorizing action", event.Id)
	_, err = bridge.Contract.BridgeContractTransactor.AuthorizeAction(bridge.Transactor, event.Id)
	if err != nil {
		log.Println("Error authorizing action", err)
		return
	}
}
