package handlers

import (
	"bytes"
	"testing"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"

	"github.com/vovikhangcdv/GOFS/chainspammer/internal/generated_contracts/entity_registry"
	"github.com/ethereum/go-ethereum/common"
)

func TestHashEntity(t *testing.T) {
	domainSeparator := "0xd530561497a594305c786d319b98c21212400efaf389c38e47b9ca72194b781c"
	skVerifier := &ecdsa.PrivateKey{
		PublicKey: ecdsa.PublicKey{
			Curve: crypto.S256(),
		},
		D: big.NewInt(0x123),
	}
	skVerifier.PublicKey.X, skVerifier.PublicKey.Y = crypto.S256().ScalarBaseMult(skVerifier.D.Bytes())

	entity := entityRegistry.Entity{
		EntityAddress: common.HexToAddress("0x93BDBe2c9f0F5cec59175C51D0a39fAee42A4a6e"),
		EntityType: 1,
		EntityData: []byte{},
		Verifier: common.HexToAddress("0x476C88ED464EFD251a8b18Eb84785F7C46807873"),
	}

	hash := hashEntity(domainSeparator, entity)
	if hash != common.HexToHash("0x0ecb63805f2da61a85a88072ad3ffcaf15b62b5d10914dff05d72eaedda64c66") {
		t.Errorf("hashEntity returned wrong hash")
	}

	signature, err := crypto.Sign(hash.Bytes(), skVerifier)
	if err != nil {
		t.Errorf("error: %v", err)
	}
	signature[64] += 27

	if !bytes.Equal(signature, common.FromHex("0xe9d9866b9767b8448e1c36b480c07fbc8822afdf785c558b4194cf9b0ab9ad0b738859438906e2948585876843d11d1452a20d5edca631930fd77819db20d0e31b")) {
		t.Errorf("wrong signatur, want: %s, got: %s", "0xe9d9866b9767b8448e1c36b480c07fbc8822afdf785c558b4194cf9b0ab9ad0b738859438906e2948585876843d11d1452a20d5edca631930fd77819db20d0e31b", common.Bytes2Hex(signature))
	}
}