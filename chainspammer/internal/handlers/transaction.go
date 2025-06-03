package handlers

import (
    "fmt"
    "crypto/ecdsa"
    "github.com/vovikhangcdv/GOFS/chainspammer/internal/config"
    "github.com/ethereum/go-ethereum/common"
    "math/big"
    "github.com/ethereum/go-ethereum/crypto"
)

// TODO
func Airdrop(config *config.Config, addr common.Address, value *big.Int) error {
    fmt.Println("Airdroped to: ", addr.Hex())
    return nil
}

// TODO
func Spam(config *config.Config, sk *ecdsa.PrivateKey) error {
    fmt.Println("Spamming from: ", crypto.PubkeyToAddress(sk.PublicKey).Hex())
    return nil
}