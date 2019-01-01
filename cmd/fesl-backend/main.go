package main

// Generate setters and getters for stats
// $ go generate ./cmd/fesl-backend
//go:generate go run ../stats-codegen/main.go -scan ../../model --getters ../../inter/ranking/getters.go --setters ../../inter/ranking/setters.go --adders ../../inter/ranking/adders.go

import (
	"context"
	"flag"

	"gitlab.com/oiacow/nextfesl/config"
	"gitlab.com/oiacow/nextfesl/inter/fesl"
	"gitlab.com/oiacow/nextfesl/inter/matchmaking"
	"gitlab.com/oiacow/nextfesl/inter/theater"
	"gitlab.com/oiacow/nextfesl/network"
	"gitlab.com/oiacow/nextfesl/storage/database"
	"gitlab.com/oiacow/nextfesl/storage/kvstore"

	//"github.com/google/gops/agent"
	"github.com/sirupsen/logrus"
	"github.com/subosito/gotenv"
)

func main() {
	initConfig()
	initLogger()

	ctx, cancel := context.WithCancel(context.Background())
	defer func() {
		cancel()
	}()
	startServer(ctx)

	// Use "github.com/google/gops/agent"
	//if err := agent.Listen(&agent.Options{}); err != nil {
		//logrus.Fatal(err)
	//}

	logrus.Println("Serving...")
	<-ctx.Done()
}

func initConfig() {
	var (
		configFile string
	)
	flag.StringVar(&configFile, "config", ".env", "Path to configuration file")
	flag.Parse()

	gotenv.Load(configFile)
	config.Initialize()
}

func initLogger() {
	logrus.SetLevel(config.LogLevel())
	logrus.SetFormatter(&logrus.JSONFormatter{})


	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
}

func startServer(ctx context.Context) {
	db, err := database.New()
	if err != nil {
		logrus.Fatal(err)
	}

	network.InitClientData()
	kvs := kvstore.NewInMemory()
	mm := matchmaking.NewPool()

	fesl.New(config.FeslClientAddr(), false, db, kvs, mm).ListenAndServe(ctx)
	fesl.New(config.FeslServerAddr(), true, db, kvs, mm).ListenAndServe(ctx)

	theater.New(config.ThtrClientAddr(), db, kvs, mm).ListenAndServe(ctx)
	theater.New(config.ThtrServerAddr(), db, kvs, mm).ListenAndServe(ctx)
}
