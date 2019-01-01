package pnow

import (
	"gitlab.com/oiacow/nextfesl/network"
)

type reqStart struct {
	// TXN=Start
	TXN string `fesl:"TXN"`
	// partition.partition=
	Partition statusPartition `fesl:"partition"`
	// debugLevel=off
	DebugLevel string `fesl:"debugLevel"`
	// version=1
	Version int `fesl:"version"`
	// players.[]=1
	Players []reqStartPlayer
}

type reqStartPlayer struct {
	// players.0.ownerId=9
	OwnerID int `fesl:"ownerId"`

	// players.0.ownerType=1
	OwnerType int `fesl:"ownerId"`
	
	Properties map[string]interface{} `fesl:"props"`
}

const (
	debugLevelOff  = "off"
	debugLevelHigh = "high"
	debugLevelMed  = "med"
	debugLevelLow  = "low"
)

type ansStart struct {
	Txn string          `fesl:"TXN"`
	ID  statusPartition `fesl:"id"`
}

// Start handles pnow.Start
func (pnow *PlayNow) Start(event network.EventClientCommand) {
	pnow.answer(
		event.Client,
		event.Command.PayloadID,
		ansStart{
			Txn: pnowStart,
			ID:  statusPartition{1, event.Command.Message["eagames/bfwest-dedicated"]},
		},
	)
}
