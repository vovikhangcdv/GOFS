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
		Value: "0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0",
		Usage: "Entity registry address",
	}

	ComplianceRegistry = &cli.StringFlag{
		Name:  "compliance-registry",
		Value: "0x0165878A594ca255338adfa4d48449f69242Eb8F",
		Usage: "Compliance registry address",
	}

	EVNDToken = &cli.StringFlag{
		Name:  "evnd-token",
		Value: "0x2279B7A0a67DB372996a5FaB50D91eAA73d2eBe6",
		Usage: "EVND token address",
	}
	USDTToken = &cli.StringFlag{
		Name:  "musd-token",
		Value: "0x09635F643e140090A9A8Dcd712eD6285858ceBef",
		Usage: "mUSD token address",
	}

	ExchangePortal = &cli.StringFlag{
		Name:  "exchange-portal",
		Value: "0x67d269191c92Caf3cD7723F116c85e6E9bf55933",
		Usage: "Exchange portal address",
	}

	RegisterEntityWeight = &cli.IntFlag{
		Name:  "register-entity.weight",
		Value: 1,
		Usage: "Weight for register entity transactions",
	}

	SendEVNDWeight = &cli.IntFlag{
		Name:  "send-evnd.weight",
		Value: 1,
		Usage: "Weight for send EVND transactions",
	}

	ExchangeVNDUSDWeight = &cli.IntFlag{
		Name:  "exchange-vnd-usd.weight",
		Value: 1,
		Usage: "Weight for exchange VND to USD transactions",
	}

	MaxKeys = &cli.IntFlag{
		Name:  "max-keys",
		Value: 100,
		Usage: "Maximum number of keys to generate",
	}

	LargeAmountTransfersWeight = &cli.IntFlag{
		Name:  "large-amount-transfers.weight",
		Value: 1,
		Usage: "Weight for Event: Large amount transfers",
	}

	MultipleOutgoingTransfersWeight = &cli.IntFlag{
		Name:  "multiple-outgoing-transfers.weight",
		Value: 1,
		Usage: "Weight for Event: Multiple outgoing transfers",
	}

	MultipleIncomingTransfersWeight = &cli.IntFlag{
		Name:  "multiple-incoming-transfers.weight",
		Value: 1,
		Usage: "Weight for Event: Multiple incoming transfers",
	}

	SuspiciousAddressInteractionsWeight = &cli.IntFlag{
		Name:  "suspicious-address-interactions.weight",
		Value: 1,
		Usage: "Weight for Event: Suspicious address interactions",
	}

	LargeAmountTransfersTotalAmount = &cli.Int64Flag{
		Name:  "large-amount-transfers.total-amount",
		Value: 1_000,
		Usage: "Total amount for Event: Large amount transfers (in Ether)",
	}

	MultipleOutgoingTransfersBlockDuration = &cli.Int64Flag{
		Name:  "multiple-outgoing-transfers.block-duration",
		Value: 10,
		Usage: "Block duration for Event: Multiple outgoing transfers",
	}

	MultipleOutgoingTransfersTotalTxs = &cli.Int64Flag{
		Name:  "multiple-outgoing-transfers.total-txs",
		Value: 5,
		Usage: "Total transactions for Event: Multiple outgoing transfers",
	}

	MultipleIncomingTransfersBlockDuration = &cli.Int64Flag{
		Name:  "multiple-incoming-transfers.block-duration",
		Value: 10,
		Usage: "Block duration for Event: Multiple incoming transfers",
	}

	MultipleIncomingTransfersTotalAmount = &cli.Int64Flag{
		Name:  "multiple-incoming-transfers.total-amount",
		Value: 1_000,
		Usage: "Total amount for Event: Multiple incoming transfers (in Ether)",
	}

	SuspiciousAddressInteractionsBlacklistAddresses = &cli.StringSliceFlag{
		Name:  "suspicious-address-interactions.blacklisted-addresses",
		Value: cli.NewStringSlice(
			"0xb1Ee7A142d267C1f36714E4a8F75612F20a79720",
		),
		Usage: "Blacklist addresses for Event: Suspicious address interactions",
	}

	EntityDataPath = &cli.StringFlag{
		Name:  "entity-data-path",
		Value: "../contracts/test/entity_data.csv",
		Usage: "Path to entity data CSV file",
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
		MaxKeys,
		LargeAmountTransfersWeight,
		MultipleOutgoingTransfersWeight,
		MultipleIncomingTransfersWeight,
		SuspiciousAddressInteractionsWeight,
		LargeAmountTransfersTotalAmount,
		MultipleOutgoingTransfersBlockDuration,
		MultipleOutgoingTransfersTotalTxs,
		MultipleIncomingTransfersBlockDuration,
		MultipleIncomingTransfersTotalAmount,
		SuspiciousAddressInteractionsBlacklistAddresses,
		EntityDataPath,
	}
)
