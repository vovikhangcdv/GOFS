# ChainSpammer

A tool for stress testing/creating traffic on blockchain networks by sending various types of transactions with configurable probability weights.

## Features

- Send 3 types of transactions:
  - Register new entity
  - Send eVND between verified entities
  - Exchange eVND to USD

## Installation

```bash
cd chainspammer
go mod tidy
go build -o build/spammer cmd/cli/main.go
```

## Usage

(Optional) Run `anvil` and deploy contracts:

```bash
anvil
cd contracts
forge script script/SimpleDeployAll.s.sol:SimpleDeployAll  --rpc-url 127.0.0.1:8545 --broadcast
```

```bash
chainspammer spam [flags]
```

### Available Flags

```bash
NAME:
   chainspammer spam - Spam transactions

USAGE:
   chainspammer spam [command options]

OPTIONS:
   --rpc value                                                                                                                      RPC URL (default: "http://127.0.0.1:8545")
   --seed value                                                                                                                     Seed for random number generator (default: 0)
   --faucet-sk value                                                                                                                Faucet private key (default: "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80")
   --delay-time value                                                                                                               Delay time in seconds (default: 1)
   --mnemonic value                                                                                                                 Mnemonic for wallet (default: "test test test test test test test test test test test junk")
   --entity-registry value                                                                                                          Entity registry address (default: "0x5FbDB2315678afecb367f032d93F642f64180aa3")
   --compliance-registry value                                                                                                      Compliance registry address (default: "0xCf7Ed3AccA5a467e9e704C703E8D87F634fB0Fc9")
   --evnd-token value                                                                                                               EVND token address (default: "0xDc64a140Aa3E981100a9becA4E685f962f0cF6C9")
   --musd-token value                                                                                                               mUSD token address (default: "0x9A676e781A523b5d0C0e43731313A708CB607508")
   --exchange-portal value                                                                                                          Exchange portal address (default: "0x0B306BF915C4d645ff596e518fAf3F9669b97016")
   --register-entity.weight value                                                                                                   Weight for register entity transactions (default: 1)
   --send-evnd.weight value                                                                                                         Weight for send EVND transactions (default: 1)
   --exchange-vnd-usd.weight value                                                                                                  Weight for exchange VND to USD transactions (default: 1)
   --max-keys value                                                                                                                 Maximum number of keys to generate (default: 100)
   --large-amount-transfers.weight value                                                                                            Weight for Event: Large amount transfers (default: 1)
   --multiple-outgoing-transfers.weight value                                                                                       Weight for Event: Multiple outgoing transfers (default: 1)
   --multiple-incoming-transfers.weight value                                                                                       Weight for Event: Multiple incoming transfers (default: 1)
   --suspicious-address-interactions.weight value                                                                                   Weight for Event: Suspicious address interactions (default: 1)
   --large-amount-transfers.total-amount value                                                                                      Total amount for Event: Large amount transfers (default: 1000000000000000000)
   --multiple-outgoing-transfers.block-duration value                                                                               Block duration for Event: Multiple outgoing transfers (default: 10)
   --multiple-outgoing-transfers.total-txs value                                                                                    Total transactions for Event: Multiple outgoing transfers (default: 5)
   --multiple-incoming-transfers.block-duration value                                                                               Block duration for Event: Multiple incoming transfers (default: 10)
   --multiple-incoming-transfers.total-amount value                                                                                 Total amount for Event: Multiple incoming transfers (default: 1000000000000000000)
   --suspicious-address-interactions.blacklisted-addresses value [ --suspicious-address-interactions.blacklisted-addresses value ]  Blacklist addresses for Event: Suspicious address interactions (default: "0x0000000000000000000000000000000000006969", "0x0000000000000000000000000000000000696969")
   --help, -h                                                                                                                       show help
```

### Example

```bash
# Run with default settings
chainspammer spam

# Run with custom RPC and transaction weights
chainspammer spam \
  --rpc http://localhost:8545 \
  --register-entity-weight 1 \
  --send-evnd-weight 0 \
  --exchange-vnd-usd-weight 0 \
  --delay-time 3
```

### Example logs
