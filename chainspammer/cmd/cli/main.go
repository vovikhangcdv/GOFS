package main

import (
	"os"
	"log"
    "fmt"
	"time"
    "github.com/vovikhangcdv/GOFS/chainspammer/internal/flags"
	"github.com/vovikhangcdv/GOFS/chainspammer/internal/config"
	"github.com/vovikhangcdv/GOFS/chainspammer/internal/utils"
	"github.com/vovikhangcdv/GOFS/chainspammer/internal/handlers"
	"github.com/ethereum/go-ethereum/params"
	"math/big"
	"github.com/urfave/cli/v2"
)

var spamCommand = &cli.Command{
	Name: "spam",
	Usage: "Spam transactions",
	Flags: flags.Flags,
	Action: runSpam,
}

// TODO: Implement spammer
func runSpam(c *cli.Context) error {
	fmt.Println("Spammer targeted RPC URL: ", c.String("rpc"))
	config, err := config.NewConfigFromContext(c)
	if err != nil {
		return err
	}
	for {
		sk, addr := utils.RandomSkAndAddressFromList(config.Keys)
		fmt.Println("Airdropping to: ", addr.Hex())
		// 1 ETH
		aidropValue := big.NewInt(params.Ether)
		if err := handlers.Airdrop(config, addr, aidropValue); err != nil {
			fmt.Println("Error airdropping: ", err)
			continue
		}
		fmt.Println("Spamming")
		if err := handlers.Spam(config, sk); err != nil {
			fmt.Println("Error spamming: ", err)
			continue
		}
		time.Sleep(1 * time.Second)
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