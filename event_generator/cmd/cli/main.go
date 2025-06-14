package main

import (
	"math/rand"
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
		choice := rand.Intn(2)
		if choice == 0 {
			txHash, err := handlers.Spam(config)
			if err != nil {
				log.Println("❌ Error while spamming: ", err)
			} else if txHash != (common.Hash{}) {
				log.Println("✅ Tx hash: ", txHash.Hex())
			} else {
				log.Println("✅ Skipped")
			}
		} else {
			success, err := handlers.SpamEvent(config)
			if err != nil {
				log.Println("🚨 Error while creating event: ", err)
			} else if success {
				log.Println("✅ Event created successfully")
			} else {
				log.Println("✅ Skipped")
			}
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
