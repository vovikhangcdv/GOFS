package utils

import (
	"crypto/ecdsa"
	"fmt"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
)

func GetNthPrivateKey(wallet *hdwallet.Wallet, n int) (*ecdsa.PrivateKey, error) {
	path := hdwallet.MustParseDerivationPath(fmt.Sprintf("m/44'/60'/0'/0/%d", n))
	account, err := wallet.Derive(path, false)
	if err != nil {
		return nil, err
	}
	return wallet.PrivateKey(account)
}

func GetNextPrivateKey(wallet *hdwallet.Wallet, cnt *int) (*ecdsa.PrivateKey, error) {
	key, err := GetNthPrivateKey(wallet, *cnt)
	if err != nil {
		return nil, err
	}
	*cnt++
	return key, nil
}