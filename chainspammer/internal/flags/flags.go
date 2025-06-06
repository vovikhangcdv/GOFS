package flags

import "github.com/urfave/cli/v2"

var (
	RpcURL = &cli.StringFlag{
		Name:  "rpc",
		Value: "http://127.0.0.1:8545",
		Usage: "RPC URL",
	}
	Seed = &cli.IntFlag{
		Name:  "seed",
		Value: 0,
		Usage: "Seed for random number generator",
	}

	FaucetSk = &cli.StringFlag{
		Name:  "faucet-sk",
		Value: "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80",
		Usage: "Faucet private key",
	}

	DelayTime = &cli.Int64Flag{
		Name:  "delay-time",
		Value: 1,
		Usage: "Delay time in seconds",
	}

	Mnemonic = &cli.StringFlag{
		Name:  "mnemonic",
		Value: "test test test test test test test test test test test junk",
		Usage: "Mnemonic for wallet",
	}

	EntityRegistry = &cli.StringFlag{
		Name:  "entity-registry",
		Value: "0x5FbDB2315678afecb367f032d93F642f64180aa3",
		Usage: "Entity registry address",
	}

	ComplianceRegistry = &cli.StringFlag{
		Name:  "compliance-registry",
		Value: "0xCf7Ed3AccA5a467e9e704C703E8D87F634fB0Fc9",
		Usage: "Compliance registry address",
	}

	EVNDToken = &cli.StringFlag{
		Name:  "evnd-token",
		Value: "0xDc64a140Aa3E981100a9becA4E685f962f0cF6C9",
		Usage: "EVND token address",
	}
	USDTToken = &cli.StringFlag{
		Name:  "musd-token",
		Value: "0x9A676e781A523b5d0C0e43731313A708CB607508",
		Usage: "mUSD token address",
	}

	ExchangePortal = &cli.StringFlag{
		Name:  "exchange-portal",
		Value: "0x0B306BF915C4d645ff596e518fAf3F9669b97016",
		Usage: "Exchange portal address",
	}

	RegisterEntityWeight = &cli.IntFlag{
		Name:  "register-entity-weight",
		Value: 1,
		Usage: "Weight for register entity transactions",
	}

	SendEVNDWeight = &cli.IntFlag{
		Name:  "send-evnd-weight",
		Value: 1,
		Usage: "Weight for send EVND transactions",
	}

	ExchangeVNDUSDWeight = &cli.IntFlag{
		Name:  "exchange-vnd-usd-weight",
		Value: 1,
		Usage: "Weight for exchange VND to USD transactions",
	}

	Flags = []cli.Flag{
		RpcURL,
		Seed,
		FaucetSk,
		DelayTime,
		Mnemonic,
		EntityRegistry,
		ComplianceRegistry,
		EVNDToken,
		USDTToken,
		ExchangePortal,
		RegisterEntityWeight,
		SendEVNDWeight,
		ExchangeVNDUSDWeight,
	}
)
