package main

import (
	"log"
	"os"
	"time"

	"github.com/urfave/cli/v2"
	"github.com/vovikhangcdv/GOFS/chainspammer/internal/config"
	"github.com/vovikhangcdv/GOFS/chainspammer/internal/flags"
	"github.com/vovikhangcdv/GOFS/chainspammer/internal/handlers"
	"github.com/ethereum/go-ethereum/common"
)

var spamCommand = &cli.Command{
	Name:   "spam",
	Usage:  "Spam transactions",
	Flags:  flags.Flags,
	Action: runSpam,
}

// TODO: Implement spammer
func runSpam(c *cli.Context) error {
	log.Println("Spammer targeted RPC URL: ", c.String("rpc"))
	config, err := config.NewConfigFromContext(c)
	if err != nil {
		return err
	}
	for {
		txHash, err := handlers.Spam(config)
		if txHash == (common.Hash{}) || err != nil {
			log.Println("❌ Error while spamming: ", err)
		} else {
			log.Println("✅ Tx hash: ", txHash.Hex())
		}
		time.Sleep(time.Duration(config.DelayTime) * time.Second)
	}
}

func initApp() *cli.App {
	app := cli.NewApp()
	app.Name = "chainspammer"
	app.Usage = "Spammer for sending transactions"
	app.Commands = []*cli.Command{
		spamCommand,
	}
	return app
}

var app = initApp()

func main() {
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
