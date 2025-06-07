package utils

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"math/big"
	mathRand "math/rand"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	localtypes "github.com/vovikhangcdv/GOFS/chainspammer/internal/types"
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

func RandomIdx(max int) int {
	return mathRand.Intn(max)
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

func RandomAddressFromList(arr []*ecdsa.PrivateKey) common.Address {
	if len(arr) == 0 {
		return common.Address{}
	}
	return crypto.PubkeyToAddress(arr[mathRand.Intn(len(arr))].PublicKey)
}

func RandomSkFromList(arr []*ecdsa.PrivateKey) *ecdsa.PrivateKey {
	if len(arr) == 0 {
		return nil
	}
	sk := arr[mathRand.Intn(len(arr))]
	return sk
}

func RandomSkAndAddressFromList(arr []*ecdsa.PrivateKey) (*ecdsa.PrivateKey, common.Address) {
	sk := RandomSkFromList(arr)
	return sk, crypto.PubkeyToAddress(sk.PublicKey)
}

func SelectTxType(types []localtypes.TxType) string {
	totalWeight := 0
	for _, t := range types {
		totalWeight += t.Weight
	}

	r := mathRand.Intn(totalWeight)
	currentWeight := 0

	for _, t := range types {
		currentWeight += t.Weight
		if r < currentWeight {
			return t.Type
		}
	}

	return types[0].Type
}

func SelectEvent(events []localtypes.Event) localtypes.Event {
	totalWeight := 0
	for _, e := range events {
		totalWeight += e.GetWeight()
	}

	r := mathRand.Intn(totalWeight)
	currentWeight := 0

	for _, e := range events {
		currentWeight += e.GetWeight()
		if r < currentWeight {
			return e
		}
	}

	return events[0]
}

func GetRandomEntityType() uint8 {
	return uint8(mathRand.Intn(3))
}

// RandomBigInt generates a random big.Int between 0 and max (inclusive)
func RandomBigInt(max *big.Int) *big.Int {
	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		panic(err)
	}
	return n
}
