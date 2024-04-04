package main

import (
    "os"
    "os/signal"
)

func main() {
    for i, bridgeContract := range bridgeContracts {
        bridges[i] = &Bridge{
            Contract:   bridgeContract,
            Peer:       bridgeContracts[i^1],
            Transactor: signers[i],
            Channel:    make(chan *BridgeContractActionRequested),
        }

        bridges[i].Subscribe()
        go bridges[i].Run()
    }

    signalChan := make(chan os.Signal, 1)
    signal.Notify(signalChan, os.Interrupt)
    <-signalChan
}
