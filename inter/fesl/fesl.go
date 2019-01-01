package fesl

import (
	"context"
	"time"

	"gitlab.com/oiacow/nextfesl/inter/fesl/acct"
	"gitlab.com/oiacow/nextfesl/inter/fesl/fsys"
	"gitlab.com/oiacow/nextfesl/inter/fesl/gsum"
	"gitlab.com/oiacow/nextfesl/inter/fesl/pnow"
	"gitlab.com/oiacow/nextfesl/inter/fesl/rank"
	"gitlab.com/oiacow/nextfesl/inter/matchmaking"
	"gitlab.com/oiacow/nextfesl/network"
	"gitlab.com/oiacow/nextfesl/storage/database"
	"gitlab.com/oiacow/nextfesl/storage/kvstore"

	"github.com/sirupsen/logrus"
)

// Manager - handles incoming and outgoing FESL data
type Manager struct {
	db     database.Adapter
	socket *network.Socket
	server bool

	acct *acct.Account
	fsys *fsys.ConnectSystem
	gsum *gsum.GameSummary
	pnow *pnow.PlayNow
	rank *rank.Ranking
}

// New
func New(bind string, server bool, db database.Adapter, kvs *kvstore.Storage, mm *matchmaking.Pool) *Manager {
	socket, err := network.NewSocketTLS(bind)
	if err != nil {
		logrus.Fatal(err)
		return nil
	}

	fm := &Manager{
		db,
		socket,
		server,
		&acct.Account{DB: db},
		&fsys.ConnectSystem{ServerMode: server},
		&gsum.GameSummary{},
		&pnow.PlayNow{MM: mm},
		&rank.Ranking{DB: db},
	}

	return fm
}

func (fm *Manager) ListenAndServe(ctx context.Context) {
	go fm.Run(ctx)
}

func (fm *Manager) Run(ctx context.Context) {
	for {
		select {
		case event := <-fm.socket.EventChan:
			fm.handleTCP(event)
		case <-ctx.Done():
			return
		}
	}
}

func (fm *Manager) handleTCP(event network.SocketEvent) {
	ev, ok := event.Data.(network.EventClientCommand)
	if !ok {
		logrus.Error("Logic error: Cannot cast event to network.EventClientCommand")
		return
	}

	// if !ev.Client.IsActive {
	// 	logrus.WithField("command", ev.Command).Warn("Inactive client")
	// 	return
	// }

	switch event.Name {
	case "newClient":
		fm.newClient(ev.Client) // TLS
	case "client.command.Hello":
		if !fm.server {
			fm.gsum.GetSessionID(ev.Client, ev.Command)
		}
		fm.fsys.Hello(ev.Client, ev.Command)
	case "client.command.NuLogin":
		fm.acct.NuLogin(ev)
	case "client.command.NuGetPersonas":
		fm.acct.NuGetPersonas(ev)
	case "client.command.NuGetAccount":
		fm.acct.NuGetAccount(ev)
	case "client.command.NuLoginPersona":
		fm.acct.NuLoginPersona(ev)
	case "client.command.GetStatsForOwners":
		fm.rank.GetStatsForOwners(ev)
	case "client.command.GetStats":
		fm.rank.GetStats(ev)
	case "client.command.NuLookupUserInfo":
		fm.acct.NuLookupUserInfo(ev)
	case "client.command.GetPingSites":
		fm.fsys.GetPingSites(ev)
	case "client.command.UpdateStats":
		fm.rank.UpdateStats(ev)
	// case "client.command.GetTelemetryToken": //used in 2k9 & blazeclient
	// 	fm.acct.GetTelemetryToken(ev)
	case "client.command.Start":
		fm.pnow.Start(ev)
		fm.pnow.Status(ev)
	case "client.command.MemCheck":
		// By now: Nothing interesting here, we can skip it.
		// TODO: Use MemCheck response for telemetry.
		return
	default:
		logrus.
			WithFields(logrus.Fields{
				"event":   event.Name,
				"payload": ev.Command.Message,
				"query":   ev.Command.Query,
			}).
			Warn("fesl.UnhandledRequest")
	}
}

// TLS
func (fm *Manager) newClient(client *network.Client) {
	fm.fsys.MemCheck(client)

	client.HeartTicker = time.NewTicker(55 * time.Second)
	go func() {
		for client.IsActive {
			select {
			case <-client.HeartTicker.C:
				if !client.IsActive {
					return
				}
				fm.fsys.MemCheck(client)
			}
		}
	}()

	logrus.Debug("NEW client connected")
}
