package matchmaking

import (
	"gitlab.com/oiacow/nextfesl/network"
)

type Game struct {
	ID         int
	LobbyID    int
	GameServer *network.Client

	PlayersJoining int
	PlayersPlaying int
	PlayerSlots    int
	// VipSlots       int
}
