package theater

import (
	"github.com/sirupsen/logrus"

	"gitlab.com/oiacow/nextfesl/network"
	"gitlab.com/oiacow/nextfesl/network/codec"
)

type reqGDAT struct {
	// TID=3
	TID int `fesl:"TID"`

	// LID=0
	LobbyID int `fesl:"LID"`
	// GID=1
	GameID int `fesl:"GID"`
}

type ansGDAT struct {
	TID string `fesl:"TID"`

	EloRank             string `fesl:"B-U-elo_rank"`
	AvgAllyRank         string `fesl:"B-U-avg_ally_rank"`
	AvgAxisRank         string `fesl:"B-U-avg_axis_rank"`
	ArmyDistribution    string `fesl:"B-U-army_distribution"`
	ArmyBalance         string `fesl:"B-U-army_balance"`
	PercentFull         string `fesl:"B-U-percent_full"`
	AvailSlotsNational  string `fesl:"B-U-avail_slots_national"`
	AvailSlotsRoyal     string `fesl:"B-U-avail_slots_royal"`
	AvailableVipsNation string `fesl:"B-U-avail_vips_national"`
	AvailableVipsRoyal  string `fesl:"B-U-avail_vips_royal"`
	IsRanked            string `fesl:"B-U-ranked"`
	Easyzone            string `fesl:"B-U-easyzone"`
	ServerType          string `fesl:"B-U-servertype"`
	ServerState         string `fesl:"B-U-server_state"`
	MapName             string `fesl:"B-U-map_name"`
	PunkBusterEnabled   string `fesl:"B-U-punkb"`
	StdDevLevel         string `fesl:"B-U-lvl_sdv"`
	AvgLevel            string `fesl:"B-U-lvl_avg"`

	GameID     int    `fesl:"GID"`
	Join       string `fesl:"JOIN"`
	ServerName string `fesl:"NAME"`

	AP      string `fesl:"AP"` // PlayerTypeCount, int
	LobbyID int    `fesl:"LID"`

	// GameType          string `fesl:"TYPE"` // = "P" ?

	// IpAddr string `fesl:"I,omitempty"` //string
	// Port string `fesl:"P,omitempty"` // int

	// Password string `fesl:"PW,omitempty"` // string

	// MaxPlayersCount string `fesl:"MP,omitempty"` // PlayerTypeCapacity, int
	// ActualPlayersCount    string `fesl:"AP"` // PlayerTypeCount, int
	// PlayerMaxObservers string `fesl:"B-maxObservers,omitempty"` // int
	// PlayerActualObservers string `fesl:"B-numObservers,omitempty"` // int
	// JoiningPlayersCount string `fesl:"JP,omitempty"` // int
	// QueuedPlayersCount string `fesl:"QP,omitempty"` // int

	// HostPlayerName string `fesl:"HN,omitempty"` // string
	// HostPlayerUserID string `fesl:"HU,omitempty"` // int

	// Version string `fesl:"V,omitempty"`
	// GameProtocolVersion string `fesl:"B-version,omitempty"`

	// IsFavorite string `fesl:"F,omitempty"` // int
	// FavPlayerCount string `fesl:"NF,omitempty"` // int

	// Platform string `fesl:"PL,omitempty"` // if PL=XBOX then XUID (XboxUserID must be specified)

	// JoinMode string `fesl:"J,omitempty"` // O / W / C (default=O, but if there will some random string not eaual to W and C it will also work)
}

// GDAT - CLIENT called to get data about the server
func (tm *Theater) GameData(event network.EventClientCommand) {
	gameID, err := event.Command.Message.IntVal("GID")
	if err != nil {
		logrus.WithError(err).Warn("Cannot parse GID in theater.GDAT")
		return
	}

	game, err := tm.mm.GetGame(gameID)
	if err != nil {
		logrus.
			WithError(err).
			WithField("gameID", gameID).
			Warn("Cannot find Game in matchmaking pool")
		return
	}

	event.Client.WriteEncode(&codec.Answer{
		Type: codec.ThtrGamesData,
		Payload: ansGDAT{
			LobbyID:             game.LobbyID,
			AP:                  game.GameServer.ServerData.Get("AP"),
			TID:                 event.Command.Message["TID"],
			GameID:              game.ID,
			Join:                game.GameServer.ServerData.Get("JOIN"),
			ServerName:          game.GameServer.ServerData.Get("NAME"),
			EloRank:             game.GameServer.ServerData.Get("B-U-elo_rank"),
			AvgAllyRank:         game.GameServer.ServerData.Get("B-U-avg_ally_rank"),
			AvgAxisRank:         game.GameServer.ServerData.Get("B-U-avg_axis_rank"),
			ArmyDistribution:    game.GameServer.ServerData.Get("B-U-army_distribution"),
			ArmyBalance:         game.GameServer.ServerData.Get("B-U-army_balance"),
			PercentFull:         game.GameServer.ServerData.Get("B-U-percent_full"),
			AvailSlotsNational:  game.GameServer.ServerData.Get("B-U-avail_slots_national"),
			AvailSlotsRoyal:     game.GameServer.ServerData.Get("B-U-avail_slots_royal"),
			AvailableVipsNation: game.GameServer.ServerData.Get("B-U-avail_vips_national"),
			AvailableVipsRoyal:  game.GameServer.ServerData.Get("B-U-avail_vips_royal"),
			IsRanked:            game.GameServer.ServerData.Get("B-U-ranked"),
			Easyzone:            game.GameServer.ServerData.Get("B-U-easyzone"),
			ServerType:          game.GameServer.ServerData.Get("B-U-servertype"),
			ServerState:         game.GameServer.ServerData.Get("B-U-server_state"),
			PunkBusterEnabled:   game.GameServer.ServerData.Get("B-U-punkb"),
			MapName:             game.GameServer.ServerData.Get("B-U-map_name"),
			AvgLevel:            game.GameServer.ServerData.Get("B-U-lvl_avg"),
			StdDevLevel:         game.GameServer.ServerData.Get("B-U-lvl_sdv"),
		},
	})
}
