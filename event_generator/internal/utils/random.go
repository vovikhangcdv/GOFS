package utils

import (
	"encoding/csv"
	"fmt"
	"os"
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

// readEntityData reads the CSV file and returns a slice of EntityData
func readEntityData(filePath string, idx int) (localtypes.EntityData, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return localtypes.EntityData{}, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	// Skip header
	if _, err := reader.Read(); err != nil {
		return localtypes.EntityData{}, fmt.Errorf("error reading header: %w", err)
	}

	records, err := reader.ReadAll()
	if err != nil {
		return localtypes.EntityData{}, fmt.Errorf("error reading records: %w", err)
	}

	var entity localtypes.EntityData
	for i, record := range records {
		if i != idx {
			continue
		}
		if len(record) < 18 { // Ensure we have all required fields
			continue
		}
		entity = localtypes.EntityData{
			Name:        record[0],
			IDNumber:    record[1],
			Birthday:    record[2],
			Gender:      record[3],
			Email:       record[4],
			Phone:       record[5],
			Address:     record[6],
			Nationality: record[7],
			Others:      record[8],
			Root:        record[9],
		}
	}

	return entity, nil
}

// getRandomEntityData returns a random entity data from the CSV file
func GetRandomEntityData(filePath string) (localtypes.EntityData, error) {
	idx := mathRand.Intn(1000)
	entity, err := readEntityData(filePath, idx)
	if err != nil {
		return localtypes.EntityData{}, fmt.Errorf("error reading entity data: %w", err)
	}
	return entity, nil
}
