package main

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/event"
	"log"
)

type Bridge struct {
	Contract   *BridgeContract
	Peer       *BridgeContract
	Transactor *bind.TransactOpts
	Channel    chan *BridgeContractActionRequested
	Subscriber event.Subscription
}

func (b *Bridge) Run() {
	log.Println("Started listening for ActionCreate")
	for {
		select {
		case eventLog := <-b.Channel:
			b.HandleLog(eventLog)
		case err := <-b.Subscriber.Err():
			log.Println(err)
			b.Subscriber.Unsubscribe()
			b.Subscribe()
		}
	}
}

func (b *Bridge) Subscribe() {
	sub, err := b.Contract.BridgeContractFilterer.WatchActionRequested(&bind.WatchOpts{}, b.Channel, nil)
	if err != nil {
		log.Fatal("Error resubscribing to events! ", err.Error())
	}
	b.Subscriber = sub
}

func (b *Bridge) HandleLog(event *BridgeContractActionRequested) {
	IdString := hexutil.Encode(event.Id[:])
	log.Println("Received event", IdString)
	action, err := b.Peer.BridgeContractCaller.Actions(&bind.CallOpts{}, event.Id)
	if err != nil {
		log.Println("Error calling peer contract [Actions]", err.Error())
		return
	}

	if !action {
		log.Println(IdString, "not found in peer contract")
		return
	}

	actionConsumed, err := b.Contract.BridgeContractCaller.ConsumedActions(&bind.CallOpts{}, event.Id)
	if err != nil {
		log.Println("Error calling us contract [ConsumedActions]", err.Error())
		return
	}

	if actionConsumed {
		log.Println(IdString, "already consumed")
		return
	}

	log.Println("Authorizing action", IdString)
	transaction, err := b.Contract.BridgeContractTransactor.AuthorizeAction(b.Transactor, event.Id)
	if err != nil {
		log.Println("Error authorizing action", err.Error())
		return
	}

	log.Println("Action authorized tx:", transaction.Hash().Hex())
}
