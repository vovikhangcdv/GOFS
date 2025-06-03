package config

import (
    "os"
    "strconv"

    "crypto/ecdsa"

    "github.com/ethereum/go-ethereum/rpc"
    "github.com/vovikhangcdv/GOFS/chainspammer/internal/utils"
    "github.com/urfave/cli/v2"
)

type Config struct {
    Backend     *rpc.Client
    Faucet      *ecdsa.PrivateKey
    Keys        []*ecdsa.PrivateKey

    Seed        int64
}

func getEnv(key, defaultValue string) string {
    value := os.Getenv(key)
    if value == "" {
        return defaultValue
    }
    return value
}

func getEnvAsInt(key string, defaultValue int) int {
    value := getEnv(key, strconv.Itoa(defaultValue))
    intValue, err := strconv.Atoi(value)
    if err != nil {
        return defaultValue
    }
    return intValue
}

func NewConfigFromEnv() *Config {
    rpcUrl := getEnv("RPC_URL", "http://localhost:8545")
    backend, err := rpc.Dial(rpcUrl)
    if err != nil {
        panic(err)
    }
    seed := getEnvAsInt("SEED", 0)
    sks := utils.RandomSks(10)
    return &Config{
        Backend: backend,
        Seed: int64(seed),
        Keys: sks,
    }
}

func NewConfigFromContext(c *cli.Context) (*Config, error) {
    rpcUrl := c.String("rpc")
    backend, err := rpc.Dial(rpcUrl)
    if err != nil {
        return nil, err
    }
    seed := c.Int64("seed")
    sks := utils.RandomSks(10)
    return &Config{
        Backend: backend,
        Seed: seed,
        Keys: sks,
    }, nil
}