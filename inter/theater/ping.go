package theater

import (
	"gitlab.com/oiacow/nextfesl/network"
	"gitlab.com/oiacow/nextfesl/network/codec"
)

type reqPING struct {
	TID string `fesl:"TID"`
}

type ansPING struct {
	TID string `fesl:"TID"`
}

func (tm *Theater) PING(client *network.Client) {
	client.WriteEncode(&codec.Answer{
		Type:    codec.ThtrPing,
		Payload: ansPING{"0"},
	})
}
