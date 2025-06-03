package utils

import (
	"crypto/rand"
	"crypto/ecdsa"
	"crypto/elliptic"
	mathRand "math/rand"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

)

func RandomCallData(length uint64) []byte {
	if length == 0 {
		return []byte{}
	}
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return b
}

func RandomSk() *ecdsa.PrivateKey {
	sk, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}
	return sk
}

func RandomSks(count int) []*ecdsa.PrivateKey {
	sk := make([]*ecdsa.PrivateKey, count)
	for i := 0; i < count; i++ {
		sk[i] = RandomSk()
	}
	return sk
}

func RandomAddress() common.Address {
	b := make([]byte, 20)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return common.BytesToAddress(b)
}

func RandomSkAndAddressFromList(arr []*ecdsa.PrivateKey) (*ecdsa.PrivateKey, common.Address) {
	if len(arr) == 0 {
		return nil, common.Address{}
	}
	sk := arr[mathRand.Intn(len(arr))]
	return sk, crypto.PubkeyToAddress(sk.PublicKey)
}
