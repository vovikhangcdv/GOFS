package flags

import "github.com/urfave/cli/v2"

var (
	RpcURL = &cli.StringFlag{
		Name: "rpc",
		Value: "http://127.0.0.1:8545",
		Usage: "RPC URL",
	}
	Seed = &cli.IntFlag{
		Name: "seed",
		Value: 0,
		Usage: "Seed for random number generator",
	}
	
	Flags = []cli.Flag{
		RpcURL,
		Seed,
	}
)