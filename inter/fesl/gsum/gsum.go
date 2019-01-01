package gsum

import (
	"gitlab.com/oiacow/nextfesl/network"
	"gitlab.com/oiacow/nextfesl/network/codec"
)
// GameSummary probably stands for Game Summary
type GameSummary struct {
	//
}

func (gsum *GameSummary) answer(client *network.Client, pnum uint32, payload interface{}) {
	client.WriteEncode(&codec.Answer{
		Type:         codec.FeslGameSummary,
		PacketNumber: pnum,
		Payload:      payload,
	})
}

type Session struct {
	Txn string `fesl:"TXN"`
	// Games  []Game  `fesl:"games"`
	// Events []Event `fesl:"events"`
}

// GetSessionID handles gsum.GetSessionID command
func (gsum *GameSummary) GetSessionID(client *network.Client, event *codec.Command) {
	gsum.answer(client, 0, Session{
		Txn: "GetSessionId",
	})
}
